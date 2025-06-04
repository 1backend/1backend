/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package proxyservice

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
	"github.com/pkg/errors"
)

func (cs *ProxyService) StartCacheScanLoop(
	scanInterval time.Duration,
	cacheDir string,
	out chan<- []proxy.CertRecord,
) {
	go func() {
		ticker := time.NewTicker(scanInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				certs, err := scanCacheDir(cacheDir)
				if err != nil {
					// optionally log
					continue
				}
				out <- certs
			}
		}
	}()
}

func scanCacheDir(dir string) ([]proxy.CertRecord, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	records := []proxy.CertRecord{}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".crt") {
			continue
		}

		id := strings.TrimSuffix(file.Name(), ".crt")
		if !isValidDomain(id) {
			return nil, errors.New("invalid domain format: " + id)
		}

		certPath := filepath.Join(dir, id+".crt")
		keyPath := filepath.Join(dir, id+".key")

		certBytes, err := os.ReadFile(certPath)
		if err != nil {
			continue
		}
		keyBytes, err := os.ReadFile(keyPath)
		if err != nil {
			continue
		}

		certInfo, err := os.Stat(certPath)
		if err != nil {
			continue
		}
		keyInfo, err := os.Stat(keyPath)
		if err != nil {
			continue
		}

		created := certInfo.ModTime()
		updated := keyInfo.ModTime()
		if updated.Before(created) {
			updated = created
		}

		records = append(records, proxy.CertRecord{
			Id:        id,
			Cert:      string(certBytes),
			Key:       string(keyBytes),
			CreatedAt: created,
			UpdatedAt: updated,
		})
	}

	return records, nil
}

var domainRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9-]{1,63}\.)+[a-zA-Z]{2,}$`)

func isValidDomain(domain string) bool {
	return domainRegex.MatchString(domain)
}

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package fileservice

import (
	"crypto/sha256"
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"
)

func EncodeURLtoFileName(url string) string {
	hash := sha256.New()
	hash.Write([]byte(url))
	hashBytes := hash.Sum(nil)

	encoded := base64.URLEncoding.EncodeToString(hashBytes)

	return strings.TrimRight(encoded, "=")
}

func DownloadStorageFilePath(url string) string {
	fileName := EncodeURLtoFileName(url)
	return shardStoragePath(fileName)
}

func shardStoragePath(fileName string) string {
	return shardStoragePathWithBasis(fileName, fileName)
}

func shardStoragePathWithBasis(basis, fileName string) string {
	first := basis
	if len(first) > 2 {
		first = first[:2]
	}
	if len(first) < 2 {
		first += strings.Repeat("_", 2-len(first))
	}

	second := ""
	if len(basis) > 2 {
		second = basis[2:]
	}
	if len(second) > 2 {
		second = second[:2]
	}
	if len(second) < 2 {
		second += strings.Repeat("_", 2-len(second))
	}

	return filepath.Join(first, second, fileName)
}

// checkFileExistsAndSize checks if a file exists and returns its size.
func checkFileExistsAndSize(filename string) (int64, bool, error) {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return info.Size(), true, nil
}

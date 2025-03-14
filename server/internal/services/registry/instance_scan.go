/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package registryservice

import (
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/logger"
	registry "github.com/openorch/openorch/server/internal/services/registry/types"
	"github.com/pkg/errors"
)

func (ns *RegistryService) instanceScan() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		ns.triggerChan <- struct{}{}
	}()

	for {
		select {
		case <-ticker.C:
			ns.instanceScanCycle()

		case _, ok := <-ns.triggerChan:
			if !ok {
				logger.Error("RegistrySvc trigger channel closed")
			}

			ns.instanceScanCycle()

		case <-sigChan:
			return
		}
	}

}

func (ns *RegistryService) instanceScanCycle() {
	instances, err := ns.instanceStore.Query().Find()
	if err != nil {
		logger.Error("Failed to query instances: %v", err)
		return
	}

	for _, instance := range instances {
		err = ns.scanInstance(instance.(*registry.Instance))
		if err != nil {
			logger.Error("Failed to scan instance: %v", err)
			continue
		}
	}

}

// scan the port of the instance to see if its available, update lastHeartbeat if it is
func (ns *RegistryService) scanInstance(instance *registry.Instance) error {
	now := time.Now()
	listening := checkPortListening(instance.URL, 3*time.Second)
	lastHeartbeat := time.Now()
	duration := time.Since(now)

	var status registry.InstanceStatus

	switch {
	case !listening:
		status = registry.InstanceStatusUnreachable
	case listening && duration > 1*time.Second:
		status = registry.InstanceStatusDegraded
	case listening && duration <= 1*time.Second:
		status = registry.InstanceStatusHealthy
	}

	updateFields := map[string]any{
		"status": status,
	}

	if listening {
		updateFields["lastHeartbeat"] = lastHeartbeat
	}

	err := ns.instanceStore.Query(datastore.Equals([]string{"id"}, instance.Id)).
		UpdateFields(updateFields)
	if err != nil {
		return errors.Wrap(err, "Failed to update instance")
	}

	return nil
}

func checkPortListening(address string, timeout time.Duration) bool {
	addressParts := strings.Split(address, "://")
	address = addressParts[len(addressParts)-1]

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

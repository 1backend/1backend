/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registryservice

import (
	"net"
	"strings"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/pkg/errors"
)

const (
	instanceScanInterval           = 15 * time.Second
	instanceHeartbeatWriteInterval = 5 * time.Minute
)

func (ns *RegistryService) instanceScan() {
	ticker := time.NewTicker(instanceScanInterval)
	defer ticker.Stop()

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
	start := time.Now()
	listening := checkPortListening(instance.URL, 3*time.Second)
	lastHeartbeat := time.Now()
	duration := time.Since(start)

	var status registry.InstanceStatus

	switch {
	case !listening:
		status = registry.InstanceStatusUnreachable
	case listening && duration > 1*time.Second:
		status = registry.InstanceStatusDegraded
	case listening && duration <= 1*time.Second:
		status = registry.InstanceStatusHealthy
	}

	updateFields := map[string]any{}
	if instance.Status != status {
		updateFields["status"] = status
	}

	if listening && shouldPersistInstanceHeartbeat(instance.LastHeartbeat, lastHeartbeat) {
		updateFields["lastHeartbeat"] = lastHeartbeat
	}

	if len(updateFields) == 0 {
		return nil
	}

	err := ns.instanceStore.Query(datastore.Equals([]string{"id"}, instance.Id)).
		UpdateFields(updateFields)
	if err != nil {
		return errors.Wrap(err, "Failed to update instance")
	}

	return nil
}

func shouldPersistInstanceHeartbeat(previous time.Time, current time.Time) bool {
	return previous.IsZero() || current.Sub(previous) >= instanceHeartbeatWriteInterval
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

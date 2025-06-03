/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registry_svc

import (
	"testing"
)

// TestGenerateID tests the GenerateID method of the Instance struct.
func TestGenerateID(t *testing.T) {
	tests := []struct {
		name       string
		instance   Instance
		expectedID string
	}{
		{
			name: "With URL",
			instance: Instance{
				URL: "https://myserver.com:5981/user-svc",
			},
			expectedID: "https://myserver.com:5981/user-svc",
		},
		{
			name: "With IP",
			instance: Instance{
				Scheme: "https",
				IP:     "192.168.1.1",
				Port:   999,
			},
			expectedID: "https://192.168.1.1:999",
		},
		{
			name: "With Host",
			instance: Instance{
				Scheme: "http",
				Host:   "api.com",
				Port:   80,
			},
			expectedID: "http://api.com:80",
		},
		{
			name: "With Host and Port",
			instance: Instance{
				Scheme: "http",
				Host:   "api.com",
				Port:   8080,
			},
			expectedID: "http://api.com:8080",
		},
		{
			name: "With Missing URL, IP, and Host",
			instance: Instance{
				Scheme: "http",
				Port:   8080,
			},
			expectedID: "http://:8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.instance.BuildURL()

			if id != tt.expectedID {
				t.Errorf("BuildURL() = %v, want %v", id, tt.expectedID)
			}
		})
	}
}

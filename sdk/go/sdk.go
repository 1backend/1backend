/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	onebackendapi "github.com/1backend/1backend/clients/go"
	"github.com/google/uuid"

	"github.com/sony/sonyflake"
)

var sonyFlake *sonyflake.Sonyflake

func init() {
	sonyFlake = sonyflake.NewSonyflake(sonyflake.Settings{})
	if sonyFlake == nil {
		panic("Sonyflake not created")
	}
}

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Id generates a short, human-readable unique ID with a prefix using Sonyflake.
// Example: Id("usr") might return "usr_fm7lcQnPni".
// This format is compact and memorable, making it ideal for user-friendly identifiers.
// However, because Sonyflake is time-based and IDs can be guessed or enumerated,
// this ID type is not suitable for exposing sensitive resources without authentication.
// For scenarios requiring non-enumerable, opaque identifiers, use OpaqueId instead.
func Id(prefixes ...string) string {
	number, err := sonyFlake.NextID()
	if err != nil {
		panic(err)
	}

	if number == 0 {
		return string(base62[0])
	}

	b := make([]byte, 0)
	for number > 0 {
		remainder := number % 62
		number = number / 62
		b = append([]byte{base62[remainder]}, b...)
	}

	return strings.Join(prefixes, "_") + "_" + string(b)
}

// OpaqueId generates a non-enumerable, opaque ID with a prefix using UUID v4.
// Example: OpaqueId("file") might return "file_5f906bb0_10e8_4066_a032_a9ad0eae1fdb".
// These IDs are suitable for public exposure and access control scenarios
// where predictability or enumeration must be avoided.
func OpaqueId(prefix string) string {
	return fmt.Sprintf("%v_%v",
		prefix,
		strings.Replace(uuid.New().String(), "-", "_", -1),
	)
}

// DeterministicId generates a deterministic, human-readable ID from a prefix and a source identifier.
// It is designed for use cases where consistent derived IDs are required—especially in eventually consistent systems.
//
// Use cases include:
//   - Idempotent operations (e.g., same insert won't duplicate data)
//   - Derived records (e.g., credit/debit entries from a transaction ID)
//   - Stable keys in distributed writes
//
// The resulting ID is formatted as: <prefix>_<cleaned source ID>
// Hyphens in either component are replaced with underscores for consistency.
//
// Example:
//
//	DeterministicId("txn_debit", "order123")   → "txn_debit_order123"
//	DeterministicId("txn_credit", "order123")  → "txn_credit_order123"
//
// This avoids clashing with auto-generated IDs while remaining readable and traceable.
func DeterministicId(prefix, id string) string {
	prefix = strings.TrimSuffix(prefix, "_")
	prefix = strings.ReplaceAll(prefix, "-", "_")
	id = strings.ReplaceAll(id, "-", "_")

	return fmt.Sprintf("%s_%s", prefix, id)
}

// OneBackendAPIError checks if an error is a GenericOpenAPIError and returns a meaningful error.
func OneBackendAPIError(err error) error {
	if err == nil {
		return nil
	}

	// Check if it's a GenericOpenAPIError
	if apiErr, ok := err.(*onebackendapi.GenericOpenAPIError); ok {
		var errorResponse map[string]interface{}
		if unmarshalErr := json.Unmarshal(apiErr.Body(), &errorResponse); unmarshalErr == nil {
			if message, exists := errorResponse["error"]; exists {
				return errors.New(message.(string))
			}
			return fmt.Errorf("unknown error format: %v", string(apiErr.Body()))
		}
		return fmt.Errorf("failed to unmarshal API error response: %v", string(apiErr.Body()))
	}

	// Return the original error if it's not a GenericOpenAPIError
	return err
}

func Marshal(value any) *string {
	jsonBytes, _ := json.Marshal(value)

	v := string(jsonBytes)
	return &v
}

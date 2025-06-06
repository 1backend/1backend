/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import (
	"time"
)

type KeyPair struct {
	Id        string    `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	PrivateKey string `json:"privateKey" binding:"required"`
	PublicKey  string `json:"publicKey" binding:"required"`
}

func (k *KeyPair) GetId() string {
	return k.Id
}

type GetPublicKeyRequest struct{}

type GetPublicKeyResponse struct {
	PublicKey string `json:"publicKey" binding:"required"`
}

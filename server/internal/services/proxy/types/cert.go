package proxy_svc

import "time"

type CertRecord struct {
	Id        string    `json:"id"`
	Cert      string    `json:"cert"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

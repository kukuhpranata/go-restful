package token

import "time"

type Token struct {
	AuthId         int       `json:"auth_id"`
	ExpirationTime time.Time `json:"expiration_time"`
}

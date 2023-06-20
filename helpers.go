package gochat_socket

import "github.com/gofrs/uuid"

func newV4UUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

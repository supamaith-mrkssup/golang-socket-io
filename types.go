package gochat_socket

import (
	"go-chat-api/internal/pkg/gochat-socket/parser"
	"reflect"
)

// namespace
const (
	aliasRootNamespace = "/"
	rootNamespace      = ""
)

// message
const (
	clientDisconnectMsg = "client namespace disconnect"
)

type readHandler func(c *conn, header parser.Header) error

var (
	defaultHeaderType = []reflect.Type{reflect.TypeOf("")}
)

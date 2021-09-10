package odoo

import (
	"fmt"
	"strings"
)

type Server struct {
	Host, Database, User, HiddenPassword, Password string
	UID                                            int
}

func NewServer(host, database, user, password string) *Server {
	return &Server{
		Host:     host,
		Database: database,
		User:     user,
		Password: password,
	}
}
func (server Server) GetName() string {
	return fmt.Sprintf("%s/?db=%s", cleanHost(server.Host), server.Database)
}
func cleanHost(host string) string {
	host = strings.Trim(strings.ToLower(host), "/")
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}
	return host
}

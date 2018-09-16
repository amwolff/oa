package common

import "fmt"

func GetDsn(user, pass, host string, port int, name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, name)
}

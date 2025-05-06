// Package ramsql contains RamSQL driver registration for xk6-sql.
package gocql

import (
	"github.com/grafana/xk6-sql/sql"

	// Blank import required for initialization of driver.
	_ "github.com/gocql/gocql"
)

func init() {
	sql.RegisterModule("gocql")
}

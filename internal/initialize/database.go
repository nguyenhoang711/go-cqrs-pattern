package initialize

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	database "github.com/nguyenhoang711/go-cqrs-pattern/database/sqlc"
	"github.com/nguyenhoang711/go-cqrs-pattern/global"
)

func InitDB() {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		"root",
		global.Config.PostgreSQL.Password,
		"localhost",
		5432,
		global.Config.PostgreSQL.DBName,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect database", err)
	}
	// conn.SetMaxOpenConns(25)
	// conn.SetMaxIdleConns(25)

	global.Logger.Info("Connect database successfully!")
	global.Db = database.NewStore(conn)
}

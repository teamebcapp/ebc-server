package postgres

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDbConnect is
func InitDbConnect() error {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=ebcmaster dbname=ebc password=ebcmaster sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return err
	}
	PostgresConn = db
	//db.DB().SetMaxIdleConns(10)
	PostgresConn.DB().SetMaxOpenConns(10)
	PostgresConn.LogMode(true)
	//PostgresConn.SetLogger(log.New(os.Stdout, "\r\n", 0))

	return nil
}

// PostgresConn is
var PostgresConn *gorm.DB

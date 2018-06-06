package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite db driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// NewDatabase - Generic Database connection func
func NewDatabase(dbType string, dbHost string, dbUser string, dbPassword string, dbName string, dbPort string) *DB {
	var connectString string
	if dbType == "postgres" {
		connectString = "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword
	}
	if dbType == "mysql" {
		connectString = dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName
	}
	if connectString != "" {
		db, err := grom.Open(dbType, connectString)
		if err != nil {
			panic(err)
		}

		if err = db.DB().Ping(); err != nil {
			panic(err)
		}

		// db.LogMode(true)

		return &DB{db}
	} else {
		panic("Must Define connection information")
	}
}

// NewPostgresDB - postgres database
func NewPostgresDB(dataSourceName string) *DB {

	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	//db.LogMode(true)

	return &DB{db}
}

// NewSqliteDB - sqlite database
func NewSqliteDB(databaseName string) *DB {

	db, err := gorm.Open("sqlite3", databaseName)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	//db.LogMode(true)

	return &DB{db}
}

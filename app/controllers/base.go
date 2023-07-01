package controllers

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log"
	"mvc_go/database"
	"net/http"
)

type Server struct {
	DB     *sqlx.DB
	Router *echo.Echo
}

type DBConfig struct {
	DBUsername string
	DBPassword string
	DBName     string
}

func (server *Server) InitServer(dbConfig DBConfig) {
	fmt.Println("Service started")

	server.initializeDB(dbConfig)
	server.initializeRoutes()
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	connectDbStr := mysql.Config{
		DBName:               dbConfig.DBName,
		User:                 dbConfig.DBUsername,
		Passwd:               dbConfig.DBPassword,
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sqlx.Open("mysql", connectDbStr.FormatDSN())

	if err != nil {
		panic(fmt.Sprintf("DB connection established failed: %v", err))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("Ping to DB failed: %v", err))
	}

	database.CreateTable(db)
	database.SeedTable(db)

	server.DB = db
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

package main

import (
	"clinic/server"
	"clinic/server/handler"
	"clinic/server/repository"
	"clinic/server/service"
	"flag"
	_ "github.com/lib/pq"
)

func main() {
	var dbHost string
	var port string
	var username string
	var dbname string
	var dbpassword string
	flag.StringVar(&dbHost, "dbhost", "127.0.0.1", "host of database")
	flag.StringVar(&port, "port", "5432", "port of database")
	flag.StringVar(&username, "username", "postgres", "username for database")
	flag.StringVar(&dbname, "dbname", "agewell", "name of database")
	flag.StringVar(&dbpassword, "dbpassword", "root", "password for database")

	db, _ := repository.NewPostgresDB(repository.Config{
		Host:     dbHost,
		Port:     port,
		Username: username,
		DBName:   dbname,
		SSLMode:  "disable",
		Password: dbpassword,
	})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	_ = srv.Run("8087", handlers.InitRoutes())
}

package main

import (
	"clinic/server"
	"clinic/server/handler"
	"clinic/server/repository"
	"clinic/server/service"
	"flag"
	"github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	var dbHost string
	var port string
	var username string
	var dbname string
	var dbpassword string
	flag.StringVar(&dbHost, "dbhost", "localhost", "host of database")
	flag.StringVar(&port, "port", "5432", "port of database")
	flag.StringVar(&username, "username", "postgres", "username for database")
	flag.StringVar(&dbname, "dbname", "clinic", "name of database")
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

	opts := mqtt.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883").SetClientID("GoClient")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(handlers.HandleMQTTMessage)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer client.Disconnect(250)

	if token := client.Subscribe("iot/data", 0, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	select {}
}

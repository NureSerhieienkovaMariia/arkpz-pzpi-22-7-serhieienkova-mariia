package main

import (
	"clinic/server"
	"clinic/server/handler"
	"clinic/server/repository"
	"clinic/server/service"
	"flag"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
	"log"
	"os"
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
	go func() {
		if err := srv.Run("8087", handlers.InitRoutes()); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Create a new MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://broker.hivemq.com:1883") // Replace with your broker address
	opts.SetClientID("age_well_go_mqtt_client")
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected to broker!")
	}
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
	}

	// Create and connect the client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Failed to connect: %v\n", token.Error())
		os.Exit(1)
	}

	// Subscribe to a topic
	topic := "agewell/data"
	if token := client.Subscribe(topic, 0, handlers.HandleMQTTMessage); token.Wait() && token.Error() != nil {
		fmt.Printf("Failed to subscribe: %v\n", token.Error())
		os.Exit(1)
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)

	// Keep the program running
	select {}
}

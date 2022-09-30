package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"os"

	"github.com/streadway/amqp"
)

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main() {
	router := httprouter.New()
	router.POST("/publish/:message", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		submit(w, r, p)
	})
	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func submit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	message := p.ByName("message")
	fmt.Println("received message: " + message)
	conn, err := amqp.Dial("amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to Rabbit", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"publisher",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}
	fmt.Println("publish successfully!")
}

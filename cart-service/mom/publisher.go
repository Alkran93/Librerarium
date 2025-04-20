package mom

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func PublishCheckout(username string, cart interface{}) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	var amqpURL = "amqp://" + os.Getenv("MOM_USER") + ":" +
		os.Getenv("MOM_PASSWORD") + "@" +
		os.Getenv("MOM_HOST") + ":" +
		os.Getenv("MOM_PORT") + "/"
	var exchange = os.Getenv("MOM_EXCHANGE")
	var routingKey = os.Getenv("MOM_ROUTING_KEYE")

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, _ := json.Marshal(map[string]interface{}{
		"evento":  "checkout",
		"usuario": username,
		"items":   cart,
	})

	err = ch.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return err
	}

	log.Printf("[MOM] Checkout enviado por %s\n", username)
	return nil
}

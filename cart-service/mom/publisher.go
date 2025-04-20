package mom

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

const amqpURL = "amqp://user:password@3.82.109.178:5672/"
const exchange = "my_exchange"
const routingKey = "test"

func PublishCheckout(username string, cart interface{}) error {
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

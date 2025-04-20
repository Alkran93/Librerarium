package mom

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

var exchangeName = "my_exchange"
var routingKey = "test"
var amqpURL = "amqp://user:password@34.205.157.11:5672/"

func PublishCheckout(cart interface{}) error {
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

	err = ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, _ := json.Marshal(map[string]interface{}{
		"evento": "checkout",
		"items":  cart,
	})

	err = ch.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return err
	}

	log.Println("[MOM] Checkout published")
	return nil
}

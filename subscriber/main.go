package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/erry-azh/mqtt-on-go/config"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	hostname, _ := os.Hostname()

	mqttHost, err := config.GetMQTTHost()
	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("connecting: ", mqttHost)

	opts := mqtt.NewClientOptions().AddBroker(mqttHost)
	opts.SetClientID(hostname)
	opts.SetCleanSession(true)
	opts.OnConnect = registerSubs

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic(token.Error())
		return
	}

	log.Println("connect success: ", mqttHost)

	<-signals
}

func registerSubs(c mqtt.Client) {
	log.Println("subscribing mqtt-go/test at qos-2")
	if token := c.Subscribe("mqtt-go/test", byte(2), sampleSubs); token.Wait() && token.Error() != nil {
		log.Print(token.Error())
	}
}

func sampleSubs(_ mqtt.Client, msg mqtt.Message) {
	if msg.Qos() == byte(2) {
		msg.Ack()
	}

	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	fmt.Printf("FULL_MSG: %+v\n", msg)
}

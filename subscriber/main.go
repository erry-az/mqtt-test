package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	hostname, _ := os.Hostname()
	server := flag.String("server", "tcp://127.0.0.1:1883", "The full URL of the mqtt server to connect to")
	topic := flag.String("topic", "", "Topic to publish the messages on")
	qos := flag.Int("qos", 0, "The QoS to send the messages at")
	flag.Parse()

	log.Println("connecting: ", *server)

	opts := mqtt.NewClientOptions().AddBroker(*server)
	opts.SetClientID(hostname)
	opts.SetCleanSession(false)
	opts.SetPingTimeout(10 * time.Second)
	opts.SetKeepAlive(10 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(10 * time.Second)

	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		fmt.Printf("!!!!!! mqtt connection lost error: %s\n" + err.Error())
	})

	opts.SetReconnectingHandler(func(c mqtt.Client, options *mqtt.ClientOptions) {
		fmt.Println("...... mqtt reconnecting ......")
	})

	opts.OnConnect = func(c mqtt.Client) {
		log.Printf("subscribing %s at qos-%d\n", *topic, *qos)
		if token := c.Subscribe(*topic, byte(*qos), sampleSubs); token.Wait() && token.Error() != nil {
			log.Print(token.Error())
		}
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic(token.Error())
		return
	}

	log.Println("connect success: ", *server)

	<-signals

	client.Disconnect(5)
	client.Unsubscribe("mqtt-go/test")
}

func sampleSubs(_ mqtt.Client, msg mqtt.Message) {
	if msg.Qos() == byte(2) {
		msg.Ack()
	}

	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	fmt.Printf("FULL_MSG: %+v\n", msg)
	fmt.Println("---------------------------------")
}

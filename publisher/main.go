package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	stdin := bufio.NewReader(os.Stdin)
	hostname, _ := os.Hostname()

	server := flag.String("server", "tcp://127.0.0.1:1883", "The full URL of the mqtt server to connect to")
	topic := flag.String("topic", hostname, "Topic to publish the messages on")
	qos := flag.Int("qos", 0, "The QoS to send the messages at")
	retained := flag.Bool("retained", false, "Are the messages sent with the retained flag")
	clientID := flag.String("client_id", hostname+strconv.Itoa(time.Now().Second()), "A clientID for the connection")
	username := flag.String("username", "", "A username to authenticate to the mqtt server")
	password := flag.String("password", "", "Password to match username")
	flag.Parse()

	connOpts := mqtt.NewClientOptions().AddBroker(*server).SetClientID(*clientID).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)
	connOpts.SetPingTimeout(10 * time.Second)
	connOpts.SetKeepAlive(10 * time.Second)
	connOpts.SetAutoReconnect(true)
	connOpts.SetMaxReconnectInterval(10 * time.Second)

	connOpts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		fmt.Printf("!!!!!! mqtt connection lost error: %s\n" + err.Error())
	})

	connOpts.SetReconnectingHandler(func(c mqtt.Client, options *mqtt.ClientOptions) {
		fmt.Println("...... mqtt reconnecting ......")
	})

	connOpts.OnConnect = func(client mqtt.Client) {
		fmt.Println("...... connect success ......")
		for {
			message, err := stdin.ReadString('\n')
			if err == io.EOF {
				os.Exit(0)
			}
			client.Publish(*topic, byte(*qos), *retained, message)
		}
	}

	client := mqtt.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
	fmt.Printf("Connected to %s\n", *server)

	<-signals

	client.Disconnect(10)
}

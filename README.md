# mqtt-on-go

## Info
- more info about MQTT https://mqtt.org/
- VerneMQ is MQTT message broker that will be use for this test https://vernemq.com/
- HAProxy as loadbalancer https://github.com/lelylan/haproxy-mqtt/blob/master/haproxy.cfg
- change ulimit with this instruction https://docs.vernemq.com/guides/change-open-file-limits
- use this tools if you want to benchmark connection and message
https://github.com/krylovsk/mqtt-benchmark

## Commands

- ### Run VerneMQ
```
make run
```
use this command for run docker verneMQ (mqtt broker) 4 nodes

- ### Stop VerneMQ 
```
make stop
``` 
use this command for stop docker compose

- ### Publisher
```
make publish PUB_TOPIC=<topic_name> PUB_QOS=<topic_qos>
``` 
use this command to create publisher with defined topic and qos

- ### Consumer
```
make subscribe SUB_TOPIC=<topic_name> SUB_QOS=<topic_qos>
``` 
use this command to create subscriber with defined topic and qos

## Web client
you can use web client on this repo on `mqtt_web_client` folder, but you must use websocket over mqtt port
(created by : matias alvin https://github.com/alvinmatias69) 

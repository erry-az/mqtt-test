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
use this command for run docker verneMQ (mqtt broker) 3 nodes.

- ### Clear Docker VerneMQ 
```
make clear
``` 
use this command for clear container

- ### Publisher
```
make publish ARG='-server=0.0.0.0:1883 -topic=abc/www -qos=1/2/3 -retained=true/false -id=str -username=str -password=str'
``` 
use this command to create publisher with defined topic and qos

- ### Subscribe
```
make subscribe ARG='-server=0.0.0.0:1883 -topic=abc/www -qos=1/2/3 -id=str -clean=true/false'
``` 
use this command to create subscriber with defined topic and qos

## Web client
you can use web client on this repo on `mqtt_web_client` folder, but you must use websocket over mqtt port
(created by : matias alvin https://github.com/alvinmatias69) 

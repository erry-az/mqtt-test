# mqtt-on-go


---
## Info
- more info about MQTT https://mqtt.org/
- VerneMQ is MQTT message broker that will be use for this test https://vernemq.com/

---
## Commands

- ### Run VerneMQ
```
make run VMQ_REPLICA=<number>
```
use this command for run docker verneMQ (mqtt broker), 
  `VMQ_REPLICA` will be use for define how many hosts will be created

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
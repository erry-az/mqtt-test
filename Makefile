.PHONY: run
run:
	@docker-compose -f docker-compose.yaml up --scale vmq=3

stop:
	@docker-compose -f docker-compose.yaml down --remove-orphans

.PHONY: build-webhook
build-webhook:
	@go build -o bin/webhook webhook/*.go

build-sub:
	@go build -o bin/subscriber subscriber/*.go

build-pub:
	@go build -o bin/publisher publisher/*.go

build:
	@make build-sub
	@make build-pub

subscribe:
	@make build-sub
	@./bin/subscriber -topic=$(SUB_TOPIC) -qos=$(SUB_QOS)

publish:
	@make build-pub
	@./bin/publisher -topic=$(PUB_TOPIC) -qos=$(PUB_QOS) -retained=$(PUB_RETAINED)

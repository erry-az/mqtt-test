.PHONY: run
run:
	@docker-compose -f docker-compose.yaml up

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
	@./bin/subscriber $(ARG)

publish:
	@make build-pub
	@./bin/publisher $(ARG)

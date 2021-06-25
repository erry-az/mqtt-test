.PHONY: run
run:
	docker-compose -f docker-compose.yaml up --build vmq

stop:
	docker-compose -f docker-compose.yaml down --remove-orphans
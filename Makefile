.PHONY: db
db:
	docker-compose exec db mysql -uroot -ppassword -Dvuln_app  --default-character-set=utf8mb4

.PHONY: clear
clear:
	docker-compose down -v

.PHONY: build
build:
	docker-compose build

.PHONY: start
start:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down

.PHONY: restart
restart: build stop start

.PHONY: reset
reset: clear start

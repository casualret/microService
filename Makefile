rm:
	docker-compose stop \
	&& docker-compose rm \
	&& sudo rm -rf data

build:
	docker-compose up --build

up:
	docker-compose up
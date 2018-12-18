all: clean build run
clean:
	docker-compose down
build: 
	docker-compose build
run:
	docker-compose up

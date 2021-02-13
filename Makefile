build:
	docker-compose build --no-cache
stop:
	docker-compose stop; docker-compose down
ps:
	docker-compose ps
start:
	docker-compose up
db:
	docker exec -it calender-db bash

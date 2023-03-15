start:
	docker-compose --project-name="postgres15.1" up -d
stop:
	docker-compose --project-name="postgres15.1" down
psql:
	docker exec -it pg psql -U user -d store
clean:
	rm -rf ./pgdata
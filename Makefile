
start:
	docker-compose up -d  postgres prometheus grafana

stop:
	docker-compose stop postgres prometheus grafana

full_start:
	docker-compose up -d

full_stop:
	docker-compose stop
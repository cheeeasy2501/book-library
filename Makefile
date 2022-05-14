
start:
	docker-compose --env-file configs/.env.local up -d postgres-master prometheus grafana

stop:
	docker-compose stop postgres-master prometheus grafana

full_start:
	docker-compose --env-file configs/.env.local up -d

full_stop:
	docker-compose stop
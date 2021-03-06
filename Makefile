
start:
	docker-compose --env-file configs/.env.local up -d postgres-master prometheus grafana

stop:
	docker-compose stop postgres-master prometheus grafana

full_start:
	docker-compose --env-file configs/.env.local up -d

full_stop:
	docker-compose stop

migration_drop:
	migrate -path ./db/migration -database 'postgres://postgres:postgres@localhost:5432/books-library?sslmode=disable' drop

migration_up:
	migrate -path ./db/migration -database 'postgres://postgres:postgres@localhost:5432/books-library?sslmode=disable' up
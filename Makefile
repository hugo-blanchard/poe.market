.PHONY: postgres adminer migrate demigrate

postgres:
	docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -p 8080:8080 adminer

migrate:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

demigrate:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable down
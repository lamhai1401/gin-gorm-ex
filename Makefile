SHELL := /bin/bash

ql:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen init
	go run ./server.go

gen:
	go run github.com/99designs/gqlgen generate

sql:
	docker-compose up

run:
	go run server.go

.PHONY: test
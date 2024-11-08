#!/bin/bash

.PHONY: dev
dev:
	make buildbin
	docker buildx build -t eip/api:latest --platform linux/amd64 -f Dockerfile .
	# docker push eip/api:latest 
	rm app

.PHONY: patch
patch:
	bump_version patch main.go

.PHONY: buildbin
buildbin:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .
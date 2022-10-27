build:
	@echo "Building app"
	go build -o app cmd/server/main.go

local-run:
	@echo "Running docker images on this instance!"
	docker-compose up --build
	
run:
	@echo "Running docker container in the background"
	docker-compose up --build -d

stop:
	@echo "Stopped the images"
	docker-compose stop
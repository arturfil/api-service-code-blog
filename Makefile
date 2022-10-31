build:
	@echo "Building app"
	go build -o app cmd/server/main.go

local-run:
	@echo "Running docker images on this instance!"
	docker-compose up --build
	
run:
	@echo "Running docker container in the background"
	docker-compose up --build -d

test:
	curl --location --request POST "http://localhost:8080/api/v1/posts/post" -v \
	--header "Content-Type: application/json" \
	--header 'Authorization: bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.GjnHO554SznHdy1LdjeTcfG6zdQr1OHXA6AibKaQ8VpbdrXb6pWly51H8cNHjsIQyFCPHgFwffAWC6v0Jn2ARA' \
	--data-raw '{
		"title":"Test",
		"name": "Arturo F",
		"content":"const test = () => return some",
		"category":"javascript"
	}'

stop:
	@echo "Stopped the images"
	docker-compose stop
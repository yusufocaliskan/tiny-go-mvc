
# Starts the app
run:
	go run main.go

#Teting
test-all:
	docker-compose -f docker-compose.dev.yml exec gptv_backend go test -v ./...

test:
	docker-compose -f docker-compose.dev.yml exec gptv_backend go test -v 

# start the application in hotreload using fresh.yaml 
# Use it in dev mode
# FormatSwagger
# And generate

hot:
	swag fmt
	swag init --parseDependency --parseInternal
	fresh

PHONY: run hot test test-all
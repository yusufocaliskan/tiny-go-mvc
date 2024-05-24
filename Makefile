
# Starts the app
run:
	go run main.go

# start the application in hotreload using fresh.yaml 
# Use it in dev mode
# FormatSwagger
# And generate

hot:
	swag fmt
	swag init --parseDependency --parseInternal
	fresh

.PHONY: run hot
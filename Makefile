
# Starts the app
run:
	go run main.go

# start the application in hotreload using fresh.yaml 
# Use it in dev mode
start hot:
	fresh

.PHONY: run start hot
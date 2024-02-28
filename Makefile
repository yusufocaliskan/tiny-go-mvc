
# Starts the app
run:
	go run main.go

# start the application in hotreload using fresh.yaml 
# Use it in dev mode
hot:
	fresh

.PHONY: run hot
sender:
	@ go run OBU/main.go

receiver:
	@go build -o bin/data_receiver ./data_receiver
	@./bin/receiver 

run:
	@ go run main.go
sender:
	@ go build -o bin/obu OBU/main.go
	@ ./bin/obu

receiver:
	@ go build -o bin/data_receiver ./data_receiver
	@ ./bin/data_receiver

calculator:
	@ go build -o bin/calculator ./distance_calculator
	@ ./bin/distance_calculator

run:
	@ go run main.go

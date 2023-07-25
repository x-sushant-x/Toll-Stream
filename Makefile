sender:
	@ go build -o bin/obu OBU/main.go
	@ ./bin/obu

receiver:
	@ go build -o bin/data_receiver ./data_receiver
	@ ./bin/data_receiver

calculator:
	@ go build -o bin/calculator ./distance_calculator
	@ ./bin/calculator

invoicer_aggregator:
	@ go build -o bin/invoicer_aggregator ./invoicer_aggregator
	@ ./bin/invoicer_aggregator

proto:
	@ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto

.PHONY: invoicer_aggregator
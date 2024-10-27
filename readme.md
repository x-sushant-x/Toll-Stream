# Real Time Toll Calculator - Microservices

### Introduction

The toll tax calculator is designed to provide real-time toll calculations based on data received from On-Board Units (OBUs). The system is implemented using microservices architecture, where each service performs a specific task and communicates through Kafka for data exchange. The toll tax calculation results are stored in MongoDB, and the invoicer generates invoices for the tolls incurred.

### Installation and Setup
To run the toll tax calculator system, follow these steps:
1. Clone the repository from GitHub:
```
git clone https://github.com/sushant102004/Real-Time-Toll-Calculator-Microservices/
```
2. Ensure you have Kafka and MongoDB set up and running.
3. Run each service using the provided make commands as mentioned below.

```
make sender
make receiver
make calculator
make aggregator
make invoicer
```

### Services
#### 1. OBU (On-Board Unit)

The OBU service is responsible for collecting toll-related data from vehicles passing through toll booths. It communicates with the Receiver service to send the collected data.

#### 2. Receiver

The Receiver service receives toll-related data from the OBUs and forwards it to the Calculator service via Kafka for processing.

#### 3. Calculator

The Calculator service calculates the toll tax based on the received data and stores the results in the Database Aggregator service.

#### 4. Database Aggregator

The Database Aggregator service aggregates the toll tax calculation results received from the Calculator service and stores them in MongoDB for further processing.

#### 5. Invoicer

The Invoicer service generates invoices based on the toll tax calculation results stored in MongoDB and provides a way to bill the customers.



### How It Works

1. The OBU service collects toll-related data from vehicles passing through toll booths.
2. The Receiver service receives the data from the OBUs and forwards it to the Calculator service via Kafka.
3. The Calculator service processes the received data and calculates the toll tax.
4. The Database Aggregator service stores the toll tax calculation results in MongoDB for further processing and analysis.
5. The Invoicer service generates invoices based on the toll tax calculation results stored in MongoDB and allows billing the customers.

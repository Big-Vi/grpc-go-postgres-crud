## Install grpc-go

`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`  
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`    


## Create proto files

`protoc --go_out=. --go-grpc_out=. orderpb/order.proto`  


## Atlas schema migration

`atlas schema inspect -u "postgres://postgres:root@localhost:3002/realtime-dashboard-grpc?sslmode=disable" > migrations/schema.hcl`  
`atlas schema apply -u "postgres://postgres:root@localhost:3002/realtime-dashboard-grpc?sslmode=disable" --to file://migrations/schema.hcl`  


## Run server

`docker compose -f docker-compose.yml up` 


## Run client

`go run client/client.go`  

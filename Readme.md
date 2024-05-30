Welcome to the Go gRPC User Service! This service provides functionalities for managing user details and includes a search capability.

##To run the application

1.go run main.go

#To run the test

1.go test ./service


#To test the service locally, we can use a gRPC client grpcurl.

#install grpcurl

 go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

#To Access the gRPC service endpoints.

1.grpcurl -plaintext -d '{\"id\": 1}' localhost:50051 user.UserService/GetUser
2.grpcurl -plaintext -d '{\"city\": \"LA\"}' localhost:50051 user.UserService/SearchUsers




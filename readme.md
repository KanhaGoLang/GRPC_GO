** Info **
https://www.youtube.com/watch?v=a6G5-LUlFO4&t=3228&ab_channel=AkhilSharma

protoc --go_out=. --go-grpc_out=. proto/greet.proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

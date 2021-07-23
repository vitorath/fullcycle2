protoc --proto_path=proto proto/*.proto --go_out=pb 
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

apt-get install libprotobuf-dev protobuf-compiler

evans -r repl --host localhost --port 50051
> service UserSevice
> call AddUser
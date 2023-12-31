codegen:
	protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import  proto/*.proto
	grpc_tools_ruby_protoc -I . --ruby_out=rb --grpc_out=rb proto/*.proto

winbuild:
	GOOS=windows GOARCH=amd64 go build -o main.exe

server:
	go run main.go --foobarPath=foobar.exe

clgo:
	go run client/main.go

clgo-build:
	go build -o client/client client/main.go

clrb:
	ruby rb/client/notify.rb

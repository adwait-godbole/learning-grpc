clean:
	rm pb/*.go
	
gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

evans-client:
	evans -r repl -p 8080

test:
	go test -v -cover -race ./...

.PHONY: clean gen server client evans-client test

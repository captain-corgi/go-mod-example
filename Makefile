build:
	go build -o go-mod-example main.go
run:
	./go-mod-example
clean:
	rm go-mod-example
run-client:
	go run cmd/client/main.go
BINARY_NAME=firewall-proxy

build:
	go build -o bin/$(BINARY_NAME) cmd/proxy/main.go

run:
	go run cmd/proxy/main.go

clean:
	rm -rf bin/

docker-build:
	docker build -t ebpf-llm-firewall -f docker/Dockerfile .
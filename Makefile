redis:
	docker run -d -p 6379:6379 redis redis

server:
	go run cmd/app/main.go

test:
	go test cmd/test/stress_test.go
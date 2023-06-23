build:
	go build -o main main.go

run:
	go run main.go

exec:
	go build -o main main.go && ./main

clean:
	rm -f ./main

.PHONY: build run exec clean
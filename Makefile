.PHONY: build run clean

build:
	go build -o service-a ./service-a
	go build -o service-b ./service-b

run:
	./service-a &
	./service-b

clean:
	rm -f service-a service-b


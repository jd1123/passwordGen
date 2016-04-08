build:
	go build

clean:
	rm passwordGen

all:
	make clean
	make build

build:
	mkdir -p functions
	go build -o functions/compute src/main.go src/equations.go
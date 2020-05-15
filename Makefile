build:
	mkdir -p functions
	go get "github.com/aws/aws-lambda-go/events"
	go get "github.com/aws/aws-lambda-go/lambda"
	go build -o functions/compute src/main.go src/equations.go
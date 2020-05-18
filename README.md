# <img src="https://github.com/karen-poon/econ-calculator/blob/master/site/favicon.ico" width="25"> ECE472 Calculator

Site has been deployed: https://ece472-calculator.netlify.app

This website is dedicated for students taking ECE472 Engineering Economic Analysis & Entrepreneurship at the University of Toronto.

You can use this website to find your desired interest factors for discrete compounding. 
If you just hate flipping through your thick textbook and tables with overcrowded numbers,
this website got you covered!

![website look](https://github.com/karen-poon/econ-calculator/blob/master/look.png)

## Mechanisms
To increase the complexity of this project, I added a backend system to this website. It is written in Go and powered by AWS Lambda,
which is a serverless compute service that handles HTTP requests via Amazon API Gateway. 

When user clicks on the "Submit" button, it will send an HTTP request to Lambda. 
After going through a handler written by me, it will provide a response and output it onto the website.
I am so glad that Netlify provides deployments involving Lambda without having to create an AWS account (since I don't have one either haha).
I will provide a flow diagram here later.

To make it work on Netlify, I have to compile and create an executable of my Go files and save it under the "functions" folder.
I decided to do this on Netlify's server. I tried to do it locally, but it didn't work. I am using Windows, and Lambda reads with Linux.
Loading a Windows executable on a Linux environment will make my deployment fail. Luckily, Netlify allows me to build on their server
with my own `Makefile` by enabling settings on a `netlify.toml` file.

Makefile:
```Makefile
build:
	mkdir -p functions
	go get "github.com/aws/aws-lambda-go/events"
	go get "github.com/aws/aws-lambda-go/lambda"
	go build -o functions/compute src/main.go src/equations.go
```

netlify.toml:
```netlify.toml
[build]
    command = "make build"
    functions = "functions"
    publish = "site"
[build.environment]
    GO_IMPORT_PATH = "github.com/karen-poon/econ-calculator"
```

## My advice
Although this website makes things convenient, please still take some time to actually flip through your textbook, 
since you are only allowed to bring your textbook and calculator to the quizzes and exam.
The more practice you have on flipping through the textbook, the quicker you find the desired interest factors!

Have fun using this website! :)

package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type responseMessage struct {
	Result     float64 `json:"result"`
	TypeStatus string  `json:"typeStatus"`
	IStatus    string  `json:"iStatus"`
	JStatus    string  `json:"jStatus"`
	NStatus    string  `json:"nStatus"`
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	typeInt, typeErr := strconv.Atoi(request.QueryStringParameters["type"])
	iFloat, iErr := strconv.ParseFloat(request.QueryStringParameters["i"], 64)
	nFloat, nErr := strconv.ParseFloat(request.QueryStringParameters["j"], 64)

	jsonMessage := responseMessage{
		Result:     0,
		TypeStatus: "OK",
		IStatus:    "OK",
		JStatus:    "OK",
		NStatus:    "OK",
	}

	if typeErr != nil {
		jsonMessage.TypeStatus = "error"
	}
	if iErr != nil {
		jsonMessage.IStatus = "error"
	}
	if nErr != nil {
		jsonMessage.NStatus = "error"
	}

	var result float64

	if typeInt == 9 || typeInt == 10 {
		jFloat, jErr := strconv.ParseFloat(request.QueryStringParameters["j"], 64)
		if jErr != nil {
			jsonMessage.JStatus = "error"
		} else if typeErr == nil && iErr == nil && nErr == nil && jErr == nil {
			switch typeInt {
			case 9:
				result = PGivenAWithJ(iFloat, jFloat, nFloat)
			case 10:
				result = FGivenAWithJ(iFloat, jFloat, nFloat)
			}
			jsonMessage.Result = result
		}

	} else if typeErr == nil && iErr == nil && nErr == nil {
		switch typeInt {
		case 1:
			result = PGivenF(iFloat, nFloat)
		case 2:
			result = FGivenP(iFloat, nFloat)
		case 3:
			result = PGivenA(iFloat, nFloat)
		case 4:
			result = AGivenP(iFloat, nFloat)
		case 5:
			result = FGivenA(iFloat, nFloat)
		case 6:
			result = AGivenF(iFloat, nFloat)
		case 7:
			result = PGivenG(iFloat, nFloat)
		case 8:
			result = AGivenG(iFloat, nFloat)
		}
		jsonMessage.Result = result
	}

	marshalledJSONMsg, _ := json.Marshal(jsonMessage)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(marshalledJSONMsg),
	}, nil

}

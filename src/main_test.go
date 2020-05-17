package main

import (
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

type testResponse struct {
	request  events.APIGatewayProxyRequest
	expected events.APIGatewayProxyResponse
}

var testSuccessWithoutJ = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "1",
			"i":    "0.1",
			"j":    "",
			"n":    "5",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"0.6209\",\"typeStatus\":\"OK\",\"iStatus\":\"OK\",\"jStatus\":\"OK\",\"nStatus\":\"OK\"}",
	},
}

var testSuccessWithJ = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "9",
			"i":    "0.1",
			"j":    "-0.2",
			"n":    "5",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"2.6551\",\"typeStatus\":\"OK\",\"iStatus\":\"OK\",\"jStatus\":\"OK\",\"nStatus\":\"OK\"}",
	},
}

var testErrorWithoutI = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "1",
			"i":    "",
			"j":    "",
			"n":    "5",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"error\",\"typeStatus\":\"OK\",\"iStatus\":\"error\",\"jStatus\":\"OK\",\"nStatus\":\"OK\"}",
	},
}

var testErrorWithoutJ = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "1",
			"i":    "",
			"j":    "",
			"n":    "5",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"error\",\"typeStatus\":\"OK\",\"iStatus\":\"OK\",\"jStatus\":\"error\",\"nStatus\":\"OK\"}",
	},
}

var testErrorWithoutN = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "9",
			"i":    "0.1",
			"j":    "0.2",
			"n":    "",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"error\",\"typeStatus\":\"OK\",\"iStatus\":\"OK\",\"jStatus\":\"OK\",\"nStatus\":\"error\"}",
	},
}

// should never happen
var testErrorWithoutType = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "",
			"i":    "0.1",
			"j":    "0.2",
			"n":    "5",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"error\",\"typeStatus\":\"error\",\"iStatus\":\"OK\",\"jStatus\":\"OK\",\"nStatus\":\"OK\"}",
	},
}

var testSuccessITooLarge = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "1",
			"i":    "99999",
			"j":    "",
			"n":    "5",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"0.0000\",\"typeStatus\":\"OK\",\"iStatus\":\"OK\",\"jStatus\":\"OK\",\"nStatus\":\"OK\"}",
	},
}

var testErrorTooLargeInf = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "5",
			"i":    "99999999",
			"j":    "",
			"n":    "99999999",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"error\",\"typeStatus\":\"OK\",\"iStatus\":\"OK\",\"jStatus\":\"OK\",\"nStatus\":\"OK\"}",
	},
}

var testErrorTooLargeNaN = testResponse{
	request: events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"type": "10",
			"i":    "99999999",
			"j":    "99999999",
			"n":    "99999999",
		},
	},
	expected: events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{\"result\":\"error\",\"typeStatus\":\"OK\",\"iStatus\":\"OK\",\"jStatus\":\"OK\",\"nStatus\":\"OK\"}",
	},
}

func TestSuccessWithoutJ(t *testing.T) {
	actualResponse, err := handler(testSuccessWithoutJ.request)
	if actualResponse.StatusCode != testSuccessWithoutJ.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testSuccessWithoutJ.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testSuccessWithoutJ.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testSuccessWithoutJ.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

func TestSuccessWithJ(t *testing.T) {
	actualResponse, err := handler(testSuccessWithJ.request)
	if actualResponse.StatusCode != testSuccessWithJ.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testSuccessWithJ.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testSuccessWithJ.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testSuccessWithJ.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

func TestErrorWithoutI(t *testing.T) {
	actualResponse, err := handler(testErrorWithoutI.request)
	if actualResponse.StatusCode != testErrorWithoutI.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testErrorWithoutI.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testErrorWithoutI.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testErrorWithoutI.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

func TestErrorWithoutJ(t *testing.T) {
	actualResponse, err := handler(testErrorWithoutJ.request)
	if actualResponse.StatusCode != testErrorWithoutJ.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testErrorWithoutJ.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testErrorWithoutI.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testErrorWithoutJ.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

func TestErrorWithoutN(t *testing.T) {
	actualResponse, err := handler(testErrorWithoutN.request)
	if actualResponse.StatusCode != testErrorWithoutN.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testErrorWithoutN.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testErrorWithoutN.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testErrorWithoutN.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

func TestErrorWithoutType(t *testing.T) {
	actualResponse, err := handler(testErrorWithoutType.request)
	if actualResponse.StatusCode != testErrorWithoutType.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testErrorWithoutType.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testErrorWithoutType.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testErrorWithoutType.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

func TestSuccessITooLarge(t *testing.T) {
	actualResponse, err := handler(testSuccessITooLarge.request)
	if actualResponse.StatusCode != testSuccessITooLarge.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testSuccessITooLarge.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testSuccessITooLarge.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testSuccessITooLarge.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

func TestErrorTooLargeInf(t *testing.T) {
	actualResponse, err := handler(testErrorTooLargeInf.request)
	if actualResponse.StatusCode != testErrorTooLargeInf.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testErrorTooLargeInf.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testErrorTooLargeInf.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testErrorTooLargeInf.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

func TestErrorTooLargeNaN(t *testing.T) {
	actualResponse, err := handler(testErrorTooLargeNaN.request)
	if actualResponse.StatusCode != testErrorTooLargeNaN.expected.StatusCode {
		t.Errorf("got %d, want %d", actualResponse.StatusCode, testErrorTooLargeNaN.expected.StatusCode)
	}
	if strings.Compare(actualResponse.Body, testErrorTooLargeNaN.expected.Body) != 0 {
		t.Errorf("got %s, want %s", actualResponse.Body, testErrorTooLargeNaN.expected.Body)
	}
	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}
}

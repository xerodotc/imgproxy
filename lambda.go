package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

type lambdaHandler struct{}

type labdaResponse struct {
	IsBase64Encoded bool              `json:"'isBase64Encoded'"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
}

func (s lambdaHandler) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	ctx = timerWithSince(ctx)

	ctx, err := parsePath(ctx, strings.TrimSpace(string(payload)), &processingHeaders{})
	if err != nil {
		return []byte{}, err
	}

	ctx, downloadcancel, err := downloadImage(ctx)
	defer downloadcancel()
	if err != nil {
		return []byte{}, err
	}

	checkTimeout(ctx)

	data, err := processImage(ctx)
	if err != nil {
		return []byte{}, err
	}

	po := getProcessingOptions(ctx)

	resp := labdaResponse{
		IsBase64Encoded: true,
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": mimes[po.Format]},
		Body:            base64.StdEncoding.EncodeToString(data),
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		return []byte{}, err
	}

	return jsonData, nil
}

func startLambdaServer() {
	lambda.StartHandler(new(lambdaHandler))
}

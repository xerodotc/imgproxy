package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/imgproxy/imgproxy/v2/obs"
)

type obsTransport struct {
	client *obs.ObsClient
}

func newOBSTransport() (http.RoundTripper, error) {
	client, err := obs.New(conf.OBSAccessKey, conf.OBSSecretKey, conf.OBSEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Can't create OBS client: %s", err)
	}

	return obsTransport{client}, nil
}

func (t obsTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	input := &obs.GetObjectInput{}
	input.Bucket = req.URL.Host
	input.Key = strings.TrimPrefix(req.URL.Path, "/")

	output, err := t.client.GetObject(input)

	if err != nil {
		return nil, err
	}

	header := make(http.Header)
	header.Set("Cache-Control", output.CacheControl)

	return &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.0",
		ProtoMajor:    1,
		ProtoMinor:    0,
		Header:        header,
		ContentLength: output.ContentLength,
		Body:          output.Body,
		Close:         true,
		Request:       req,
	}, nil
}

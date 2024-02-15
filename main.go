package main

import (
	"context"
	"log"
	"net/http"

	elizav1 "connect-examples-go/internal/gen/connectrpc/eliza/v1"
	"connect-examples-go/internal/gen/connectrpc/eliza/v1/elizav1connect"

	"connectrpc.com/connect"
)

func main() {
	log.SetFlags(0)
	client := elizav1connect.NewElizaServiceClient(
		&http.Client{
			Transport: &transport{},
		},
		"http://localhost:8082",
	)
	res, err := client.Say(
		context.Background(),
		connect.NewRequest(&elizav1.SayRequest{
			Sentence: "Hey",
		}),
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg.Sentence)
	sres, err := client.Introduce(context.Background(), connect.NewRequest(&elizav1.IntroduceRequest{
		Name: "John",
	}))
	if err != nil {
		log.Fatalln(err)
	}
	_ = sres.Close()
}

type transport struct{}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	return (&http.Transport{}).RoundTrip(req)
}

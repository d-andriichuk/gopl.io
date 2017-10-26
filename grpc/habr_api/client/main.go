package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	pb "gopl.io/grpc/habr_api/habrapi"
)

const (
	addr = "localhost:50051"
)

var username = flag.String("user", "mad_hero", "User name")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewHabrApiClient(conn)

	karmaResp, err := c.GetKarma(context.Background(), &pb.KarmaRequest{Username: *username})
	if err != nil {
		log.Fatalf("Could not get karma %v\n", err)
	}
	log.Printf("User %s have %f karma", karmaResp.Username, karmaResp.Karma)

	articleResp, err := c.PostArticle(context.Background(), &pb.PostArticleRequest{
		Title: "title",
		Body:  "body",
		Tag:   []string{"tag1", "tag2"},
		Hub:   []string{"hub1", "hub2"},
	})

	if articleResp.Posted {
		log.Printf("Article was posted at %s with error %s\n", articleResp.Time, articleResp.ErrorCode)
		log.Printf("Your article live here %s", articleResp.Url)
	} else {
		log.Printf("Artciel was not posted at %s with error %s\n", articleResp.Time, articleResp.ErrorCode)
	}
}

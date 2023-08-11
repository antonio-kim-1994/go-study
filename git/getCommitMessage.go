package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v53/github"
	"github.com/rs/zerolog/log"
)

var (
	sourceOwner = ""
	sourceRepo  = ""
	branch      = ""
)

var client *github.Client
var ctx = context.Background()

func main() {
	token := ""
	client = github.NewTokenClient(ctx, token)

	ref, _, err := client.Git.GetRef(ctx, sourceOwner, sourceRepo, "refs/heads/"+branch)
	if err != nil {
		log.Err(err).Msg("")
	}

	commit, _, err := client.Git.GetCommit(ctx, sourceOwner, sourceRepo, *ref.Object.SHA)
	if err != nil {
		log.Err(err).Msg("")
	}

	fmt.Println(*commit.Message)
}

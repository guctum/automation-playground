package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/google/go-github/v35/github"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/oauth2"
)

func HandleRequest(ctx context.Context) {
	createRepo(*name, *private)
}

func main()  {
	lambda.Start(HandleRequest)
}

func boolCheck(privacy string) bool {
	var privacyCheck bool
	if privacy == "true" {
		privacyCheck = true
	} else if privacy == "false" {
		privacyCheck = false
	}
	return privacyCheck
}

var (
	name        = flag.String("name", os.Args[1], "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "", "Description of created repo.")
	private     = flag.Bool("private", boolCheck(os.Args[2]), "Whether or not the created repo will be private")
)

func createRepo(repoName string, repoPrivacy bool) {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("github-token") // this needs to be a Personal Access Token from GitHub with at least repo permissions set
	if token == "" {
		log.Fatal("Unauthorized: No Token Present")
	}
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	oauthClient := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(oauthClient)
	repository := &github.Repository{Name: name, Private: private}
	repo, _, err := client.Repositories.Create(ctx, "", repository)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func main()  {
	fmt.Println("Now creating a new github repo named: ", os.Args[1])
	setupRepo()
}

var (
	name        = flag.String("name", os.Args[1], "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "", "Description of created repo.")
	private     = flag.Bool("private", false, "Will created repo be private.")
)

func setupRepo() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("github-token") // this needs to be a Personal Access Token from GitHub with at least repo permissions set
	if token == "" {
		log.Fatal("Unauthorized: No Token Present")
	}
	if *name == "" {
		log.Fatal("Script needs a repo name")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	r := &github.Repository{Name: name, Private: private, Description: description}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}
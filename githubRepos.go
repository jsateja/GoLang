package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/go-github/v43/github"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the GitHub user: ")
	username, _ := r.ReadString('\n')

	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := client.Repositories.List(ctx, strings.TrimRight(username,"\n"), opt)


	if err != nil {
		fmt.Println(err)
	}
	for _,v := range repos {
		fmt.Println(*v.Name)
	}
}
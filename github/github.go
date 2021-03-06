package github

import (
  "context"
  "log"

  "github.com/aytdm/github-playground/config"
  "github.com/google/go-github/v33/github"
  "golang.org/x/oauth2"
)

type Issues struct {
  Issues []Issue `json:"issues"`
}

type Issue struct {
  Title string   `json:"title"`
  Body string   `json:"body"`
  Labels []string `json:"labels"`
}

var (
  authClient *github.Client
  conf *config.Config
)

func GetIssues() {
  client, err := getClient()
  if err != nil {
    log.Fatal("cannot get client: ", err)
    return
  }

  opts := &github.IssueListByRepoOptions{
    Milestone: "none",
    Direction: "asc",
    ListOptions: github.ListOptions{
      PerPage: 50,
    },
  }
  
  log.Println("existed issues list:")

  issues, _, err := client.Issues.ListByRepo(context.Background(), conf.Github.Owner, conf.Github.Repository, opts)
  if err != nil {
    log.Fatal("cannot list issue:", err)
    return
  }

  for _, issue := range issues {
    log.Printf("No: %02d, title: %s\n", *issue.Number, *issue.Title)
  }
}

func CreateIssues(issues *Issues) {
  client, err := getClient()
  if err != nil {
    log.Fatal("cannot get client: ", err)
    return
  }

  result := make([]int64, 0)

  for i := 0; i < len(issues.Issues); i++ {
    newIssueRequest := github.IssueRequest{}
    newIssueRequest.Title = &issues.Issues[i].Title
    newIssueRequest.Body = &issues.Issues[i].Body
    if (len(issues.Issues[i].Labels) > 0) {
      newIssueRequest.Labels = &issues.Issues[i].Labels
    }

    newIssue, resp, err := client.Issues.Create(context.Background(), conf.Github.Owner, conf.Github.Repository, &newIssueRequest)
    if err != nil {
      log.Println(resp)
      log.Fatal("Failed to request issues: %s", err.Error())
      return
    }
    result = append(result, *newIssue.ID)
  }

  log.Printf("created issue nums: %d\n", len(result))
  return
}

func init() {
  if conf == nil {
    config, err := config.LoadConfig()
    if err != nil {
      log.Fatal("cannot load config: ", err)
    }
    conf = config
  }
}

func getClient() (client *github.Client, err error) {
  if authClient != nil {
    client = authClient
    return
  }

  ctx := context.Background()
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: conf.Github.Token})
  tc := oauth2.NewClient(ctx, ts)
  client = github.NewClient(tc)
  return
}

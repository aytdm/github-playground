package service

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "os"

  "github.com/aytdm/createissue/github"
)

func Start() {
  github.CreateIssues(getIssuesFromJson())

  github.GetIssues()
}

func getIssuesFromJson() (issues *github.Issues) {
  if issues != nil {
    return issues
  }

  jsonFile, err := os.Open("issues.json")
  if err != nil {
    log.Fatal(err)
  }
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)
  issueInstance := github.Issues{}
  json.Unmarshal(byteValue, &issueInstance)
  issues = &issueInstance

  return
}

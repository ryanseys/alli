package main

import (
    "bufio"
    "os"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "encoding/json"
)

const delim = '\n'

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

  var token string = os.Getenv("GH_TOKEN_ALLI") // env var name
  var home string = os.Getenv("HOME") // env var name

  if token == "" {
    fmt.Printf("No token found! Using highly rate-limited public access.\n")
  } else {
    fmt.Printf("Yay! Using authentication token!\n")
  }

  r := bufio.NewReader(os.Stdin)

  var username string
  saved_username, err := ioutil.ReadFile(home + "/.alli")

  if string(saved_username) == "" {
    print("Enter Github username: ")
    username, err = r.ReadString(delim)
    check(err)
  } else {
    username = string(saved_username)
    fmt.Printf("Using saved username: %s\n", username)
  }

  username = strings.TrimSpace(username) // trim newline

  // write username to save
  d1 := []byte(username)
  ioutil.WriteFile(home + "/.alli", d1, 0644)

  client := &http.Client{}
  req, _ := http.NewRequest("GET", "https://api.github.com/users/" + username + "/repos?per_page=100", nil)
  req.Header.Set("Accept", "application/vnd.github.v3+json")

  if token != "" {
    req.SetBasicAuth(token, "x-oauth-basic") // user, password
  }

  repos, err := client.Do(req)
  check(err)

  defer repos.Body.Close()
  contents, err := ioutil.ReadAll(repos.Body);
  check(err)

  byt := []byte(contents)

  var f interface{}
  err = json.Unmarshal(byt, &f)
  check(err)

  array := f.([]interface {})

  println()

  for i := range array {
    repo := array[i].(map[string]interface {})
    var countFloat float64 = repo["open_issues_count"].(float64)
    var countInt int = int(countFloat)
    if(countInt != 0) {
      var name string = repo["full_name"].(string)
      fmt.Printf("%s\n", name)

      req, _ = http.NewRequest("GET", "https://api.github.com/repos/" + name + "/issues?state=open", nil)
      req.Header.Set("Accept", "application/vnd.github.v3+json")
      if token != "" {
        req.SetBasicAuth(token, "x-oauth-basic") // user, password
      }
      issues, err := client.Do(req)
      check(err)

      defer issues.Body.Close()
      iss, err := ioutil.ReadAll(issues.Body);
      check(err)

      byt = [] byte(iss)
      var g interface{}
      err = json.Unmarshal(byt, &g)
      check(err)

      issue_array := g.([]interface {})

      for j := range issue_array {
        issue := issue_array[j].(map[string]interface {})
        var number int = int(issue["number"].(float64))
        var title string = issue["title"].(string)
        fmt.Printf("#%d %s\n", number, title)
      }
      fmt.Printf("\n");
    }
  }
}

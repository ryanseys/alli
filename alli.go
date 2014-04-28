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

func main() {

  r := bufio.NewReader(os.Stdin)

  print("Enter Github username: ")
  line, err := r.ReadString(delim)

  if err != nil {
    fmt.Println("Error occured: %s\n", err)
  }

  line = strings.TrimSpace(line)
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "https://api.github.com/users/" + line + "/repos?per_page=100", nil)
  req.Header.Set("Accept", "application/vnd.github.v3+json")
  repos, err := client.Do(req)

  if err != nil {
    fmt.Println("Error occured: %s\n", err)
  }

  defer repos.Body.Close()
  contents, err := ioutil.ReadAll(repos.Body);

  if err != nil {
    fmt.Println("Error occured: %s\n", err)
  }

  byt := []byte(contents)

  var f interface{}
  err = json.Unmarshal(byt, &f)

  if err != nil {
    fmt.Println("Error occured: %s\n", err)
  }

  array := f.([]interface {})

  for i := range array {
    repo := array[i].(map[string]interface {})
    var countFloat float64 = repo["open_issues_count"].(float64)
    var countInt int = int(countFloat)
    if(countInt != 0) {
      fmt.Println(repo["full_name"])
    }
  }
}

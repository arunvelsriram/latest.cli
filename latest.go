package main

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "time"
  "errors"

  "github.com/urfave/cli"
)

type Gem struct {
  Name string
  Version string
}

func latestRubyGemVersion(gemName string) (Gem, error) {
  var gem Gem
  var httpClient = &http.Client{
    Timeout: time.Second * 10,
  }

  url := fmt.Sprintf("https://rubygems.org//api/v1/gems/%s.json", gemName)
  resp, err := httpClient.Get(url)
  if err != nil {
    return gem, errors.New("Failed to fetch gem details")
  }

  if(resp.StatusCode == http.StatusNotFound) {
    return gem, errors.New("Gem not found")
  }

  jsonBlob, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    return gem, errors.New("Failed to read gem details")
  }

  err = json.Unmarshal(jsonBlob, &gem)
  if err != nil {
    return gem, errors.New("Failed to parse gem details JSON")
  }

  return gem, nil
}

func main() {
  var queryGem bool
  var queryNodePkg bool

  app := cli.NewApp()
  app.Name = "latest"
  app.Usage = "CLI to find the latest version of a ruby gem, node package etc."

  app.Flags = []cli.Flag {
    cli.BoolFlag {
      Name:  "gem, g",
      Usage: "query for latest version of a ruby gem",
      Destination: &queryGem,
    },
    cli.BoolFlag {
      Name: "node-pkg, n",
      Usage: "query for latest version of a node package",
      Destination: &queryNodePkg,
    },
  }

  app.Action = func(context *cli.Context) error {
    if queryGem {
      gemName := context.Args().Get(0)

      gem, err := latestRubyGemVersion(gemName)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      fmt.Println("Gem found:")
      fmt.Println(gem.Name, gem.Version)
    } else if queryNodePkg {
      fmt.Println("node")
    } else {
      fmt.Println("gem")
      fmt.Println("node")
    }
    return nil
  }

  app.Run(os.Args)
}

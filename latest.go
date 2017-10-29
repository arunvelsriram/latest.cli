package main

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "time"
  "log"

  "github.com/urfave/cli"
)

type Gem struct {
  Name string
  Version string
}

func latestRubyGemVersion(gemName string) Gem {
  var gem Gem
  var httpClient = &http.Client{
    Timeout: time.Second * 10,
  }

  url := fmt.Sprintf("https://rubygems.org//api/v1/gems/%s.json", gemName)
  resp, err := httpClient.Get(url)
  if err != nil {
    log.Fatal(err)
  }

  if(resp.StatusCode == http.StatusNotFound) {
    return gem
  }

  jsonBlob, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(jsonBlob, &gem)
  if err != nil {
    log.Print(gem)
    log.Fatal(err)
  }

  return gem
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
      fmt.Println("gem")
      gemName := context.Args().Get(0)
      gem := latestRubyGemVersion(gemName)
      fmt.Printf("%s: %s\n", gem.Name, gem.Version)
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

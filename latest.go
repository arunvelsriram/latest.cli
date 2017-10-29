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

type Package struct {
  Name string
  Version string
}

func fetch(url string) (Package, error) {
  var pkg Package
  var httpClient = &http.Client{
    Timeout: time.Second * 10,
  }

  resp, err := httpClient.Get(url)
  if err != nil {
    return pkg, errors.New(fmt.Sprintf("Failed to get response from %s", url))
  }

  if(resp.StatusCode == http.StatusNotFound) {
    return pkg, errors.New("Not found")
  }

  jsonBlob, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    return pkg, errors.New("Failed to read response body")
  }

  err = json.Unmarshal(jsonBlob, &pkg)
  if err != nil {
    return pkg, errors.New("Failed to parse JSON")
  }

  return pkg, nil
}

func latestRubyGem(name string) (Package, error) {
  var pkg Package

  pkg, err := fetch(fmt.Sprintf("https://rubygems.org/api/v1/gems/%s.json", name))
  if err != nil {
    return pkg, err
  }

  return pkg, nil
}

func latestNodePackage(name string) (Package, error) {
  var pkg Package

  pkg, err := fetch(fmt.Sprintf("https://registry.npmjs.org/%s/latest", name))
  if err != nil {
    return pkg, err
  }

  return pkg, nil
}

func main() {
  var isRubyGem bool
  var isNodePackage bool

  app := cli.NewApp()
  app.Name = "latest"
  app.Usage = "CLI to find the latest version of a ruby gem, node package etc."

  app.Flags = []cli.Flag {
    cli.BoolFlag {
      Name:  "gem, g",
      Usage: "query for latest version of a ruby gem",
      Destination: &isRubyGem,
    },
    cli.BoolFlag {
      Name: "node-pkg, n",
      Usage: "query for latest version of a node package",
      Destination: &isNodePackage,
    },
  }

  app.Action = func(context *cli.Context) error {
    if isRubyGem {
      name := context.Args().Get(0)

      gem, err := latestRubyGem(name)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      fmt.Println("Gem found:")
      fmt.Println(gem.Name, gem.Version)
    } else if isNodePackage {
      name := context.Args().Get(0)

      nodePackage, err := latestNodePackage(name)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      fmt.Println("Node package found:")
      fmt.Println(nodePackage.Name, nodePackage.Version)
    } else {
      fmt.Println("gem")
      fmt.Println("node")
    }
    return nil
  }

  app.Run(os.Args)
}

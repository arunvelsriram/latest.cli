package main

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "time"

  "github.com/urfave/cli"
)

type NotFoundError struct {
  Message string
}

func (error *NotFoundError) Error() (string) {
  return error.Message
}

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
    return pkg, fmt.Errorf("Failed to get response from %s", url)
  }

  if(resp.StatusCode == http.StatusNotFound) {
    return pkg, &NotFoundError{"Not found!"}
  }

  jsonBlob, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    return pkg, fmt.Errorf("Failed to read response body")
  }

  err = json.Unmarshal(jsonBlob, &pkg)
  if err != nil {
    return pkg, fmt.Errorf("Failed to parse JSON")
  }

  return pkg, nil
}

func latestRubyGem(name string) (error) {
  fmt.Println("Ruby Gem:")

  gem, err := fetch(fmt.Sprintf("https://rubygems.org/api/v1/gems/%s.json", name))
  if err != nil {
    return err
  }

  fmt.Println(gem.Name, gem.Version)

  return nil
}

func latestNodePackage(name string) (error) {
  fmt.Println("Node Module:")

  nodePackage, err := fetch(fmt.Sprintf("https://registry.npmjs.org/%s/latest", name))
  if err != nil {
    return err
  }

  fmt.Println(nodePackage.Name, nodePackage.Version)
  return nil
}

func reportError(err error) {
  fmt.Fprintln(os.Stderr, err)
}

func exitStatus(err error) (int) {
  if _, ok := err.(*NotFoundError); ok {
    return 0
  }

  return 1
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
    name := context.Args().Get(0)
    fmt.Println()

    if isRubyGem {
      err := latestRubyGem(name)
      if err != nil {
        reportError(err)
        os.Exit(exitStatus(err))
      }
    } else if isNodePackage {
      err := latestNodePackage(name)
      if err != nil {
        reportError(err)
        os.Exit(exitStatus(err))
      }
    } else {
      fmt.Println("gem")
      fmt.Println("node")
    }
    return nil
  }

  app.Run(os.Args)
}
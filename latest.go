package main

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "time"
  "sort"
  "strings"

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
  httpClient := &http.Client{
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
  fmt.Println()
  fmt.Println("Ruby Gem:")

  gem, err := fetch(fmt.Sprintf("https://rubygems.org/api/v1/gems/%s.json", name))
  if err != nil {
    return err
  }

  fmt.Println(gem.Name, gem.Version)

  return nil
}

func latestNodeModule(name string) (error) {
  fmt.Println()
  fmt.Println("Node Module:")

  nodeModule, err := fetch(fmt.Sprintf("https://registry.npmjs.org/%s/latest", name))
  if err != nil {
    return err
  }

  fmt.Println(nodeModule.Name, nodeModule.Version)
  return nil
}

func latestAll(name string) ([]error) {
  var errs []error

  if err := latestRubyGem(name); err != nil {
    fmt.Fprintln(os.Stderr, err)
    errs = append(errs, err)
  }

  if err := latestNodeModule(name); err != nil {
    fmt.Fprintln(os.Stderr, err)
    errs = append(errs, err)
  }

  return errs
}

func isEmpty(content string) bool {
  if strings.TrimSpace(content) == "" {
    return true
  }

  return false
}

func exitStatus(err error) (int) {
  if _, ok := err.(*NotFoundError); ok {
    return 0
  }

  return 1
}

func main() {
  app := cli.NewApp()
  app.Name = "latest"
  app.Usage = "A CLI to find the latest version of a Ruby Gem, Node module, Java JAR etc."

  before :=  func(cliContext *cli.Context) (error) {
    name := cliContext.Args().Get(0)
    if ok := isEmpty(name); ok {
      err := fmt.Errorf("name not given")
      return cli.NewExitError(err, exitStatus(err))
    }

    return nil
  }

  app.Commands = []cli.Command {
    {
      Name:  "gem",
      Aliases: []string{"g"},
      Usage: "get latest version of ruby gem <name>",
      ArgsUsage: "<name>",
      Before: before,
      Action: func(cliContext *cli.Context) error {
        name := cliContext.Args().Get(0)
        err := latestRubyGem(name)
        if err != nil {
          return cli.NewExitError(err, exitStatus(err))
        }

        return nil
      },
    },
    {
      Name: "node-module",
      Aliases: []string{"n"},
      Usage: "get latest version of node module <name>",
      ArgsUsage: "<name>",
      Before: before,
      Action: func(cliContext *cli.Context) error {
        name := cliContext.Args().Get(0)
        err := latestNodeModule(name)
        if err != nil {
          return cli.NewExitError(err, exitStatus(err))
        }

        return nil
      },
    },
    {
      Name: "all",
      Aliases: []string{"a"},
      Usage: "get latest version of <name>",
      ArgsUsage: "<name>",
      Before: before,
      Action: func(cliContext *cli.Context) error {
        name := cliContext.Args().Get(0)
        if errs := latestAll(name); errs != nil {
          var computedExitStatus int

          for _, err := range errs {
            computedExitStatus += exitStatus(err)
          }

          os.Exit(computedExitStatus)
        }

        return nil
      },
    },
  }

  sort.Sort(cli.CommandsByName(app.Commands))

  app.Run(os.Args)
}
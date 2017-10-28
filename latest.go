package main


import (
  "fmt"
  "os"
  "github.com/urfave/cli"
)

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

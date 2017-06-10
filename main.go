package main

import (
  "bufio"
  "os"
  "fmt"
  "crypto/tls"

  "gopkg.in/resty.v0"
  "github.com/danielhood/loco.cli/config"
  "github.com/danielhood/loco.cli/clientApis"
)

func main() {
  fmt.Printf("loco.cli v%v starting\n", config.Version())

  resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })

  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')

    switch (text) {
    case "quit\n":
      os.Exit(0)
    case "help\n":
      showHelp()
    case "objects\n":
      getObjects()
    default:
      fmt.Printf("Unknown command: %vType help for supported commands.\n", text)
    }
  }
}

func showHelp() {
  fmt.Print("quit\t\tTerminates application\n")
  fmt.Print("help\t\tDisplays a list of commands\n")
  fmt.Print("objects\t\tList status of all objects\n")
}

func getObjects() {
  //fmt.Print("Getting token from server: ", config.LocoServer)

  token, err := clientApis.GetToken()
  if (err != nil) {
    fmt.Printf("Unable to get token: %v\n", err)
    return
  }

  //fmt.Print("Got token: ", token)

  //fmt.Print("Getting object list")
  objects, err := clientApis.GetObjects(token)

  for _, o := range objects {
      fmt.Printf("%v(%v): %v @ (%v, %v)\n", o.Id, o.Type, o.Name, o.X, o.Y)
  }

}

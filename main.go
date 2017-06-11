package main

import (
  "bufio"
  "os"
  "errors"
  "fmt"
  "time"
  "crypto/tls"

  "gopkg.in/resty.v0"
  "github.com/danielhood/loco.cli/config"
  "github.com/danielhood/loco.cli/clientApis"
)

func main() {
  fmt.Printf("loco.cli v%v starting\n", config.Version())

  resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })

  token, err := doLogin()
  if (err != nil) {
    os.Exit(1)
  }

  // Setup backround process
  ticker := time.NewTicker(5 * time.Second)
  quit := make(chan struct{})
  go func() {
      for {
         select {
          case <- ticker.C:
              performSync()
          case <- quit:
              ticker.Stop()
              fmt.Print("Exiting sync loop\n")
              return
          }
      }
   }()

  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')

    switch (text) {
    case "quit\n":
      close (quit)
      os.Exit(0)
    case "help\n":
      showHelp()
    case "objects\n":
      getObjects(token)
    default:
      fmt.Printf("Unknown command: %vType help for supported commands.\n", text)
    }
  }
}

func performSync() {
  //fmt.Print("synching...\n")
}

func showHelp() {
  fmt.Print("quit\t\tTerminates application\n")
  fmt.Print("help\t\tDisplays a list of commands\n")
  fmt.Print("objects\t\tList status of all objects\n")
}

func doLogin() (string, error) {
  //fmt.Print("user: ")
  //user, _ := reader.ReadString('\n')

  token, err := clientApis.GetToken()
  if (err != nil) {
    fmt.Printf("Login failed. Unable to get token: %v\n", err)
    return "", errors.New("Unable to get token")
  }

  return token, nil
}

func getObjects(token string) {
  objects, _ := clientApis.GetObjects(token)

  for _, o := range objects {
      fmt.Printf("%v(%v): %v @ (%v, %v)\n", o.Id, o.Type, o.Name, o.X, o.Y)
  }

}

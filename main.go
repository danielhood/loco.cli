package main

import (
  "log"
  "crypto/tls"

  "gopkg.in/resty.v0"
  "github.com/danielhood/loco.cli/config"
  "github.com/danielhood/loco.cli/clientApis"
)

func main() {
  log.Printf("loco.cli v%v starting", config.Version)

  resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })

  log.Print("Getting token from server: ", config.LocoServer)

  token, err := clientApis.GetToken()
  if (err != nil) {
    log.Printf("Unable to get token: ", err)
    return
  }

  log.Print("Got token: ", token)

  log.Print("Getting object list")
  objects, err := clientApis.GetObjects(token)

  log.Print("Got Objects: ", objects)

  log.Print("Done")
}

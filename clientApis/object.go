package clientApis

import (
  "encoding/json"

  "gopkg.in/resty.v0"

  "github.com/danielhood/loco.cli/config"
  "github.com/danielhood/loco.server/entities"
)

func GetObjects(token string) ([]entities.Object, error) {
  resty.SetAuthToken(token)
  objectsResp, err := resty.R().Get(config.LocoServer() + "/object")

  if err != nil {
    return nil, err
  }

  var o []entities.Object
  err =  json.Unmarshal([]byte(objectsResp.String()), &o)

  if (err != nil) {
    return nil, err
  }

  return o, nil
}

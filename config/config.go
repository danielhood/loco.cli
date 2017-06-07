package config

const (
  server = "https://loco.local:8080"
  version = "0.1"
)

func LocoServer() string {
  return server
}

func Version() string {
  return version
}

package main

import (
  "encoding/json"
  "io/ioutil"
  "github.com/golang/glog"
)

type Config struct {
	CommandFile string `json:"commandFile"`
	ComPort string `json:"comPort"`
	Prompt string `json:"prompt"`
	Delay int `json:"delay"`
}

func ReadConfig(configfile string) Config {
  raw, err := ioutil.ReadFile(configfile)
  if err != nil {
    glog.Fatal(err)
  }

  var config Config
  json.Unmarshal(raw, &config)

  glog.Info(config)
  return config
}

/**func DisplayConfig() {
  fmt.Println(toJson(config))
}**/

package config
import (
    "encoding/json"
    "io/ioutil"
    "log"
)

var (
    Token     string
    Prefix string
    config    *configStruct
)

type configStruct struct {
    Token     string `json:"token"`
    Prefix string `json:"prefix"`
}

func ReadConfig() error {
    file, err := ioutil.ReadFile("./config.json")
    if err != nil {
        log.Fatal(err)
        return err
    }
    err = json.Unmarshal(file, &config)
    if err != nil {
        log.Fatal(err)
        return err
    }
    Token = config.Token
    Prefix = config.Prefix

    return nil
}
package config

import (
    "io/ioutil"
    "encoding/json"
)

type ConfigData struct {
    Request HttpRequest
    Concurrent ConcurrentSetting
}

type Header struct {
    Key string
    Value string
}

type HttpRequest struct {
    Url string
    ContentType string
    Method string
    Body string
}

type ConcurrentSetting struct {
    Threads int
    Times int
}

func LoadData() ConfigData {
    rawdata, err := ioutil.ReadFile("config/config.json")
    if(err != nil) {
        panic("error reading file: " + err.Error())
    }
    var configData ConfigData
    err = json.Unmarshal(rawdata, &configData)
    if(err != nil) {
        panic("error parsing json configuration: " + err.Error())
    }
    return configData
}
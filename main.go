package main

import (
    "goMeter/config"
    "goMeter/webServiceMeter"
)

func main() {
    conf := config.LoadData()
    webServiceMeter.Run(conf)
}
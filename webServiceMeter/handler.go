package webServiceMeter

import (
    "goMeter/config"
    "fmt"
    "goMeter/util"
    "goMeter/myHttp"
)

type executionResult struct {
    statusCode int
    executionTime int
    duration int
    err error
}

var conf config.ConfigData

func Run(configData config.ConfigData) {
    conf = configData
    c := sendRequests(conf)
    handleResult(c)
}

func sendRequests(conf config.ConfigData) chan executionResult{
    c := make(chan executionResult)
    for i:=0; i < conf.Concurrent.Threads; i++ {
        go util.Repeat(conf.Concurrent.Times, func(){
            var response myHttp.MyResponse
            var err error
            duration := util.InvokeAndGetTime(func() {
                response, err = myHttp.MakeRequest(conf.Request)
            })
            
            if err != nil {
                c <- executionResult{err: err, duration:duration}
            } else {
                c <- executionResult{statusCode: response.StatusCode(), duration:duration}
            }
        })
    }
    return c
}

func handleResult(c chan executionResult) {
    count := 0
    errorCount := 0
    requestNum := conf.Concurrent.Threads * conf.Concurrent.Times
    durationCount := 0
    for (count < requestNum) {
        select {
        case result := <- c:
            if result.err != nil {
                fmt.Println("error occured: " + result.err.Error())
                errorCount++
            } else {
                durationCount += result.duration
                fmt.Printf("%d / %d response received, statusCode: %d\n", count+1, requestNum, result.statusCode)
            }
            count++
        }
    }
    fmt.Printf("Finished with %d errors\n", errorCount)
    fmt.Printf("Total Duration: %d ms. Average Duration: %d ms\n", durationCount, durationCount / requestNum)
}
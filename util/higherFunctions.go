package util

import (
    "time"
)

type executable func()

func Repeat(times int, f executable) {
    for i := 0; i < times; i++ {
        f()
    }
}

func InvokeAndGetTime(f executable) int{
    now := time.Now()
    f()
    return int(time.Since(now)/time.Millisecond)
}
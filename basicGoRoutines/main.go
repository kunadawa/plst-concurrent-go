package main

import (
  "time"
  "runtime"
)

func main() {
  dur10ms, _ := time.ParseDuration("10ms")
  runtime.GOMAXPROCS(2)
  
  go func () {
    for i := 0; i < 100; i++ {
      println("hello")
      time.Sleep(dur10ms)
    }
  }()

  go func() {
    for i := 0; i < 100; i++ {
      println("go")
      time.Sleep(dur10ms)
    }
  }()

  dur, _ := time.ParseDuration("1s")
  time.Sleep(dur)
}

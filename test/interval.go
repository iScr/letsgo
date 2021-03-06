/*=============================================================================
#     FileName: gate.go
#         Desc: game gate server
#       Author: sunminghong
#        Email: allen.fantasy@gmail.com
#     HomePage: http://weibo.com/5d13
#      Version: 0.0.1
#   LastChange: 2013-06-09 10:09:28
#      History:
=============================================================================*/
package helper

import (
    "fmt"
    "time"
    "errors"
    "runtime"
)


type LGInterval struct {
    duration time.Duration
    quit chan bool
    stop chan bool
    callback func(self *LGInterval)

    isRun bool

    timer *time.Ticker

    isExecuting bool
}

func NewLGInterval(duration time.Duration,callbackfn func(self *LGInterval)) *LGInterval {
    return &LGInterval{duration,make(chan bool),make(chan bool),callbackfn,false,nil,false}
}

func (self *LGInterval) Start(newDuration ...time.Duration) error {
    if len(newDuration)>0 {
        self.duration = newDuration[0]
    }
    
    if self.isRun {
        return errors.New("the instance is already running!")
    }

    go func() {
        fmt.Println("is starting...")
        self.isRun = true
        self.timer = time.NewTicker(self.duration)
        for {
            if !self.isRun {
                //<-self.quit
                goto quit
            }

            select {
            //case <-self.quit:
                //fmt.Println("quit")
                //goto quit
            case <-self.timer.C:
                self.isExecuting = true
                self.callback(self)
                self.isExecuting = false
            }
        }

        quit:
        self.isRun = false
        self.stop <- true
    }()

    return nil
}

var istop int =0
func (self *LGInterval) Stop(callback ...func(interval *LGInterval)) {
    go func() {
        if self.isRun {
            self.isRun = false
            //self.quit<- true
            fmt.Println("enter _stop()")

            a := istop
            istop++
            //for self.isExecuting {
                //fmt.Println("self.isExecuting is true")
                //time.Sleep(100*time.Millisecond)
            //}
            fmt.Println("istop a:",a)
            <-self.stop
            self.timer.Stop()
            fmt.Println("istop a:",a)

            if len(callback)>0 {
                callback[0](self)
            }

            fmt.Println("-----||```------")
        }
        fmt.Println("-----------")
    }()
}

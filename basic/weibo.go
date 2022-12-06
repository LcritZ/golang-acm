package basic

import (
    "fmt"
)

/**
多个chan输入

一个输出

c1 := make(chan int, 5) 1-5
c2 := make(chan int, 5) 5-6


c3 := make(chan int, 5)


 */

func input(o chan int, c <-chan int) {
    for {
        select {
        case x := <-c:
            o<-x
        }
    }
}

func Send() {
    c1 := make(chan int, 5)
    c2 := make(chan int, 5)

    c3 := make(chan int, 5)


    go func() {
        for i := 1; i<=5; i++ {
            c1<-i
        }
    }()

    go func() {
        for i := 6; i<=10; i++ {
            c2<-i
        }
    }()

    go func() {
        for {
            select {
            case m := <-c1:
                c3<-m
            case n := <-c2:
                c3<-n
            default:
                break
            }
        }
    }()


    for {
        select {
        case x := <-c3:
            fmt.Println(x)
        default:
            break
        }
    }

}


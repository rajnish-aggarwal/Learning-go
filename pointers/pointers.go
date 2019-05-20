package main

import "fmt"

func main() {
    i, j := 1, 200

    p := &i
    fmt.Println(*p)

    *p = 32
    fmt.Println(i)

    p = &j
    *p /= 20
    fmt.Println(j)
}

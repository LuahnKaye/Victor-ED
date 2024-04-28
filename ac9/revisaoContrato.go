package main

import (
    "fmt"
    "strings"
)

func main() {
    var d string
    var v string

    fmt.Scanf("%s", &d)
    fmt.Scanf("%s", &v)

    for d != "0" {
        n := len(v)
        for i := 0; i < n; i++ {
            if strings.Contains(v, d) {
                v = strings.ReplaceAll(v, d, "")
                n = len(v)
            }
        }

        v = strings.TrimLeft(v, "0")

        if len(v) == 0 {
            v = "0"
        }

        fmt.Println(v)

        fmt.Scanf("%s", &d)
        fmt.Scanf("%s", &v)
    }
}
package main

import "fmt"
import "math/rand"
import "time"

const GROUNDHOG = "           __________________\n          /                  \\\n         |    ___      ___    |\n       __|   /   \\    /   \\   |__\n      /      | o |    | o |       \\\n     |  C    \\___/    \\___/    D  |\n      \\__                       __/\n         |        ====         |\n         |         ][          |\n         |                     |"
const SHADOW = "         \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n          \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n         \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n           \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n             \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\"

func groundhog_day() {
    var today = time.Now()

    if today.Month() != time.February || today.Day() != 2 {
        fmt.Println("It's not Groundhog Day!")
    } else {
        rand.Seed(today.UTC().UnixNano())

        fmt.Println(GROUNDHOG)

        if rand.Intn(2) == 1 {
            fmt.Println(SHADOW)
            fmt.Println("6 More weeks of winter")
        } else {
            fmt.Println("It looks like an early spring")
        }
    }
}

func main() {
    groundhog_day()
}
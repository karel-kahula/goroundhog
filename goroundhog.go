package main

import "fmt"
import "io"
import "math/rand"
import "os"
import "os/user"
import "strings"
import "strconv"
import "time"

const GROUNDHOG = "           __________________\n          /                  \\\n         |    ___      ___    |\n       __|   /   \\    /   \\   |__\n      /      | o |    | o |       \\\n     |  C    \\___/    \\___/    D  |\n      \\__                       __/\n         |        ====         |\n         |         ][          |\n         |                     |"
const SHADOW = "         \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n          \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n         \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n           \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n             \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\"

func get_filepath() string{
    usr, _ := user.Current()
    dir := usr.HomeDir
    var path = strings.Replace("~/.goroundhog", "~", dir, 1)

    fmt.Println(path)
    return path
}

func groundhog_day() {
    var today = time.Now()

    if today.Month() != time.February || today.Day() != 2 {
        fmt.Println("It's not Groundhog Day!")
    } else {
        var value = -1;

        file, err := os.Open(get_filepath())
        if err == nil {
            // file exists
            defer file.Close()

            stat, err := file.Stat()
            if err == nil {
                // file has stuff
                bs := make([]byte, stat.Size())
                _, err = file.Read(bs)
                if err == nil {
                    // lets see if its the right year
                    str := string(bs)

                    if strings.HasPrefix(str, strconv.Itoa(today.Year()) ) {
                        if strings.HasSuffix(str, "1") {
                            value = 1
                        } else {
                            value = 0
                        }
                    }
                }
            }
        }

        if value == -1 {
            rand.Seed(today.UTC().UnixNano())
            value = rand.Intn(2)
        }

        fmt.Println(GROUNDHOG)

        if value == 1 {
            fmt.Println(SHADOW)
            fmt.Println("6 More weeks of winter")
        } else {
            fmt.Println("It looks like an early spring")
        }

        out_file, err := os.Create(get_filepath())
        if err == nil {
            io.WriteString(out_file, strconv.Itoa(today.Year()) + " " + strconv.Itoa(value))
            out_file.Close()
        }
    }
}

func main() {
    groundhog_day()
}
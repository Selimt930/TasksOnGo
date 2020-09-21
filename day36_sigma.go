//https://www.hackerrank.com/challenges/game-of-thrones/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Complete the gameOfThrones function below.
func gameOfThrones(s string) string {
var res string
    letters_count := make(map[byte]int)
    for ix := range s {
        letters_count[s[ix]]++
    }
    var odd_count int = 0
    for l := range letters_count {
        if letters_count[l]%2 != 0 {
            odd_count++
        }
    }

    if odd_count > 1 {
        res ="NO"
    } else {
        if len(s)%2 != 0 {
            if odd_count == 1 {
                res ="YES"
            } else {
                res ="NO"
            }
        } else {
            if odd_count == 0 {
                res ="YES"
            } else {
                res ="NO"
            }
        }
    }
return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := gameOfThrones(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

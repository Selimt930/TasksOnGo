//https://www.hackerrank.com/challenges/strange-code/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the strangeCounter function below.
func strangeCounter(t int64) int64 {
var res int64 = 3
    for t > res {
        t = t - res
        res *= 2
    }
return res-t+1
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    t, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    result := strangeCounter(t)

    fmt.Fprintf(writer, "%d\n", result)

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

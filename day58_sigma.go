//https://www.hackerrank.com/challenges/priyanka-and-toys/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

// Complete the toys function below.
func toys(w []int32) int32 {

n := len(w)
ws := make([]int, len(w))

    for i, c := range w {
        ws[i] = int(c)
    }

    sort.Ints(ws)

    var count, inc int
    for i := 0; i < n; i += inc {
        inc = 0
        for k := i; k < n && ws[k] <= ws[i]+4; k++ {
            inc++
        }
        count++
    }
return int32(count)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    wTemp := strings.Split(readLine(reader), " ")

    var w []int32

    for i := 0; i < int(n); i++ {
        wItemTemp, err := strconv.ParseInt(wTemp[i], 10, 64)
        checkError(err)
        wItem := int32(wItemTemp)
        w = append(w, wItem)
    }

    result := toys(w)

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

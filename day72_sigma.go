//https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func abs(n int32) int32 {
    if n >= 0 {
        return n
    } else {
        return -n
    }
}
// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
var num string
var rev string
var revnum int32
var res int32 = 0
    for p := i; p <= j; p++ {
        num = strconv.FormatInt(int64(p), 10)
        rev = ""
        for q := len(num) - 1; q >= 0; q-- {
            rev += string(num[q])
        }
        revnum64, err := strconv.ParseInt(rev, 10, 64)
        if err != nil {
            panic(err)
        }
        revnum = int32(revnum64)
        if abs(p - revnum) % k == 0 {
            res++
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

    ijk := strings.Split(readLine(reader), " ")

    iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
    checkError(err)
    i := int32(iTemp)

    jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
    checkError(err)
    j := int32(jTemp)

    kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := beautifulDays(i, j, k)

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

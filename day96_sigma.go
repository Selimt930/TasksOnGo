//https://www.hackerrank.com/challenges/weighted-uniform-string/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the weightedUniformStrings function below.
func weightedUniformStrings(s string, queries []int32) []string {
var res []string
var temp  int32 = 0
    shift := ' '
    m := make(map[int32]int)
    for _, ch := range s {
        if ch != shift {
            temp = int32(ch - 'a' + 1)
        } else {
            temp += int32(ch - 'a' + 1)
        }
        shift = ch
        m[temp]++
    }
    for i := 0; i < len(queries); i++ {
        if _, ok := m[queries[i]]; ok {
            res = append(res, "Yes")
        } else {
            res = append(res, "No")
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

    queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    var queries []int32

    for i := 0; i < int(queriesCount); i++ {
        queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        queriesItem := int32(queriesItemTemp)
        queries = append(queries, queriesItem)
    }

    result := weightedUniformStrings(s, queries)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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

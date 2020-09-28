//https://www.hackerrank.com/challenges/gem-stones/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the gemstones function below.
func gemstones(arr []string) int32 {
var Ltr [26]byte
const ACODE = 97
var N = len(arr)
var i int
for i = 0; int(i) < N; i++ {
        line := arr[i]
        var visited [26]bool

        for k := range line {
            if !visited[line[k]-ACODE] {
                Ltr[line[k]-ACODE] += 1
                visited[line[k]-ACODE] = true
            }
        }
}
        var res int32 = 0
        for i = 0; i < 26; i++ {
            if int(Ltr[i]) == N {
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

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var arr []string

    for i := 0; i < int(n); i++ {
        arrItem := readLine(reader)
        arr = append(arr, arrItem)
    }

    result := gemstones(arr)

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

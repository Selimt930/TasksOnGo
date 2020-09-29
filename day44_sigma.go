//https://www.hackerrank.com/challenges/funny-string/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the funnyString function below.
func funnyString(s string) string {
var res string
leng := len(s)
    var is bool = true
    for i, r := 1, leng-2; i < leng; i++ {
if math.Abs(float64(s[i])-float64(s[i-1])) != math.Abs(float64(s[r])-float64(s[r+1])) {
                is = false
                break
            }
            r--
        }
        if is {
            res = "Funny"
        } else {
            res = "Not Funny"
        }
return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s := readLine(reader)

        result := funnyString(s)

        fmt.Fprintf(writer, "%s\n", result)
    }

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

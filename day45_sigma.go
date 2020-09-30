//https://www.hackerrank.com/challenges/alternating-characters/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the alternatingCharacters function below.
func alternatingCharacters(s string) int32 {
char := s[0]
    var current, res, size int32 = 0, 0, int32(len(s))
    var flag bool = true
    
        for i := 1; int32(i) < size; i++ {
            if char != s[i] {
                res += (int32(i) - current - 1)

                current = int32(i)
                char = s[i]

                flag = false
            } else {
                flag = true
            }
        }
        if flag {
            res += (size - current - 1)
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

        result := alternatingCharacters(s)

        fmt.Fprintf(writer, "%d\n", result)
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

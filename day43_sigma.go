//https://www.hackerrank.com/challenges/pangrams/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Complete the pangrams function below.
func pangrams(s string) string {
letters_count := make(map[byte]bool)
var line string = strings.ToLower(s)
var res string 
for i := range line {
        if line[i] >= 97 && line[i] <= 122 { 
            letters_count[line[i]] = true
        }
    }
    if len(letters_count) != 26 {
        res = "not pangram" 
    } else {
        res= "pangram"
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

    result := pangrams(s)

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

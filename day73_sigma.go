//https://www.hackerrank.com/challenges/append-and-delete/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
var res string
var min string
var length int32
var temp int32 
var count int32
    if len(s) <= len(t) {
        min = s
    } else {
        min = t
    }
    for i := 0; i < len(min); i++ {
        if s[i] != t[i] {
            break
        }
        count++
    }
    var len1,len2 int32 = int32(len(s)), int32(len(t))
    temp = (len1 - count) + (len2 - count)
    length = len1 + len2
    res = "No"
    if temp>k{
        res = "No"
    }else if  temp ==k{
        res="Yes"
    }else if length % 2 == k%2{
        res = "Yes"
    }else if length<k{
        res = "Yes"
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

    t := readLine(reader)

    kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := appendAndDelete(s, t, k)

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

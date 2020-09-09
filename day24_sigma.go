//https://www.hackerrank.com/challenges/mini-max-sum/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
var elems int64
var maxel int32 = arr[0]
var minel int32 = arr[0]

for i:=0; i<5;i++{
    if arr[i]>maxel{
        maxel = arr[i]
    } else if arr[i]<minel{
        minel = arr[i]
    }
    elems += int64(arr[i])
}
fmt.Println(elems-int64(maxel),elems-int64(minel))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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

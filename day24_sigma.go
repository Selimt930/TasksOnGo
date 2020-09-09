//https://www.hackerrank.com/challenges/mini-max-sum/problem
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

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
var arr2 []int
sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
for i :=range arr{
    arr2 = append(arr2, int(arr[i]))
}
//sort.Ints(arr2)
min := arr2[0]+arr2[1]+arr2[2]+arr2[3]
max := arr2[1]+arr2[2]+arr2[3]+arr2[4]
fmt.Println(min,max)
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

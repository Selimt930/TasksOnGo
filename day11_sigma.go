//https://www.hackerrank.com/challenges/greedy-florist/problem
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

// Complete the getMinimumCost function below.
func getMinimumCost(k int32, c []int32) int32 {
s := []int {}
k1 := int(k)
for i := range c{
    s = append(s, int(c[i]))
}
sort.Sort(sort.Reverse(sort.IntSlice(s)))
var cost int = 0
var prev_purch int = 0

for i := range s{
    cost += (prev_purch +1) * s[i]
    if (i+1)%k1 ==0{
        prev_purch +=1
    }
}
return int32(cost)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    cTemp := strings.Split(readLine(reader), " ")

    var c []int32

    for i := 0; i < int(n); i++ {
        cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
        checkError(err)
        cItem := int32(cItemTemp)
        c = append(c, cItem)
    }

    minimumCost := getMinimumCost(k, c)

    fmt.Fprintf(writer, "%d\n", minimumCost)

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

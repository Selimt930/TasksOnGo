//https://www.hackerrank.com/challenges/ctci-coin-change/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the ways function below.
func ways(n int32, coins []int32) int64 {
var m = len(coins)
var res int64
NN := int64(n)
cns := make([]int64, m)
nways := make([]int64, n+1)

for i := 0; i < m; i++ {
        cns[i] = int64(coins[i])
    }
    nways[0] = 1
    for j := 0; j < m; j++ {
        c := cns[j]
        for k := c; int64(k) <= NN; k++ {
            nways[k] += nways[k-c]
        }
    }
    res = nways[n]
    return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    coinsTemp := strings.Split(readLine(reader), " ")

    var coins []int32

    for i := 0; i < int(m); i++ {
        coinsItemTemp, err := strconv.ParseInt(coinsTemp[i], 10, 64)
        checkError(err)
        coinsItem := int32(coinsItemTemp)
        coins = append(coins, coinsItem)
    }

    res := ways(n, coins)

    fmt.Fprintf(writer, "%d\n", res)

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

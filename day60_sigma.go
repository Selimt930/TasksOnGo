//https://www.hackerrank.com/challenges/the-grid-search/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the gridSearch function below.
func gridSearch(G []string, P []string) string {
R := len(G)
m := len(P)
var res string
    var found bool

        for ix := 0; ix < R-m; ix++ {

            if found_index := strings.Index(G[ix], P[0]); found_index != -1 {
                b := 1
                for a := ix + 1; b < m; b++ {
                    if currix := strings.Index(G[a], P[b]); currix != found_index {
                        if strings.Index(G[a][found_index:], P[b]) != 0 {
                            break
                        }
                    }
                    a++
                }

                if b == m {
                    found = true
                    break
                }
            }
        } // for

        if found {
            res = "YES"
        } else {
            res = "NO"
        }
    return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        RC := strings.Split(readLine(reader), " ")

        RTemp, err := strconv.ParseInt(RC[0], 10, 64)
        checkError(err)
        R := int32(RTemp)

        //CTemp, err := strconv.ParseInt(RC[1], 10, 64)
        checkError(err)
        //C := int32(CTemp)

        var G []string

        for i := 0; i < int(R); i++ {
            GItem := readLine(reader)
            G = append(G, GItem)
        }

        rc := strings.Split(readLine(reader), " ")

        rTemp, err := strconv.ParseInt(rc[0], 10, 64)
        checkError(err)
        r := int32(rTemp)

        //cTemp, err := strconv.ParseInt(rc[1], 10, 64)
        checkError(err)
        //c := int32(cTemp)

        var P []string

        for i := 0; i < int(r); i++ {
            PItem := readLine(reader)
            P = append(P, PItem)
        }

        result := gridSearch(G, P)

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

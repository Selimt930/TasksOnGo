//https://www.hackerrank.com/challenges/find-a-word/problem?h_r=internal-search
package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

var reg = regexp.MustCompile(`[[:word:]]+`)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    var buffer bytes.Buffer
    for ; n > 0; n-- {
        scanner.Scan()
        buffer.WriteString(scanner.Text())
        buffer.WriteString("\n")
    }

    scanner.Scan()
    t, _ := strconv.Atoi(scanner.Text())

    var text string = buffer.String()

    words_map := make(map[string]int)
    all_words := reg.FindAllString(text, -1)

    for ix := range all_words {
        var current string = strings.ToLower(all_words[ix])
        if _, exists := words_map[current]; !exists {
            words_map[current] = 0
        }
        words_map[current]++
    }

    for ; t > 0; t-- {
        scanner.Scan()
        word := strings.ToLower(scanner.Text())
        var res int = 0
        if cr, exists := words_map[word]; exists {
            res = cr
        }
        fmt.Println(res)
    }
}
//https://www.hackerrank.com/challenges/find-substring/problem?h_r=internal-search
package main

import (
    "bufio"
    "container/list"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

var reg = regexp.MustCompile(`[[:word:]]+`)
var n, T int

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())

    lines := list.New()

    for ; n > 0 && scanner.Scan(); n-- {
        var words []string = reg.FindAllString(scanner.Text(), -1)
        lines.PushBack(words)
    }

    scanner.Scan()
    T, _ := strconv.Atoi(scanner.Text())

    for ; T > 0 && scanner.Scan(); T-- {
        var current_word string = scanner.Text()
        var res, curr_word_len int = 0, len(current_word)

        for elem := lines.Front(); elem != nil; elem = elem.Next() {
            sentence := elem.Value.([]string)

            for i := range sentence {
                var word_sentence = sentence[i]
                if len(word_sentence) > curr_word_len {
                    ftix := strings.Index(word_sentence, current_word)

                    if ftix > 0 && ftix < len(word_sentence)-curr_word_len {
                        res++
                    }
                }
            }
        }
        fmt.Println(res)
    } 
}
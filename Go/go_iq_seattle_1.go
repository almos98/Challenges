package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {
    inputs := []string{     // expected ouput
        "a",                // a
        "a{b}",             // ab
        "a{bc}",            // ab, ac
        "a{bc}d{ef}",       // abde, abdf, acde, acdf
    }

    for _, input := range inputs {
        fmt.Println(genStrings(input))
    }
}

func normalize (is []int, groups []string) []int {
    for i := len(is)-1; i > 0; i-- {
        if is[i] > len(groups[i])-1 {
            is[i] = 0
            if i > 0 {
                is[i-1]++
            }
        }
    }

    return is
}

func count (groups []string) int {
    if len(groups) == 0 {
        return 0
    }

    res := 1
    for _, g := range groups {
        res *= len(g)
    }
    return res
}
func genStrings(str string) []string {
    reg, err := regexp.Compile("{[^{}]+}")
    if err != nil {
        fmt.Printf("Error compiling RegExp: %s\n", err)
    }
    groups := reg.FindAllString(str, -1)
    if len(groups) == 0 {
        return []string{str}
    }
    // Remove brackets
    for i, group := range groups {
        groups[i] = group[1:len(group)-1]
    }

    totalRuns := count(groups)
    base := reg.ReplaceAllString(str, "{}")
    res := []string{}
    indices := make([]int, len(groups))
    for i := 0; i < totalRuns; i++ {
        s := base
        for i, j := range indices {
            s = strings.Replace(s, "{}", string(groups[i][j]), 1)
        }
        res = append(res, s)

        indices[len(indices)-1]++
        indices = normalize(indices, groups)
    }

    return res
}

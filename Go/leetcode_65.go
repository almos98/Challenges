/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 *
 * https://leetcode.com/problems/valid-number/description/
 *
 * algorithms
 * Hard (14.45%)
 * Total Accepted:    139.7K
 * Total Submissions: 966.6K
 * Testcase Example:  '"0"'
 *
 * Validate if a given string can be interpreted as a decimal number.
 * 
 * Some examples:
 * "0" => true
 * " 0.1 " => true
 * "abc" => false
 * "1 a" => false
 * "2e10" => true
 * " -90e3   " => true
 * " 1e" => false
 * "e3" => false
 * " 6e-1" => true
 * " 99e2.5 " => false
 * "53.5e93" => true
 * " --6 " => false
 * "-+3" => false
 * "95a54e53" => false
 * 
 * Note: It is intended for the problem statement to be ambiguous. You should
 * gather all requirements up front before implementing one. However, here is a
 * list of characters that can be in a valid decimal number:
 * 
 * 
 * Numbers 0-9
 * Exponent - "e"
 * Positive/negative sign - "+"/"-"
 * Decimal point - "."
 * 
 * 
 * Of course, the context of these characters also matters in the input.
 * 
 * Update (2015-02-10):
 * The signature of the C++ function had been updated. If you still see your
 * function signature accepts a const char * argument, please click the reload
 * button to reset your code definition.
 * 
 */

 import (
    "strconv"
    "regexp"
    "strings"
    "math"
)

func ParseFloat(s string) (float64, error) {
    val, err := strconv.ParseFloat(s, 64)
    if err == nil {
        return val, nil
    }
    
    var base float64
    var exp int64
    
    pos := strings.IndexAny(s, "e")
    if pos < 0 {
        return strconv.ParseFloat(s, 64)
    }
    
    baseStr := s[0:pos]
    base, err = strconv.ParseFloat(baseStr, 64)
    if err != nil {
        return 0, err
    }
    
    expStr := s[(pos + 1):]
    exp, err = strconv.ParseInt(expStr, 10, 64)
    if err != nil {
        return 0, err
    }
    
    return base * math.Pow10(int(exp)), nil
}

func isNumber(s string) bool {
    s = strings.TrimSpace(s)
    reg, _ := regexp.Compile(".?[0-9e+-.]+.?")
    res := reg.FindAllString(s, -1)
    if len(res) > 1 || len(res) == 0 {
        return false
    }
    s = res[0]
    
    if _, err := strconv.ParseInt(s, 10, 64); err == nil {
        return true
    }
    
    if _, err := ParseFloat(s); err == nil {
        return true
    }
    
    return false
}

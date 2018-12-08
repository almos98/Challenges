{- Largest palindrome product
 -
 - A palindromic number reads the same both ways. The largest palindrome made
 - from the product of two 2-digit numbers is 9009 = 91 x 99.
 -
 - Find the largest palindrome made from pdocut of two 3-digit numbers.
 -
 -
 - To check if a number is a palindrome, it is converted to a string and
 - compared to the reverse of this string.
 -
 - Generate a list of all palindromic products of 2 3-digit numbers and return
 - the maximum value.
-}

largestPalindromeProduct =
    maximum [x | y <- [100..999], z <- [y..999], let x=y*z, let s = show x, s==reverse s]

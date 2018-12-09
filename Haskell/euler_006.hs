{- Sum square difference
 - The sum of the squares of the first ten natural numbers is,
 -
 -                  1^2 + 2^2 + ... + 10^2 = 385
 -
 - The square of the sum of the first ten natural numbers is,
 -                  
 -                  (1 + 2 + ... + 10) ^ 2 = 55^2 = 3025
 -
 - Hence the difference between the sum of the squares of the first ten natural
 - numbers and the square of the sum is 3025 - 385 = 2640.
 - 
 - Find the difference between the sum of the squares of the first one hundred
 - natural numbers and the square of the sum.
 -
 -
 - Generate a list of squares and then add them.
 - Add a list of the first N natural numbers and then square it.
 - Then subtract.
-}

sumSquareDiff :: Integer -> Integer
sumSquareDiff n = (sumSquare n) - (sum $ squareList n)
 where
    squareList n = [x*x | x <- [1..n]]
    sumSquare n = sum [1..n] * sum [1..n]
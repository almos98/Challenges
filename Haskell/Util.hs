module Util
( primes
, primeFactors
, count
) where
primes :: [Integer]
primes = 2 : 3 : 5 : filter (null . tail . primeFactors) [7,9..]

primeFactors n = factor n primes
 where
    factor n (x:xs)
        | x * x > n      = [n]
        | n `mod` x == 0 = x : factor (n `div` x) (x:xs)
        | otherwise      = factor n xs

count :: Eq a => a -> [a] -> Int
count x = length . filter (x==)
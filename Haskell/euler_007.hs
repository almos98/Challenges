{- 10001st Prime
 -
 - By listing the first six prime numbers: 2, 3, 5, 7, 11 and 13, we can see
 - that the 6th prime is 13.
 -
 - What is the 10 001st prime number?
 -
 -
 - Generate a list of primes by filtering the list of odd numbers by those that
 - have no factors other than itself. Then get the 10001st element.
-}

primes = 2 : 3 : 5 : filter (null . tail . primeFactors) [7,9..]

primeFactors n = factor n primes
 where
    factor n (x:xs)
      | x*x > n         = [n]
      | n `mod` x == 0  = x : factor (n `div` x) (x:xs)
      | otherwise       = factor n xs

nthPrime n = primes !! (n-1)

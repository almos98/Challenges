{- Largest Prime Factor
 -
 - The prime factors of 13195 are 5, 7, 13 and 29.
 -
 - What is the largest prime factor of the number 600851475143 ?
 -
 -
 - Generate a list of primes by filtering the list of odd numbers by those that
 - have no factors other than itself.
 - To factorise, if X squared is greater than N, then N is the only factor.
 - Otherwise if n mod x is 0, x is a factor.
 - Otherwise x is not a factor, continue through the list.
 - The largest prime factor is the last entry in the list.
-}

primes = 2 : 3 : 5 : filter (null . tail . primeFactors) [7,9..]

primeFactors n = factor n primes
 where
    factor n (x:xs)
      | x*x > n         = [n]
      | n `mod` x == 0  = x : factor (n `div` x) (x:xs)
      | otherwise       = factor n xs

largestPrimeFactor n = last $ primeFactors n

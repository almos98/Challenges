{- Summation of primes
 - 
 - The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17
 -
 - Find the sum of all the primes below two million.
 - 
 -
 - Create a subset of the primes list in Util.hs that contains primes below N
 - and then sum them up.
 - This requires a more efficient prime generation algorithm.
 - https://www.cs.tufts.edu/~nr/cs257/archive/melissa-oneill/Sieve-JFP.pdf
-}

import Util

sumOfPrimes n = sum $ takeWhile (\ a -> a < n) primes
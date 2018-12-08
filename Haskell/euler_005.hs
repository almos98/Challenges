{- Smallest multiple
 - 2520 is the smallest number that can be divided by each of the numbers from
 - 1 to 10 without any remainder.
 -
 - What is the smallest positive number that is evenly divisible by all of the
 - numbers from 1 to 20?
 -
 - 
 - Breaking down all the numbers from 1 to 10 into its prime factors we get
 -      1:  1           = 1 ^ 1
 -      2:  2           = 2 ^ 1
 -      3:  3           = 3 ^ 1
 -      4:  2, 2        = 2 ^ 2
 -      5:  5           = 5 ^ 1
 -      6:  2, 3        = 2 ^ 1 * 3 ^ 1
 -      7:  7           = 7 ^ 1
 -      8:  2, 2, 2     = 2 ^ 3
 -      9:  3, 3        = 3 ^ 2
 -      10: 2, 5        = 2 ^ 1 * 5 ^ 1
 -
 - The goal is to eliminate all the lowest powers of each factor:
 -      5:  5           = 5 ^ 1
 -      7:  7           = 7 ^ 1
 -      8:  2, 2, 2     = 2 ^ 3
 -      9:  3, 3        = 3 ^ 2
 -
 - The product of the remaining numbers is the desired result:
 -      5 * 7 * 8 * 9 = 2520
-}
import Util
import qualified Data.Map as Map
import qualified Data.List as List

smallestMultiple n = evaluate $ highestPowers n
 where
    factorise [] = []
    factorise (x:xs) = primeFactors x : factorise xs

    highestPowers n = Map.fromList $ List.sortBy (\ (_,a) (_,b) -> compare a b) $ filter (\ x -> snd x > 0) [x | x' <- factorise [2..n], let y = maximum x', let z = minimum x', y' <- [z..y], let x = (y', (count y' x'))]

    evaluate m = fst $ Map.mapAccumWithKey (\ a k b -> (a * k ^ b, b)) 1 m
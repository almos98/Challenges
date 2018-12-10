{- Special Pythagorean triplet
 -
 - A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
 -              
 -                  a^2 + b^2 = c^2
 -
 - For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2
 - 
 - There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 - Find the product abc.
 -
 -
 - Using a list compreehension and Euclid's formula to generate the triplet then
 - get the product.
-}

pythagoreanTriplet x = head $
    [ a*b*c 
    | m <- [1..x]
    , n <- [1..x]
    , m > n
    , let a = m*m - n*n
    , let b = 2*m*n
    , let c = m*m + n*n
    , a + b + c == 1000
    ]
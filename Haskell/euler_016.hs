{- Power digit sum
 - 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8  = 26.
 -
 - What is the sum of the digits of the number 2^1000?
 -
 -
 - Separate the number 2^N into digits byt converting it into a string and then
 - converting each character to a number. Then return the sum.
-}

import Data.Char

sumOfDigits n = sum $ map (\ a -> digitToInt a) $ show $ 2 ^ n
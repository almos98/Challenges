--[[
--      The Problem: Even Fibonacci numbers
--  Each new term in the Fibonacci sequence is generated by adding the previous
--  two terms. By starting with 1 and 2, the first 10 terms will be:
--              
--              1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
--  
--  By considering the terms in the Fibonacci sequence whose values do not
--  exceed four million, find the sum of the even-valued terms.
--
--      The Solution
--  Recursively generate all Fibonacci numbers, if they pass the check function,
--  add it to the sum. The check function in this case checks if a number is
--  even.
]]


local sum = 2
local limit = 4000000

function fibonacci (num1, num2, check)
   local check = check
   local nextNum = num1 + num2
   if nextNum < limit then
      if check(nextNum) then
	     sum = sum + nextNum
      end
      fibonacci(num2, nextNum, check)
   end
end

fibonacci(1,2,function(n) return n % 2 == 0 end)
print(sum)

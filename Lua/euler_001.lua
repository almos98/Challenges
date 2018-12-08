--[[
--      The Problem: Multiples of 3 and 5
--  If we list all the natural numbers below 10 that are multiples of 3 or 5,
--  we get 3, 5, 6 and 9. The Sum of these multiples is 23.
--  
--  Find the sum of all the multiples of 3 or 5 below 1000.
--
--      The Solution
--  Generate all numbers until N-1, where N = 1000. If a number is a multiple
--  of 3 or 5 then add it to an accumulator.
]]

function sumOfMultiples (multiples, limit)
   local sum = 0
   for i = 1, limit do
      local multiple = false
      for _, n in next, multiples do
	 multiple = multiple or (i % n == 0)
      end
      sum = multiple and sum + i or sum
   end
   return sum
end

print(sumOfMultiples({3,5}, 999))

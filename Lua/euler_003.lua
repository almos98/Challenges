--[[
--      The Problem: Largest prime factor
--  The prime factors of 13195 are 5, 7, 13 and 29.
--  
--  What is the largest prime factor of the number 600851475143 ?
--      
--      The Solution
--  Generate all primes until Sqrt N and then get the prime factors of N and
--  return the largest one.
--
--  NOTE: This could do with a lot of improvements.
]]

local n = 600851475143
local LPF = 0

local sieve = {}
function buildSieve (n)
   print("Building sieve...")
   print("Generating all possible numbers to n")
   local limit = math.floor(math.sqrt(n))
   for i = 2, limit do
      local prime = true
      for j, v in ipairs(sieve) do
	 prime = prime and not (i % j == 0)
      end
      if prime then
	 sieve[i] = true
      end
   end
   print("Sieve complete")
end

buildSieve(n)
local currPrime = 2
while (n > 1) do
   if (sieve[currPrime]) then
      n2 = n / currPrime
      if (math.floor(n2) == n2) then
	 LPF = currPrime
	 n = n2
      else
	 currPrime = currPrime + 1
      end
   else
      currPrime = currPrime + 1
   end
end

print(LPF)

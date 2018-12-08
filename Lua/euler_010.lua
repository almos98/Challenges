-- Using Sieve of Sundaram
local function primeCheck (n)
   local limit = math.floor(math.sqrt(n))
   local isPrime = true
   local i = 2
   while (i <= limit) do
      isPrime = isPrime and not (n % i == 0)
      i = i + 1
   end
   return isPrime
end

function sumOfPrimes (limit)
   local sum = 0
   for i = 2, limit do
      if primeCheck(i) then
	 sum = sum + i
      end
   end
   return sum
end

print(sumOfPrimes(2000000))

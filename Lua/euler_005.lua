--Requires Lua 5.3

function smallestMultiple (n)
   -- Finding all primes up to n
   print("Generating all primes until  " .. n .. "...")
   local primes = {}
   for i = 2, n do
      local prime = true
      for j in next, primes do
	 prime = prime and not (i % j == 0)
      end
      primes[i] = prime or nil
   end

   print("Creating factor table...")
   -- Populating the factor table
   local factorTable = {}
   for i = 2, n do
      factorTable[i] = {}
      for prime in next, primes do
	 factorTable[i][prime] = 1
      end
   end

   local function nextPrime(curr)
      for i = curr+1, n do
	 if (primes[i]) then return i end
      end
   end

   print("Reducing to prime factors...")
   -- Reducing to prime factors
   for x = n, 2, -1 do
      local workingNumber = x
      local prime = nextPrime(1)
      if (primes[workingNumber]) then
	 factorTable[x][workingNumber] = factorTable[x][workingNumber + 1]
      else
	 while (not primes[workingNumber]) do
	    local temp = workingNumber/prime
	    if math.floor(temp) == temp then
	       workingNumber = temp
	       factorTable[x][prime] = factorTable[x][prime] + 1
	    else
	       prime = nextPrime(prime)
	    end
	 end
      end
   end

   print("Finding highest exponents...")
   -- Finding highest powers
   local highestPowers = {}
   for prime in next, primes do
      highestPowers[prime] = 0
   end

   for i = 2, n do
      for prime, exponent in next, factorTable[i] do
	 highestPowers[prime] = highestPowers[prime] > exponent and highestPowers[prime] or exponent
      end
   end

   print("Calculating final result...")
   local result = 1
   for prime, exponent in next, highestPowers do
      result = result * prime ^ exponent
   end

   return result
end

print("Smallest multiple is " .. smallestMultiple(20))

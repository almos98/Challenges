-- Requires lua 5.3

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

print(primeCheck(71))
function findPrime (primeCount)
   local limit = primeCount
   local increment = math.floor(limit * 0.1)
   print("Approximating limit...")
   while (limit > 0) do
      local n = limit/math.log(limit)
      if (n >= primeCount) then break end
      limit = limit + increment
   end

   if (limit < 0) then
      print("Overflow occurred")
   else
      print("Limit approximated to " .. limit)
   end

   print ("Generating prime numbers...")
   local count = 0
   local number = 1
   while (count < primeCount and number < limit) do
      number = number + 1
      count = primeCheck(number) and count + 1 or count
   end

   return number
end

print(findPrime(10001))

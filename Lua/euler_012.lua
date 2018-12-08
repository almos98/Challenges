local function factorise (n)
   local factors = {}
   factors.count = 0
   
   for i = 1, math.sqrt(n) do
      if (math.floor(n/i) == n/i) then
	 factors[i] = true
	 factors[n/i] = true
	 factors.count = factors.count + 2
      end
   end
   
   return factors
end

local function highlyDivisibleTriangle(divisors)
   for i = 1, math.huge do
      local triangle = (i + 1) * i / 2
      if factorise(triangle).count > divisors then
	 return triangle
      end
   end
   return "Solution out of range"
end

print(highlyDivisibleTriangle(500))

-- this approach uses Euclid's formula
function SPT (sum)
   local function getDivisors (x)
      local divisors = {}
      for i = 1, x/2 - 1 do
	 if (math.floor(x/i) == x/i) then
	    divisors[x/i] = true
	    divisors[i] = true
	 end
      end

      return divisors
   end

   local sumDiv = sum/2
   local k_values = getDivisors(sumDiv)
   for k in next, k_values do
      m_values = getDivisors(sumDiv/k)
      for m in next, m_values do
	 local n = sumDiv/(k*m) - m
	 local a = k * (m^2 - n^2)
	 local b = k * (2*m*n)
	 local c = k * (m^2+n^2)
	 if (a + b + c == sum) and (a > 0 and b > 0 and c > 0) then
	    return a, b, c
	 end
      end
   end
end

local a,b,c = SPT(1000)
print(a*b*c)

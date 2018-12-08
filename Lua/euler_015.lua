local function fact(n, lowerLimit)
   lowerLimit = lowerLimit or 1
   return n == lowerLimit and lowerLimit or n * fact(n-1)
end

local function binomialCoefficient(n, k)
   return (fact(n))/(fact(k)*fact(n-k))
end

function latticePaths(N)
   return binomialCoefficient(2*N, N)
end

print(latticePaths(20))

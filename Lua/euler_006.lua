-- Written for lua 5.3, not sure if it will work on previous versions.
local n = 100 -- First "n" natural numbers.

local sumOfSquares = 0
local squareOfSums = 0

for i = 1, n do
   sumOfSquares = sumOfSquares + i*i
   squareOfSums = squareOfSums + i
end

squareOfSums = squareOfSums * squareOfSums
local difference = squareOfSums - sumOfSquares
print("The difference is " .. difference)

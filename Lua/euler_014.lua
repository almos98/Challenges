function collatzSequence(startingNumber)
   function even (n)
      return n/2
   end
   
   function odd (n)
      return 3*n + 1
   end

   local length = 1
   local n = startingNumber
   while (n ~= 1) do
      if (n % 2 == 0) then
	 n = even(n)
      else
	 n = odd(n)
      end
      
      length = length + 1
   end

   return length
end

local longest = {0, 0}
for i = 999999, 2, -1 do
   local length = collatzSequence(i)
   if (length > longest[2]) then
      longest[1] = i
      longest[2] = length
   end
end

print(longest[1], longest[2])

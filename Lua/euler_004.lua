function largestPalindromeProduct (numberOfDigits)

   -- Calculating largest possible number.
   local maxNumber = tonumber(string.rep("9", numberOfDigits))
   local minNumber = tonumber("1" .. string.rep("0", numberOfDigits - 1))
   local function separate (n)
      local digits = {}
      
      while (n ~= 0) do
	 table.insert(digits, n % 10)
	 n = math.floor(n/10)
      end
      
      return digits
   end
   
   local function isPalindrome (n)
      local digits = separate(n)
      local palindrome = true
      for i = 1, #digits/2 do
	 palindrome = palindrome and digits[i] == digits[#digits - (i - 1)]
      end
      return palindrome
   end

   for i = maxNumber, 0, -1 do
      if isPalindrome(i) then
	 
      end
   end

   local largest = 0
   for i = maxNumber, minNumber, -1 do
      for j = maxNumber, minNumber, -1 do
	 local product = i * j
	 if isPalindrome(product) then
	    largest = largest > product and largest or product
	 end
      end
   end

   return largest
end

print(largestPalindromeProduct(3))

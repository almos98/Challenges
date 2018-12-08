local bigNumber = {}
bigNumber.new = function (stringValue)
   local number = {}
   stringValue = stringValue:reverse()
   for i = 1, #stringValue do
      local c = stringValue:sub(i,i)
      if tonumber(c) then
	 number[i] = c
      end
   end
   
   local mt = {
      __add = function (lhs, rhs)
	 local finalNumberString = ""
	 local carry = 0
	 local digits = #lhs > #rhs and #lhs or #rhs
	 for i = 1, digits do
	    local digit1 = lhs[i] and tonumber(lhs[i]) or 0
	    local digit2 = rhs[i] and tonumber(rhs[i]) or 0
	    local sum = digit1 + digit2 + carry
	    carry = 0
	    while (sum > 9) do
	       carry = carry + 1
	       sum = sum - 10
	    end
	    finalNumberString = finalNumberString .. sum
	 end

	 finalNumberString = carry > 0 and finalNumberString .. carry or finalNumberString
	 return bigNumber.new(finalNumberString:reverse())
      end;

      __tostring = function ()
	 local str = ""
	 for i, v in next, number do
	    if(tonumber(i) == i) then
	       str = str .. v
	    end
	 end
	 return str:reverse()
      end;
   }

   setmetatable(number, mt)
   return number
end

return bigNumber

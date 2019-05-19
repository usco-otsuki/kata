def is_month (n)
  1 <= n and n <= 12
end

def is_year(n)
  0 <= n and n <= 99
end

s = gets.chomp
f, l = s[0,2].to_i, s[2,2].to_i
 
if is_month(f) and is_month(l)
  puts "AMBIGUOUS"
elsif is_year(f) and is_month(l)
  puts "YYMM"
elsif is_month(f) and is_year(l)
  puts "MMYY"
else
  puts "NA"
end

left_list = []
right_list = []

File.foreach(ARGV[0]) do |line|
  numbers = line.scan(/\d+/)
  left_list.push(numbers[0].to_i)
  right_list.push(numbers[1].to_i)
end

left_list = left_list.sort
left_hash = left_list.to_h { |n| [n, 0] }
right_list = right_list.sort
answer1 = 0
answer2 = 0

left_list.zip(right_list).each do |left, right|
  answer1 += (left - right).abs
  left_hash[right] += 1 if left_hash.has_key? right
end

left_list.each do |n|
  answer2 += n * left_hash[n]
end

puts "Answer 1: #{answer1}"
puts "Answer 2: #{answer2}"

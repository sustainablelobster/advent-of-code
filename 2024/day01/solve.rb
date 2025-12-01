pairs = ARGF.each_line.map do |line|
  line.scan(/\d+/).map!(&:to_i)
end

left_list, right_list = pairs.transpose
left_list.sort!
right_list.sort!

answer1 = left_list.zip(right_list).sum do |left, right|
  (left - right).abs
end

puts "Answer 1: #{answer1}"

right_tally = right_list.tally
answer2 = left_list.sum do |n|
  n * right_tally.fetch(n, 0)
end

puts "Answer 2: #{answer2}"

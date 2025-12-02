def safe?(report)
  is_increasing = true

  report.each_cons(2).each_with_index do |pair, i|
    n, m = pair
    is_increasing = n < m if i == 0
    difference = (n - m).abs
    bad_trend = is_increasing ? n >= m : n <= m
    bad_gap = difference < 1 || difference > 3
    return false if bad_trend || bad_gap
  end

  true
end

def salvageable?(report)
  report.each_index do |i|
    new_report = report.dup
    new_report.delete_at i
    return true if safe? new_report
  end

  false
end

answer1 = 0
answer2 = 0

ARGF.each_line do |line|
  report = line.scan(/\d+/).map!(&:to_i)
  is_safe = safe? report
  answer1 += 1 if is_safe
  answer2 += 1 if is_safe || salvageable?(report)
end

puts "Answer 1: #{answer1}"
puts "Answer 2: #{answer2}"

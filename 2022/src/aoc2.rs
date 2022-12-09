


fn main() {
    let lines = std::io::stdin().lines();
    let mut sum = 0;
    for line in lines {
        let input_line = line.expect("Failed to read input");
        let input_str = input_line.as_str();
        let val = match input_str {
            "A X" => 0 + 3,
            "A Y" => 3 + 1,
            "A Z" => 6 + 2,
            "B X" => 0 + 1,
            "B Y" => 3 + 2,
            "B Z" => 6 + 3,
            "C X" => 0 + 2,
            "C Y" => 3 + 3,
            "C Z" => 6 + 1,
            _ => panic!("Pattern not matched: {}", input_str),
        };
        sum += val;
    }
    dbg!(sum);
}

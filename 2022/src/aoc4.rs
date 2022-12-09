

fn parse_ints(items: &(&str, &str)) -> (i32, i32) {
    return (items.0.parse::<i32>().unwrap(), items.1.parse::<i32>().unwrap());
}

fn contains(range1: (i32, i32), range2: (i32, i32)) -> bool {
    return range1.0 >= range2.0 && range1.1 <= range2.1;
}

fn overlaps(pair: Vec<(i32, i32)>) -> bool {
    return contains(pair[0], pair[1]) || contains(pair[1], pair[0]);
}

fn overlaps_any(pair: Vec<(i32, i32)>) -> bool {
    return pair[0].0 <= pair[1].1 && pair[0].1 >= pair[1].0;
}

fn part2() {
    let lines = std::io::stdin().lines();
    let mut sum = 0;
    for line in lines {
        let input_line = line.expect("Failed to read input");
        let pairs = input_line.split(",");
        let pairs_of_ranges: Vec<_> = pairs.map(|x| x.split_once("-").expect("Invalid input")).collect();
        let ranges: Vec<_> = pairs_of_ranges.iter().map(parse_ints).collect();
        if overlaps_any(ranges) {
            sum += 1;
        }
    }
    dbg!(sum);
}

fn part1() {
    let lines = std::io::stdin().lines();
    let mut sum = 0;
    for line in lines {
        let input_line = line.expect("Failed to read input");
        let pairs = input_line.split(",");
        let pairs_of_ranges: Vec<_> = pairs.map(|x| x.split_once("-").expect("Invalid input")).collect();
        let ranges: Vec<_> = pairs_of_ranges.iter().map(parse_ints).collect();
        if overlaps(ranges) {
            sum += 1;
        }
    }
    dbg!(sum);
}

fn main() {
    part2();
}
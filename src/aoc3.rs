use std::collections::HashSet;


fn part2() {
    let lines = std::io::stdin().lines();
    let mut sum = 0;
    let mut array = [String::from(""), String::from(""), String::from("")];
    for line in lines.enumerate() {
        let (line_num, input_line) = line;
        let input_line = input_line.expect("Failed to read input");
        array[line_num % 3] = input_line;
        if line_num % 3 == 2 {
            let first_set: HashSet<_> = array[0].chars().into_iter().collect();
            let second_set: HashSet<_> = array[1].chars().into_iter().collect();
            let third_set: HashSet<_> = array[2].chars().into_iter().collect();

            let inter_set: HashSet<_>  = first_set.intersection(&second_set).cloned().collect();
            let inter_set: HashSet<_>  = inter_set.intersection(&third_set).cloned().collect();

            assert_eq!(inter_set.len(), 1);
            let common_char = inter_set.iter().next().unwrap().clone();
            let common_char_ord: u32 = common_char.into();
            sum += match common_char_ord {
                97..=122 => common_char_ord - 96,
                65..=90 => common_char_ord - 65 + 27,
                _ => panic!("Value not matched: {}", common_char_ord),
            }
        }
    }
    dbg!(sum);

}

fn part1() {
    let lines = std::io::stdin().lines();
    let mut sum = 0;
    for line in lines {
        let input_line = line.expect("Failed to read input");
        let middle = input_line.len() / 2;
        let first = &input_line[..middle];
        let second = &input_line[middle..];
        let first_set: HashSet<_> = first.chars().into_iter().collect();
        let second_set: HashSet<_> = second.chars().into_iter().collect();
        let inter_set: HashSet<_>  = first_set.intersection(&second_set).collect();
        let common_char = inter_set.iter().next().unwrap().clone();
        let common_char_ord: u32 = (*common_char).into();
        sum += match common_char_ord {
            97..=122 => common_char_ord - 96,
            65..=90 => common_char_ord - 65 + 27,
            _ => panic!("Value not matched: {}", common_char_ord),
        }
    }
    dbg!(sum);
}


fn main() {
    part2();
}
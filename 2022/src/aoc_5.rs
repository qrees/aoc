use std::{io};


// 00123456789
// [Q]         [N]             [N]
// [H]     [B] [D]             [S] [M]
// [C]     [Q] [J]         [V] [Q] [D]
// [T]     [S] [Z] [F]     [J] [J] [W]
// [N] [G] [T] [S] [V]     [B] [C] [C]
// [S] [B] [R] [W] [D] [J] [Q] [R] [Q]
// [V] [D] [W] [G] [P] [W] [N] [T] [S]
// [B] [W] [F] [L] [M] [F] [L] [G] [J]
//  1   2   3   4   5   6   7   8   9

fn update_stacks(stacks: &mut Vec<Vec<char>>, count: usize, from: usize, to: usize) {
    let mut tmp: Vec<char> = vec![];
    for i in 0..count {
        let char_to_move = stacks[from].pop().expect("No more elements to pop");
        tmp.push(char_to_move);
    }
    for i in 0..count {
        let char_to_move = tmp.pop().expect("No more elements to pop");
        stacks[to].push(char_to_move);
    }
}

fn main() {

    let mut stacks: Vec<Vec<char>> = vec![];
    for line in io::stdin().lines() {
        let line = line.expect("Failed to read a line");
        if line.len() == 0 {
            break;
        }
        // stacks.push(&line);
        let stack_count = (line.len() + 1) / 4;
        if stack_count > stacks.len() {
            stacks.resize(stack_count, vec![]);
        }
        for (pos, char) in line.chars().enumerate() {
            let stack_num = match char {
                'A'..='Z' => (pos - 1) / 4,
                '[' => continue,
                ']' => continue,
                ' ' => continue,
                '0'..='9' => continue,
                _ => panic!("Invalid character {}", char),
            };
            stacks[stack_num].push(char);
        }
    }
    for stack in stacks.iter_mut() {
        stack.reverse();
    }
    dbg!(&stacks);

    for line in io::stdin().lines() {
        let line = line.expect("Failed to read a line");
        let parts: Vec<_> = line.split(" ").collect();
        let count = parts[1].parse::<usize>().expect("Not in integer");
        let from = parts[3].parse::<usize>().expect("Not in integer") - 1;
        let to = parts[5].parse::<usize>().expect("Not in integer") - 1;
        update_stacks(&mut stacks, count, from, to);
    }
    dbg!(&stacks);

}

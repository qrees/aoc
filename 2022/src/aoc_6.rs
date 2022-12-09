use std::collections::{HashSet};
use std::io;


fn main() {
    let mut line = String::from("");
    io::stdin().read_line(&mut line).expect("Failed to read input");

    for (pos, char) in line.chars().enumerate() {
        let part = &line[pos..(pos+14)];
        let set: HashSet<_> = part.chars().into_iter().collect();
        if set.len() == 14 {
            dbg!(pos+14);
            break;
        }
    }
}
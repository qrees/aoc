use std::{io};
use std::collections::BinaryHeap;

fn main() {
    let mut max = 0;
    let mut cur = 0;


    let mut heap = BinaryHeap::new();

    for line in io::stdin().lines() {
        let line = line.expect("Failed to read a line");
        let line = line.trim();
        if line.len() == 0 {
            heap.push(cur);
            if cur > max {
                max = cur;
            }
            cur = 0;
            continue;
        }
        let value = line.parse::<i32>().expect("Not in integer");
        cur += value;
    }
    heap.push(cur);
    let total = heap.pop().unwrap() + heap.pop().unwrap() + heap.pop().unwrap();
    println!("{}", total);
}

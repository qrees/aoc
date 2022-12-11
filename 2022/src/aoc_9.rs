use std::{io, cmp::{max, min}};


fn unit(integer: i32) -> i32 {
    if integer == 0 {
        return 0
    }
    if integer > 0 {
        return 1
    }
    if integer < 0 {
        return -1;
    }
    panic!("This should not happen");
}

fn calc_tail(cur_x: usize, cur_y: usize, tail_x: usize, tail_y: usize) -> (usize, usize) {
    let mut new_tail_x = tail_x;
    let mut new_tail_y = tail_y;
    let diff_x: i32 = cur_x as i32 - tail_x as i32;
    let diff_y: i32 = cur_y as i32 - tail_y as i32;
    if diff_x.abs() >= 2 || diff_y.abs() >= 2 {
        new_tail_x = (tail_x as i32 + unit(diff_x)) as usize;
        new_tail_y = (tail_y as i32 + unit(diff_y)) as usize;
    }
    return (new_tail_x, new_tail_y);
}

struct Pos {
    pub x: usize,
    pub y: usize,
}

fn main() {
    let mut cur_x = 0i32;
    let mut cur_y = 0i32;
    let mut min_x = 0i32;
    let mut max_x = 0i32;
    let mut min_y = 0i32;
    let mut max_y = 0i32;

    let lines: Vec<_> = io::stdin().lines().collect();
    for line in lines.iter() {
        let input_line = line.as_ref().expect("Input failed");
        let (dir, steps) = input_line.split_once(" ").expect("Cannot split");
        let steps: i32 = steps.parse().expect("Not an integer");

        match dir {
            "R" => cur_x += steps,
            "L" => cur_x -= steps,
            "U" => cur_y -= steps,
            "D" => cur_y += steps,
            _ => panic!("Invalid direction")
        }

        min_x = min(min_x, cur_x);
        max_x = max(max_x, cur_x);
        min_y = min(min_y, cur_y);
        max_y = max(max_y, cur_y);
    }

    dbg!(min_x);
    dbg!(max_x);
    dbg!(min_y);
    dbg!(max_y);

    let mut state: Vec<Vec<bool>> = vec![vec![false; (max_x - min_x + 1) as usize]; (max_y - min_y + 1) as usize];
    let mut cur_x = -min_x as usize;
    let mut cur_y = -min_y as usize;
    let mut tail: Vec<Pos> = vec![];

    for i in 0..10 {
        tail.push(Pos{x: cur_x, y: cur_y});
    }

    state[cur_y][cur_x] = true;
    let mut touched = 1;
    for line in lines.iter() {
        let input_line = line.as_ref().expect("Input failed");
        let (dir, steps) = input_line.split_once(" ").expect("Cannot split");
        let steps = steps.parse::<usize>().expect("Not an integer");

        for _ in 0..steps {
            match dir {
                "R" => tail[0].x += 1,
                "L" => tail[0].x -= 1,
                "U" => tail[0].y -= 1,
                "D" => tail[0].y += 1,
                _ => panic!("Invalid direction")
            }
            for i in 0..(tail.len() - 1) {
                (tail[i+1].x, tail[i+1].y) = calc_tail(tail[i].x, tail[i].y, tail[i+1].x, tail[i+1].y);
            }
            let last = tail.last().expect("No last element");
            if state[last.y][last.x] == false {
                touched += 1;
                state[last.y][last.x] = true;
            }
        }
    }
    // for i in state {
    //     for j in i {
    //         print!("{}", if j { "X" } else { "." });
    //     }
    //     println!("");
    // }
    dbg!(touched);
}
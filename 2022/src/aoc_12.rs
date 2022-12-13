use std::io;
use std::cmp::Ordering;
use std::collections::BinaryHeap;

// struct Graph {
//     map: Vec<Vec<u8>>,
//     dist: Vec<Vec<i32>>,
//     cur_x: usize,
//     cur_y: usize,
// }


// fn step(graph: &Graph) {

// }

#[derive(Copy, Clone, Eq, PartialEq)]
struct State {
    cost: i32,
    position: (usize, usize),
}

// The priority queue depends on `Ord`.
// Explicitly implement the trait so the queue becomes a min-heap
// instead of a max-heap.
impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        // Notice that the we flip the ordering on costs.
        // In case of a tie we compare positions - this step is necessary
        // to make implementations of `PartialEq` and `Ord` consistent.
        other.cost.cmp(&self.cost)
            .then_with(|| self.position.cmp(&other.position))
    }
}

// `PartialOrd` needs to be implemented as well.
impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn main() {
    let mut map: Vec<Vec<u8>> = vec![];
    let mut dist: Vec<Vec<i32>> = vec![];
    let mut y = 0;
    let mut start_x: usize = 0;
    let mut start_y: usize = 0;
    let mut end_x: usize = 0;
    let mut end_y: usize = 0;
    for line in io::stdin().lines() {
        let mut chars: Vec<_> = line.unwrap().bytes().collect();
        let mut found_start = false;
        let mut found_end = false;
        for (x, field) in chars.iter().enumerate() {
            if (*field) == 'S' as u8 {
                start_x = x;
                start_y = y;
                found_start = true;
            }
            if (*field) == 'E' as u8 {
                end_x = x;
                end_y = y;
                found_end = true;
            }
        }
        if found_start {
            chars[start_x as usize] = 'a' as u8;
        }
        if found_end {
            chars[end_x as usize] = 'z' as u8;
        }
        let chars_len = chars.len();
        map.push(chars);
        dist.push(vec![i32::MAX; chars_len]);
        y += 1;
    }


    dbg!(end_x, end_y);
    let mut pq = BinaryHeap::new();
    let mut cur_x = start_x;
    let mut cur_y = start_y;
    pq.push(State{cost: 0, position:(cur_x, cur_y)});
    dist[cur_y][cur_x] = 0;

    dbg!(map[23][161] as i32 - map[24][161] as i32);
    dbg!(map[23][161] as char);
    dbg!(map[23][161] as char);

    while ! pq.is_empty() {
        let state = pq.pop().unwrap();
        let cur_x = state.position.0 as usize;
        let cur_y = state.position.1 as usize;
        let cur_dist = dist[cur_y][cur_x];

        for (dx, dy) in [(1, 0), (0, 1), (-1, 0), (0, -1)] {
            let new_x: i32 = state.position.0 as i32 + dx;
            let new_y: i32 = state.position.1 as i32 + dy;
            if (new_x == end_x as i32) && (new_y == end_y as i32) {
                dbg!((new_x, new_y, cur_dist));
            }
            if new_x < 0 { continue; };
            if new_x >= map[0].len() as i32 { continue; };
            if new_y < 0 { continue; };
            if new_y >= map.len() as i32 { continue; };
            let new_x = new_x as usize;
            let new_y = new_y as usize;

            // dbg!((map[cur_y][cur_x] as i32 - map[new_y][new_x] as i32));
            if (map[cur_y][cur_x] as i32 - map[new_y][new_x] as i32) > -2 {

            }else {
                continue;
            }

            let alt = cur_dist + 1;
            if alt < dist[new_y][new_x] {
                dist[new_y][new_x] = alt;
                pq.push(State{cost: alt, position:(new_x, new_y)});
            }
        }
    }

    // for (y, row) in dist.iter().enumerate() {
    //     for (x, col) in row.iter().enumerate() {
    //         if (*col) < i32::MAX {
    //             print!("!{}", map[y][x] as char);
    //         } else {
    //             print!(" {}", map[y][x] as char);
    //         }
    //     }
    //     println!("");
    // }
    // dbg!(&dist[end_y]);
    // dbg!(&map);
    // dbg!((end_x, end_y));
    dbg!(dist[end_y as usize][end_x as usize]);
}
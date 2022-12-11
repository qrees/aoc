use std::io;


fn traverse(matrix: &Vec<Vec<i32>>, visible: &mut Vec<Vec<bool>>, dir_x: i32, dir_y: i32) -> u32 {
    let mut cur_x: i32 = if dir_x >= 0 { 0 } else {(matrix.len() - 1).try_into().unwrap()};
    let mut cur_y: i32 = if dir_y >= 0 { 0 } else {(matrix[0].len() - 1).try_into().unwrap()};
    let mut prev_val: i32 = -1;
    let mut count = 0;
    loop {
        let cur_val: i32 = matrix[cur_x as usize][cur_y as usize];
        if cur_val > prev_val {
            if ! visible[cur_x as usize][cur_y as usize] {
                count += 1;
            }
            visible[cur_x as usize][cur_y as usize] = true;
            prev_val = cur_val;
        }
        cur_x = cur_x + dir_x;
        cur_y = cur_y + dir_y;
        if cur_x < 0 {
            cur_x = if dir_x >= 0 { 0 } else {(matrix.len() - 1).try_into().unwrap()};
            cur_y += 1;
            prev_val = -1;
        } else if cur_x >= visible.len() as i32 {
            cur_x = if dir_x >= 0 { 0 } else {(matrix.len() - 1).try_into().unwrap()};
            cur_y += 1;
            prev_val = -1;
        } else if cur_y >= visible[cur_x as usize].len() as i32 {
            cur_y = if dir_y >= 0 { 0 } else {(matrix[0].len() - 1).try_into().unwrap()};
            cur_x += 1;
            prev_val = -1;
        }else if cur_y < 0 {
            cur_y = if dir_y >= 0 { 0 } else {(matrix[0].len() - 1).try_into().unwrap()};
            cur_x += 1;
            prev_val = -1;
        }
        if cur_x >= visible.len() as i32 || cur_y >= visible[cur_x as usize].len() as i32 {
            break
        }
    }
    return count;
}

fn scenic_score(matrix: &Vec<Vec<i32>>, loc_x: usize, loc_y: usize) -> usize {
    let size_x = matrix.len();
    let size_y = matrix[0].len();
    let start_h = matrix[loc_x][loc_y];
    let mut edge_l = loc_x;
    let mut edge_r = loc_x;
    let mut edge_t = loc_y;
    let mut edge_b = loc_y;

    for x in (loc_x + 1)..size_x {
        let cur_h = matrix[x][loc_y];
        edge_r = x;
        if cur_h >= start_h {
            break;
        }
    }
    for x in (0..(loc_x)).rev() {
        let cur_h = matrix[x][loc_y];
        edge_l = x;
        if cur_h >= start_h {
            break;
        }
    }
    for y in (loc_y + 1)..size_y {
        let cur_h = matrix[loc_x][y];
        edge_b = y;
        if cur_h >= start_h {
            break;
        }
    }
    for y in (0..(loc_y)).rev() {
        let cur_h = matrix[loc_x][y];
        edge_t = y;
        if cur_h >= start_h {
            break;
        }
    }
    // dbg!(loc_x, loc_y);
    // dbg!(edge_r, edge_l, edge_b, edge_t);
    let dist = (edge_r - loc_x) * (loc_x - edge_l) * (edge_b - loc_y) * (loc_y - edge_t);
    return dist;
}

fn main() {
    let mut matrix: Vec<Vec<i32>> = vec![];
    let mut visible: Vec<Vec<bool>> = vec![];

    for line in io::stdin().lines() {
        let input_line = line.expect("Input failed");
        let chars: Vec<i32> = input_line.chars().map(|x| x.to_digit(10).expect("Not in integer") as i32).collect();
        let vis_line = vec![false; chars.len()];
        matrix.push(chars);
        visible.push(vis_line);
    }

    let mut max_score = 0;
    for x in 0..matrix.len() {
        for y in 0..matrix[x].len() {
            let score = scenic_score(&matrix, x, y);
            if score > max_score {
                max_score = score;
            }
        }
    }
    dbg!(max_score);
}
use std::io;


fn main() {

    let mut x = 1;
    let mut tick = 0;
    let mut value;
    let mut pre_x = x;
    let mut step = 20;
    let mut total = 0;
    for line in io::stdin().lines() {
        let input_line = line.unwrap();
        match input_line.as_str() {
            "noop" => tick += 1,
            _ => {
                (_, value) = input_line.split_once(" ").unwrap();
                let value: i32 = value.parse().unwrap();
                tick += 2;
                x += value;
            }
        }
        if tick >= step {
            // dbg!(pre_x * step);
            total += pre_x * step;
            step += 40;
        }
        pre_x = x;
    }
    dbg!(total);
}

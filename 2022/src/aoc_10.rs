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
            "noop" => {
                if (tick) >= (x - 1) && (tick) <= (x + 1) {
                    print!("#");
                } else {
                    print!(" ");
                }
                tick = tick + 1;
                if tick >= 40 {
                    tick = 0;
                    println!("");
                }
            }
            _ => {
                (_, value) = input_line.split_once(" ").unwrap();
                let value: i32 = value.parse().unwrap();
                if (tick) >= (x - 1) && (tick) <= (x + 1) {
                    print!("#");
                } else {
                    print!(" ");
                }
                tick = tick + 1;
                if tick >= 40 {
                    tick = 0;
                    println!("");
                }
                if (tick) >= (x - 1) && (tick) <= (x + 1) {
                    print!("#");
                } else {
                    print!(" ");
                }
                tick = tick + 1;
                if tick >= 40 {
                    tick = 0;
                    println!("");
                }
                x += value;
            }
        }
    }
    dbg!(total);
}

use std::io;


fn main() {

    let mut X = 1;
    let mut tick = 0;
    for line in io::stdin().lines() {
        let input_line = line.unwrap();
        match input_line.as_str() {
            "noop" => tick += 1,
            _ => {
                (_, value) = input_line.split_once(" ").unwrap();
            }
        }
    }
}
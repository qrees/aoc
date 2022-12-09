use core::panic;
use std::{io, collections::{HashMap, hash_map}, rc::Rc};


struct TreeNode<'a> {

    directories: HashMap<&'a str, Rc<TreeNode<'a>>>,
    files_size: u32,
    parent: Option<Rc<TreeNode<'a>>>
}

impl<'a> TreeNode<'a> {
    pub fn new(cur: Rc<TreeNode<'a>>) -> Self {
        TreeNode { directories: HashMap::new(), files_size: 0, parent: Some(cur.clone()) }
    }
}

fn main() {
    let mut lines = io::stdin().lines();
    // let mut input_line2;
    let mut input_line: String;

    let mut line = lines.next().expect(" ");
    let mut input_line = line.expect("");
    let mut root = Rc::new(TreeNode { directories: HashMap::new(), parent: None, files_size: 0});
    let mut cur = root.clone();
    loop {

        if let '$' = input_line.chars().nth(0).unwrap() {
            let parts: Vec<_> = input_line.split(" ").collect();
            if parts[1] == "cd" {
                if parts[2] == ".." {
                    cur = cur.parent.expect("No parent");
                } else {
                    cur = cur.directories.get(parts[2]).expect(format!("Unknown directory {}", parts[2]).as_str()).clone();
                }
                input_line = lines.next().expect(" ").expect(" ");
                continue;
            } else if parts[1] == "ls" {
                loop {
                    if let '$' = input_line.chars().nth(0).unwrap() {
                        break;
                    }
                    let parts = input_line.split_once(" ").expect("Cannot split");
                    if parts.0 == "dir" {
                        let new_node = Rc::new(TreeNode::new(cur));
                        cur.directories.insert(parts.1, new_node.clone());
                    }
                    if parts.0.parse::<u32>().is_ok() {
                        let size = parts.0.parse::<u32>().expect("Not a number");
                        cur.files_size += size;
                    }
                    input_line = lines.next().expect(" ").expect(" ");
                    continue;
                }
                continue;
            }
        }
    }
}
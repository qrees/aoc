use core::panic;
use std::{io, collections::{HashMap, hash_map}};

use std::cell::RefCell;
use std::rc::Rc;


struct TreeNode {
    pub directories: HashMap<String, Rc<RefCell<TreeNode>>>,
    pub files_size: u32,
    pub parent: Option<Rc<RefCell<TreeNode>>>
}

impl TreeNode {
    pub fn new(cur: Rc<RefCell<TreeNode>>) -> Self {
        TreeNode { directories: HashMap::new(), files_size: 0, parent: Some(cur.clone()) }
    }
}

fn main() {
    let mut lines = io::stdin().lines();
    // let mut input_line2;
    let mut input_line: String;

    let mut line = lines.next().expect(" ");
    let mut input_line = line.expect("");
    let root = Rc::new(RefCell::new(TreeNode { directories: HashMap::new(), parent: None, files_size: 0}));
    let mut cur = Rc::clone(&root);
    loop {
        let cloned = Rc::clone(&cur);
        if let '$' = input_line.chars().nth(0).unwrap() {
            let parts: Vec<_> = input_line.split(" ").collect();
            if parts[1] == "cd" {
                if parts[2] == ".." {
                    // let cur_clone = Rc::clone(&cur);
                    cur = Rc::clone(Rc::clone(&cur).borrow().parent.as_ref().unwrap());
                    // cur = Rc::clone(Rc::clone(&cur).borrow().parent.as_ref().unwrap());
                    // cur = cur.borrow_mut().parent.expect("No parent").clone();
                } else {
                    cur = Rc::clone(&cur).borrow_mut().directories.get(parts[2]).expect(format!("Unknown directory {}", parts[2]).as_str()).clone();
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
                        let new_node = Rc::new(RefCell::new(TreeNode::new(Rc::clone(&cur))));
                        Rc::clone(&cur).borrow_mut().directories.insert(parts.1.to_string(), new_node.clone());
                    }
                    if parts.0.parse::<u32>().is_ok() {
                        let size = parts.0.parse::<u32>().expect("Not a number");
                        cur.borrow_mut().files_size += size;
                    }
                    input_line = lines.next().expect(" ").expect(" ");
                    continue;
                }
                continue;
            }
        }
    }
}
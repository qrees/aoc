use std::{io, collections::{HashMap}};

use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug)]
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

    let root = Rc::new(RefCell::new(TreeNode { directories: HashMap::new(), parent: None, files_size: 0}));
    let mut cur = root.clone();

    let maybe_line = lines.next();
    if maybe_line.is_none() {
        return;
    }
    let mut input_line = maybe_line.expect("first fail").expect("second fail");
    let mut total = 0;
    'outer: loop {
        let cloned = Rc::clone(&cur);
        if let '$' = input_line.chars().nth(0).unwrap() {
            let parts: Vec<_> = input_line.split(" ").collect();
            if parts[1] == "cd" {
                if parts[2] == "/" {

                } else
                if parts[2] == ".." {
                    let ret_size = cur.borrow_mut().files_size;
                    if ret_size <= 100000 {
                        total += cur.borrow_mut().files_size;
                    }
                    cur = Rc::clone(cloned.borrow().parent.as_ref().unwrap());
                    cur.borrow_mut().files_size += ret_size;
                } else {
                    cur = cloned.borrow_mut().directories.get(parts[2]).expect(format!("Unknown directory {}", parts[2]).as_str()).clone();
                }
            } else if parts[1] == "ls" {
                let maybe_line = lines.next();
                if maybe_line.is_none() {
                    break 'outer;
                }
                input_line = maybe_line.expect("First fali").expect("Second fail");
                loop {
                    if input_line.chars().nth(0).unwrap() == '$' {
                        continue 'outer;
                    }
                    let parts = input_line.split_once(" ").expect("Cannot split");
                    if parts.0 == "dir" {
                        let new_node = Rc::new(RefCell::new(TreeNode::new(Rc::clone(&cur))));
                        cur.borrow_mut().directories.insert(parts.1.to_string(), new_node.clone());
                    }
                    if parts.0.parse::<u32>().is_ok() {
                        let size = parts.0.parse::<u32>().expect("Not a number");
                        cur.borrow_mut().files_size += size;
                    }

                    let maybe_line = lines.next();
                    if maybe_line.is_none() {
                        break 'outer;
                    }
                    input_line = maybe_line.expect("First fali").expect("Second fail");
                    continue;
                }
            }
        }
        let maybe_line = lines.next();
        if maybe_line.is_none() {
            break;
        }
        input_line = maybe_line.expect("First fali").expect("Second fail");
        continue;
    }

    while cur.borrow().parent.is_some() {
        let cloned = Rc::clone(&cur);
        let ret_size = cur.borrow_mut().files_size;
        if ret_size <= 100000 {
            total += cur.borrow_mut().files_size;
        }
        cur = Rc::clone(cloned.borrow().parent.as_ref().unwrap());
        cur.borrow_mut().files_size += ret_size;
    }

    dbg!(total);
}
use std::{collections::VecDeque};

use std::cell::RefCell;
use std::rc::Rc;

struct Monkey {
    pub items: VecDeque<u32>,
    pub op: fn(old: u32) -> u32,
    pub test: u32,
    pub true_destination: usize,
    pub false_destination: usize,
}

fn main() {
    let mut monkeys: Vec<Rc<RefCell<Monkey>>> = vec![];

    let monkey: Monkey = Monkey{
        items: VecDeque::from([71, 56, 50, 73]),
        op: |old| {old * 11},
        test: 13,
        true_destination: 1,
        false_destination: 7,
    };
    monkeys.push(Rc::new(RefCell::new(monkey)));

    let monkey: Monkey = Monkey{
        items: VecDeque::from([70, 89, 82]),
        op: |old| {old + 1},
        test: 7,
        true_destination: 3,
        false_destination: 6,
    };
    monkeys.push(Rc::new(RefCell::new(monkey)));

    let monkey: Monkey = Monkey{
        items: VecDeque::from([52, 95]),
        op: |old| {old * old},
        test: 3,
        true_destination: 5,
        false_destination: 4,
    };
    monkeys.push(Rc::new(RefCell::new(monkey)));

    let monkey: Monkey = Monkey{
        items: VecDeque::from([94, 64, 69, 87, 70]),
        op: |old| {old + 2},
        test: 19,
        true_destination: 2,
        false_destination: 6,
    };
    monkeys.push(Rc::new(RefCell::new(monkey)));

    let monkey: Monkey = Monkey{
        items: VecDeque::from([98, 72, 98, 53, 97, 51]),
        op: |old| {old + 6},
        test: 5,
        true_destination: 0,
        false_destination: 5,
    };
    monkeys.push(Rc::new(RefCell::new(monkey)));

    let monkey: Monkey = Monkey{
        items: VecDeque::from([79]),
        op: |old| {old + 7},
        test: 2,
        true_destination: 7,
        false_destination: 0,
    };
    monkeys.push(Rc::new(RefCell::new(monkey)));

    let monkey: Monkey = Monkey{
        items: VecDeque::from([77, 55, 63, 93, 66, 90, 88, 71]),
        op: |old| {old * 7},
        test: 11,
        true_destination: 2,
        false_destination: 4,
    };
    monkeys.push(Rc::new(RefCell::new(monkey)));

    let monkey: Monkey = Monkey{
        items: VecDeque::from([54, 97, 87, 70, 59, 82, 59]),
        op: |old| {old + 8},
        test: 17,
        true_destination: 1,
        false_destination: 3,
    };
    monkeys.push(Rc::new(RefCell::new(monkey)));

    let mut inspect_count: Vec<_> = [0; 8].to_vec();
    // let monkeys = Rc::new(RefCell::new(monkeys));
    for _ in 0..20 {
        for i in 0..monkeys.len() {
            let mut monkey = monkeys[i].clone();
            while monkey.borrow_mut().items.len() > 0 {
                inspect_count[i] += 1;
                let item = monkey.borrow_mut().items.pop_front().unwrap();
                let new_level = (monkey.borrow().op)(item);
                let divided = new_level / 3;
                if (divided % monkey.borrow_mut().test) == 0 {
                    monkeys[monkey.borrow_mut().true_destination].borrow_mut().items.push_back(divided);
                } else {
                    monkeys[monkey.borrow_mut().false_destination].borrow_mut().items.push_back(divided);
                }
            }
        }
    }
    // let inspect_count: Vec<_> = inspect_count.iter().enumerate().collect();
    inspect_count.sort();
    inspect_count.reverse();
    // dbg!(inspect_count);
    dbg!(inspect_count[0] * inspect_count[1]);
    // dbg!()
}

package main

import (
    "fmt"
)

type Tree struct {
    left *Tree;
    right *Tree;
    value int;
}

var sum_map map[int]int = make(map[int]int);
var sum int = 0;
var total_paths int = 0;
var expected int;

func walk(cur *Tree, depth int) {
    if cur == nil {
        return
    }
    sum += cur.value;
    sum_map[sum] += 1;
    walk(cur.left, depth + 1);
    walk(cur.right, depth + 1);
    total_paths += sum_map[sum - expected];
    sum_map[sum] -= 1;
    sum -= cur.value;
}


func paths_with_sums(tree *Tree, sum int){
    sum_map[0] = 1;
    expected = sum;
    walk(tree, 0);
    fmt.Println(total_paths);
}

func main() {
    var tree * Tree = &Tree{
        left: &Tree{
            left: &Tree{
                left: &Tree{value: 3},
                right: &Tree{value: -2},
                value: 3},
            right: &Tree{
                right: &Tree{value: 1},
                value: 2},
            value: 5},
        right: &Tree{
            right: &Tree{value: 11},
            value: -3},
        value: 10};
    paths_with_sums(tree, 8);
}

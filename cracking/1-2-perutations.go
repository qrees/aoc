package main

import (
    "fmt"
)

func permutation(str1 string, str2 string) bool{
    var exists bool;
    var count int;
    var counts map[rune]int = make(map[rune]int);
    for _, char := range str1[:] {
        count, exists = counts[char];
        if (! exists){
            counts[char] = 0;
            count = 0;
        }
        counts[char] = count + 1;
    }

    for _, char := range str2[:] {
        count, exists = counts[char];
        if (! exists){
            return false;
        }
        counts[char] = count - 1;
    }

    for _, value := range counts {
        if value != 0 { return false};
    }

    return true;
}

func main(){
    var is_perm bool;
    //var char byte;

    is_perm = permutation("abc", "cbad");
    if (is_perm) {
        fmt.Println("Is perm");
    } else {
        fmt.Println("Is NOT perm");
    }
}

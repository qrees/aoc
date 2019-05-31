package main

import (
    "bufio"
    "os"
    "regexp"
    "strconv"
    "fmt"
)

const MAX int = 1000;
var lights [][MAX]int;

var input *os.File;

func turn_on(cur_x int, cur_y int){
    lights[cur_x][cur_y]++;
}

func turn_off(cur_x int, cur_y int){
    if lights[cur_x][cur_y] > 0 {
        lights[cur_x][cur_y]--;
    }
}

func toggle(cur_x int, cur_y int){
    lights[cur_x][cur_y] += 2;
}

func main(){
    var scanner *bufio.Scanner;
    var line string;
    var compilled *regexp.Regexp;
    var matches []string;
    var cur_x, cur_y int
    var fun func(int, int)

    for i := 0; i < MAX; i++ {
        lights = append(lights, [MAX]int{});
    }

    input = os.Stdin;
    scanner = bufio.NewScanner(input);

    compilled, _ = regexp.Compile("(turn on|toggle|turn off) ([0-9]*),([0-9]*) through ([0-9]*),([0-9]*)");
    for scanner.Scan() {
        line = scanner.Text();
        matches = compilled.FindStringSubmatch(line);
        fmt.Println(matches)
        if (matches[1] == "turn on") {
            fun = turn_on
        } else if (matches[1] == "turn off") {
            fun = turn_off
        } else if (matches[1] == "toggle") {
            fun = toggle
        } else {
            panic("Shit")
        }
        start_x, _ := strconv.Atoi(matches[2])
        start_y, _ := strconv.Atoi(matches[3])
        end_x, _ := strconv.Atoi(matches[4])
        end_y, _ := strconv.Atoi(matches[5])
        for cur_x = start_x; cur_x <= end_x; cur_x ++ {
            for cur_y = start_y; cur_y <= end_y; cur_y ++ {
                fun(cur_x, cur_y)
            }
        }
    }
    var counter int = 0;
    for cur_x = 0; cur_x < MAX; cur_x ++ {
        for cur_y = 0; cur_y < MAX; cur_y ++ {
            counter += lights[cur_x][cur_y];
        }
    }
    fmt.Println(counter)
}
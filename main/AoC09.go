package main

import (
    "os"
    "bufio"
    "strconv"
    "fmt"
    "regexp"
    "container/list"
)

type StringPair struct {
    from, to string;
}

var distances map[StringPair]int = make(map[StringPair]int);
var min_distance int = 0;

func traverse(cities map[string]bool, tail *list.List) {
    var found bool = false;
    for city, valid := range cities {
        if !valid {
            continue;
        }
        found = true;
        cities[city] = false;
        el := tail.PushFront(city);
        traverse(cities, tail);
        tail.Remove(el);
        cities[city] = true;
    }

    distance := 0;
    total := 0;
    if !found {
        item := tail.Front();
        start := item.Value.(string);
        var from string = start;
        fmt.Print(item.Value, " ");
        item = item.Next();
        for item != nil {
            fmt.Print(item.Value, " ");
            var to string = item.Value.(string);
            distance = distances[StringPair{from, to}];
            total += distance;
            from = to;
            item = item.Next();
        }
        if total > min_distance {
            min_distance = total;
        }
        fmt.Println(min_distance);
    }
}

func main(){
    var input *os.File;
    var scanner *bufio.Scanner;
    var line string;
    //var unquoted string;
    var err error;
    var matches []string;
    var r_distance *regexp.Regexp;
    var from, to string;
    var distance int;
    var cities map[string]bool = make(map[string]bool);

    // Tristram to AlphaCentauri = 34
    r_distance, _ = regexp.Compile(`^(\w*) to (\w*) = (\d*)$`);

    input = os.Stdin;
    scanner = bufio.NewScanner(input);

    for scanner.Scan() {
        line = scanner.Text();

        matches = r_distance.FindStringSubmatch(line);

        from = matches[1];
        to = matches[2];
        cities[from] = true;
        cities[to] = true;
        distance, err = strconv.Atoi(matches[3]);
        if err != nil {
            panic(fmt.Sprintf("Not a number %s", matches[3]));
        }
        distances[StringPair{from, to}] = distance;
        distances[StringPair{to, from}] = distance;
    }
    traverse(cities, list.New());
    println(distance);
}

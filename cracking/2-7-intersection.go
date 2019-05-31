package main

type List struct {
    next *List;
}

func min(a int, b int) int {
    if a < b {
        return a;
    }
    return b;
}

func max(a int, b int) int {
    if a > b {
        return a;
    }
    return b;
}

func intersect(list1 List, list2 List) *List {
    var last_item1 List;
    var last_item2 List;
    var count1 = 0;
    var count2 = 0;

    last_item1 = list1;
    for ; last_item1.next != nil; last_item1 = *last_item1.next {
        count1 ++;
    }

    last_item2 = list2;
    for ; last_item2.next != nil; last_item2 = *last_item2.next {
        count2 ++;
    }

    if last_item2 != last_item1 {
        return nil;
    }
    var skip = max(count1, count2) - min(count1, count2);

    if count2 > count1 {
        list2, list1 = list1, list2;
    }

    for ; skip > 0; skip-- {
        list1 = *list1.next;
    }

    for ; list1 != list2; {
        list1 = *list1.next;
        list2 = *list2.next;
    }
    return &list1;
}

func main(){

}

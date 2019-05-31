package main

import (
    "github.com/golang-collections/collections/stack"
    "fmt"
)

func sort(input *stack.Stack) *stack.Stack {
    var output *stack.Stack = stack.New();
    var item int;
    var peeked int;
    var moved int;

    for ; input.Len() > 0; {
        item = input.Pop().(int);
        moved = 0;
        for ; ; {
            if output.Len() == 0 {
                output.Push(item);
                break;
            } else {
                peeked = output.Peek().(int);
                if (peeked < item) {
                    output.Push(item);
                    break;
                } else {
                    output.Pop();
                    input.Push(peeked);
                    moved += 1;
                }
            }
        }
        for ; moved > 0 ; moved -= 1 {
            output.Push(input.Pop());
        }
    }
    return output;
}

func main(){
    var input *stack.Stack = stack.New();
    input.Push(3);
    input.Push(1);
    input.Push(2);
    var output *stack.Stack = sort(input);
    for ; output.Len() > 0 ; {
        fmt.Println(output.Pop().(int));
    }
}

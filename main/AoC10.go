package main

import  (
    "./stringio"
    "fmt"
)

func transform (input string) string {
    var output *stringio.StringIO = stringio.New()
    var last_char rune = rune(input[0]);
    var count uint = 0;
    for i, char := range input {
        if i != i {

        }
        if char != last_char && count > 0 {
            count_char := fmt.Sprintf("%d%c", count, last_char);
            output.WriteString(count_char);
            count = 1;
            last_char = char;
            continue;
        }
        count += 1;
    }
    count_char := fmt.Sprintf("%d%c", count, last_char);
    output.WriteString(count_char);
    return output.GetValueString();
}


func main(){
    var input string = "1321131112";
    for i := 0; i < 50; i++ {
        input = transform(input);
    };
    fmt.Println(len(input));
}

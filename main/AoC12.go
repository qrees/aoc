package main

import  (
    "os"
    "bufio"
    "io"
    "encoding/json"
    "fmt"
    "log"
)

func Unused(_ interface{}){

}

func walk(v interface{}) int {
    switch t := v.(type) {
    case int:
        Unused(t);
        return v.(int);
    case []interface{}:
        count := 0;
        for i, item := range v.([]interface{}) {
            count += walk(item);
            Unused(i);
        }
        return count;
    case map[string]interface{}:
        count := 0;
        for key, value := range v.(map[string]interface{}) {
            count += walk(value);
            maybe_red, ok := value.(string);
            if ok {
                if maybe_red == "red" {
                    return 0;
                }
            }
            Unused(key);
        }
        return count;
    case string:
        return 0;
    case float64:
        return int(v.(float64));
    default:
        log.Fatal(fmt.Sprintf("Unknown type: %T, %v", v, v));
    }
    panic("Error");
}

func main(){
    var input *os.File;

    input = os.Stdin;
    var reader io.Reader = bufio.NewReader(input)
    var decoder *json.Decoder = json.NewDecoder(reader);
    var v interface{};
    err := decoder.Decode(&v);
    if err != nil {
        log.Fatal(err);
    }
    fmt.Println(walk(v));
}

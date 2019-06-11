package main

import (
    "math/big"
    "fmt"
)

func main(){
    var value *big.Int = big.NewInt(20151125);
    var multipler *big.Int = big.NewInt(252533);
    var modulo *big.Int = big.NewInt(33554393);

    var row int64 = 2978;
    var column int64 = 3083;
    var diagonal int64 = row + column - 1;
    fmt.Println(diagonal);
    var rowSum int64 = (diagonal * (diagonal + 1)) / 2 - row;
    fmt.Println(rowSum);
    var exp *big.Int = big.NewInt(rowSum);

    var result *big.Int = big.NewInt(0);
    result.Exp(multipler, exp, modulo);
    fmt.Println(result);
    result.Mul(value, result);
    fmt.Println(result);
    result.Mod(result, modulo);
    fmt.Println(result);
}

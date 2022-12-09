package main

import (
    "regexp"
    "os"
    "bufio"
    "strconv"
    "fmt"
)

type Operator interface {
    getValue() uint16;
}

var operators map[string]Operator;
var cache map[string]uint16 = make(map[string]uint16);

func getValue(node string) uint16 {
    value, valid := cache[node];
    if !valid {
        value = operators[node].getValue();
        cache[node] = value;
    }
    return value
}

type ConstOp struct {
    value uint16;
}

func (op ConstOp) getValue() uint16 {
    return op.value;
}

type DirectOp struct {
    input string;
}

func (op DirectOp) getValue() uint16 {
    return getValue(op.input);
}

type AndOp struct {
    input1 string;
    input2 string;
}

func (op AndOp) getValue() uint16 {
    return getValue(op.input1) & getValue(op.input2);
}

type OrOp struct {
    input1 string;
    input2 string;
}

func (op OrOp) getValue() uint16 {
    return getValue(op.input1) | getValue(op.input2);
}

type NotOp struct {
    input string;
}

func (op NotOp) getValue() uint16 {
    return ^getValue(op.input);
}

type LshiftOp struct {
    input string;
    shift uint16;
}

func (op LshiftOp) getValue() uint16 {
    return getValue(op.input) << op.shift;
}

type RshiftOp struct {
    input string;
    shift uint16;
}

func (op RshiftOp) getValue() uint16 {
    var value uint16 = getValue(op.input);
    return value >> op.shift;
}

func atoi(text string) uint16 {
    value, err := strconv.Atoi(text);
    if err != nil {
        panic("shit")
    }
    return uint16(value);
}

func maybe_const(token string) string {
    value, err := strconv.Atoi(token);
    if err == nil {
        op := ConstOp{value: uint16(value)};
        operators[token] = op;
    };
    return token;
}

func main(){
    var input *os.File;
    var scanner *bufio.Scanner;
    var line string;
    var matches []string;
    operators = make(map[string]Operator);
    var r_direct, r_constant, r_and, r_or, r_not, r_lshift, r_rshift *regexp.Regexp;
    r_constant, _ = regexp.Compile(`^(\d+) -> ([a-z]+)$`);
    r_direct, _ = regexp.Compile(`^([a-z]+) -> ([a-z]+)`);
    r_and, _ = regexp.Compile(`^([a-z0-9]+) AND ([a-z0-9]+) -> ([a-z]+)`);
    r_or, _ = regexp.Compile(`^([a-z0-9]+) OR ([a-z0-9]+) -> ([a-z]+)`);
    r_not, _ = regexp.Compile(`^NOT ([a-z0-9]+) -> ([a-z]+)`);
    r_lshift, _ = regexp.Compile(`^([a-z0-9]+) LSHIFT (\d+) -> ([a-z]+)`);
    r_rshift, _ = regexp.Compile(`^([a-z0-9]+) RSHIFT (\d+) -> ([a-z]+)`);

    input = os.Stdin;
    scanner = bufio.NewScanner(input);

    for scanner.Scan() {
        line = scanner.Text();
        fmt.Println(line);
        matches = r_constant.FindStringSubmatch(line);
        if len(matches) > 0 {
            var op Operator = ConstOp{value: atoi(matches[1])};
            operators[matches[2]] = op;
            continue;
        }
        matches = r_direct.FindStringSubmatch(line);
        if len(matches) > 0 {
            var op Operator = DirectOp{input: maybe_const(matches[1])};
            operators[matches[2]] = op;
            continue;
        }
        matches = r_and.FindStringSubmatch(line);
        if len(matches) > 0 {
            var op Operator = AndOp{input1: maybe_const(matches[1]), input2: maybe_const(matches[2])};
            operators[matches[3]] = op;
            continue;
        }
        matches = r_or.FindStringSubmatch(line);
        if len(matches) > 0 {
            var op Operator = OrOp{input1: maybe_const(matches[1]), input2: maybe_const(matches[2])};
            operators[matches[3]] = op;
            continue;
        }
        matches = r_not.FindStringSubmatch(line);
        if len(matches) > 0 {
            var op Operator = NotOp{input: maybe_const(matches[1])};
            operators[matches[2]] = op;
            continue;
        }
        matches = r_lshift.FindStringSubmatch(line);
        if len(matches) > 0 {
            var op Operator = LshiftOp{input: maybe_const(matches[1]), shift: atoi(matches[2])};
            operators[matches[3]] = op;
            continue;
        }
        matches = r_rshift.FindStringSubmatch(line);
        if len(matches) > 0 {
            var op Operator = RshiftOp{input: maybe_const(matches[1]), shift: atoi(matches[2])};
            operators[matches[3]] = op;
            continue;
        }
        fmt.Println(line);
        panic("shit")
    }
    cache["b"] = 3176;
    fmt.Println(operators["a"].getValue())
}

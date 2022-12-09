
END = object()

class Parser:

    def __init__(self, exp):
        self.exp = exp + [END]
        self.index = 0

    def peek(self):
        return self.exp[self.index]

    def accept(self):
        try:
            return self.peek()
        finally:
            self.index += 1

    def factor(self):
        if isinstance(self.peek(), int):
            return self.accept()
        if self.peek() == "(":
            self.accept()
            addition = self.addition()
            return addition
        if self.peek() == "-":
            self.accept()
            val = self.factor()
            return -val
        raise ValueError(self.peek())
    
    def mul(self):
        lval = self.factor()
        while self.peek() == "*" or self.peek() == "/":
            operator = self.accept()
            val = self.factor()
            if operator == "*":
                lval *= val
            elif operator == "/":
                lval /= val
        return lval

    def addition(self):
        lval = self.mul()
        while self.peek() == "+" or self.peek() == "-" or self.peek() == ")":
            operator = self.accept()
            if operator == ")":
                return lval
            val = self.mul()
            if operator == "+":
                lval += val
            elif operator == "-":
                lval -= val
            else:
                raise ValueError(operator)
        return lval


NUMBERS = "1234567890"

def tokenize(expression):
    tokens = []
    cur_val = ""
    for char in expression:
        if char == " ":
            if cur_val:
                tokens.append(int(cur_val))
                cur_val = ""
            continue
        if char in NUMBERS:
            cur_val += char
        if char in "()*/+-":
            if cur_val:
                tokens.append(int(cur_val))
                cur_val = ""
            tokens += char
    if cur_val:
        tokens.append(int(cur_val))
    return tokens


def calc(expression):
    tokens = tokenize(expression)
    parser = Parser(tokens)
    res = parser.addition()
    print(expression, res)
    return res


print(calc("1 + 2 * 3 * (5 - 2) - 8"))

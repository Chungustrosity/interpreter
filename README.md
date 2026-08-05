# interpreter

My own implementation of the Monkey toy-programming language interpreter,
following the book Writing an Interpreter in Go.

## Features of the Monkey Language

- C-like syntax
- variable bindings
- integer and boolean built-in primitives
- arithmetic expression resolution
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hashmap data structure

## Code Examples

Here is how we bind values to names in Monkey:

```Monkey
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

Here’s what binding an array of integers to a name looks like:

```Monkey
let myArray = [1, 2, 3, 4, 5];
```

Here is a hashmap:

```Monkey
let thorsten = {"name": "Thorsten", "age": 28};
```

Accessing the elements in arrays and hashes is done with index expressions:

```Monkey
myArray[0]        // => 1
thorsten["name"]  // => "Thorsten"
```

The `let` keyword can also be used to bind functions to names. Here's a small
function that adds two number:

```Monkey
let add = fn(a, b) {return a + b; };
```

Monkey also supports implicit return values:

```Monkey
let add = fn(a, b) {a + b; };

```

Function invocation works the way you'd expect:

```Monkey
add(1, 2);
```

A more complex function involving recursion and conditional statements might
look like this:

```Monkey
let fibonacci = fn(x) {
  if (x == 0) {
    0
    } else {
      if (x == 1) {
        1
      } else {
      fibonacci(x-1) + fibonacci(x-2)
      }
  }
}
```

Monkey also supports higher-order and first-class functions. For example:

```Monkey
let applyTwice = fn(f, x) {
  return f(f(x));
};

let addTwo = fn(x){
  return x + 2;
};

applyTwice(addTwo, 2);    // => 6
```

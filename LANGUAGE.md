# Orus Language Guide

Orus is an experimental interpreted language influenced by modern scripting languages and Rust-like syntax. This guide covers the features available in version 0.5.2 and serves both as a tutorial and reference. All examples come from the `tests/` directory.

## Getting Started

A simple program prints text using the built-in `print` function:

```orus
fn main() {
    print("Hello, Orus!")
}
```

The interpreter looks for a `main` function in the entry file. Exactly one such function must exist across the project.

## Variables and Mutability

Variables are declared with `let`.

```orus
let number: i32 = 5     // immutable
let mut count = 0       // mutable, type inferred as i32
```

- **Immutability** is the default. Reassigning an immutable binding is a compile-time error.
- Use `let mut` to allow reassignment. The type of a binding cannot change after it is set.
- Variables are block scoped and must be declared inside functions.

```orus
fn demo() {
    let mut value = 1
    value = 2       // ✅ allowed
    // value = 3.0  // ❌ type mismatch
}
```

## Primitive Types

- `i32` – 32‑bit signed integer
- `u32` – 32‑bit unsigned integer
- `f64` – double precision floating point
- `bool` – boolean (`true` or `false`)
- `string` – UTF‑8 text
- `nil` – absence of a value (returned from functions without `return`)

```orus
let flag: bool = true
let text = "hello"       // type inference
```

Numeric types never convert implicitly. Use `as` to cast:

```orus
let a: i32 = -5
let b: u32 = a as u32
```

## Comments

`//` starts a line comment. Block comments use `/* ... */`.

```orus
// single line
let x = 1 /* inline */ + 2

/*
This is a block comment.
*/
```

## Operators

Orus supports common arithmetic (`+`, `-`, `*`, `/`, `%`), comparison (`==`, `!=`, `<`, `>`, `<=`, `>=`), and logical operators (`and`, `or`, `not`). Compound assignments like `+=` and `-=` are also available.

Casting between numeric types must be explicit with `as`.

## Control Flow

### Conditionals

```orus
if n > 0 {
    print("positive")
} elif n == 0 {
    print("zero")
} else {
    print("negative")
}
```

### Loops

```orus
for i in 0..5 {          // 0 to 4
    print(i)
}

while condition {
    // repeat while true
}

break      // exit loop
continue   // next iteration
```

## Arrays

Fixed-length arrays use `[T; N]` syntax. Elements are zero indexed.

```orus
let nums: [i32; 3] = [1, 2, 3]
let first = nums[0]
nums[1] = 20
```

### Dynamic Arrays

Built-in functions can grow arrays dynamically.

```orus
let values: [i32; 1] = [0]
push(values, 10)
print(len(values))  // 2
```

### Slicing

Subarrays are created with `[start..end]` (end exclusive). The start or end may
be omitted to slice from the beginning or to the end of the array.

```orus
let part = nums[0..2]  // first to 3rd element
let part = nums[..2]   // first to 3rd element
let part = nums[0..]   // first to last element
let part = nums[..]    // entire array
```

## Structs

Structs group named fields.

```orus
struct Point {
    x: i32,
    y: i32,
}

let p = Point{ x: 1, y: 2 }
print(p.x)
```

## Methods with `impl`

Methods attach functions to a struct inside an `impl` block. This style is similar to Rust.

```orus
impl Point {
    fn new(x: i32, y: i32) -> Point {
        return Point{ x: x, y: y }
    }

    fn move_by(self, dx: i32, dy: i32) {
        self.x = self.x + dx
        self.y = self.y + dy
    }
}
```

Use the struct name for static methods and a value for instance methods:

```orus
let p = Point.new(1, 2)
p.move_by(3, 4)
```

## Functions

Functions are defined with `fn`. Parameter types are required and the return type follows `->`.

```orus
fn add(a: i32, b: i32) -> i32 {
    return a + b
}

fn greet(name: string) {    // returns nil implicitly
    print("Hello, {}!", name)
}
```

Functions may be declared after their call site. Generic functions must currently appear before use (forward declarations for generics are not yet implemented).

## Pattern Matching

`match` compares a value against patterns, similar to `switch` in other languages but with explicit patterns like Rust.

```orus
match value {
    0 => print("zero"),
    1 => print("one"),
    _ => print("other"),
}
```

The first matching branch runs. Use `_` as a wildcard.

## Error Handling

`try`/`catch` blocks handle runtime errors.

```orus
try {
    let x = 10 / 0
} catch err {
    print("Error: {}", err)
}
```

Error messages include the file, line and column as well as a short stack trace.

## Generics

Functions and structs may take type parameters using angle brackets.

```orus
fn id<T>(x: T) -> T {
    return x
}

struct Box<T> { value: T }
```

Type arguments can often be inferred:

```orus
let a = id<i32>(5)
let b: Box<string> = Box { value: "hi" }
```

Forward declarations for generic functions are not yet supported, so generic functions must appear before they are used. Constraint syntax is planned for a future version.

## Modules

Code can be split into multiple files. Use `use` to import another module at the top level of a file.

```orus
use math::utils
use datetime as dt
```

Modules are only loaded once. Importing the same file twice triggers a runtime error. Each module may define structs, functions and `impl` blocks. Only the entry file needs a `main` function.

## Built-in Functions

Common utilities are always available. See `docs/BUILTINS.md` for full details.

- `print(values...)`
- `len(value)`
- `substring(str, start, len)`
- `push(array, value)` and `pop(array)`
- `range(start, end)`
- `sum(array)`, `min(array)`, `max(array)`
- `type_of(value)`, `is_type(value, name)`
- `input(prompt)`, `int(text)`, `float(text)`
- `sorted(array, reverse)`

```orus
let arr: [i32; 1] = [1]
push(arr, 2)
print(len(arr))
```

## Feature Status

- Modules, pattern matching, error handling and `impl` blocks are **fully implemented**.
- Generics work for basic use cases but lack forward declarations and constraints.
- The standard library is minimal; more built-ins are planned.


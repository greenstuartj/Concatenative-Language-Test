# Concatenative-Language-Test
Exploring implementing a [concatenative language](https://en.wikipedia.org/wiki/Concatenative_programming_language)

An experiment into implementing a concatentive language in [GO](https://golang.org/)

Build:

`$ cd main && go build`

Run:

`$ ./main`

or

`$ ./main {file}`

(switching {file} for the name of a source code file, examples in main/examples)

Uses reverse polish notation

```
>>> 10 3 + println
13
```

Define functions

```
def add1 {
  1 +
}
```

Recursion with arguments and pattern matching

```
def factorial {
  0 -- 1 ; // The -- split arguments from body
           // The ; splits function bodies
  n -- n n 1 - recur * // recur or factorial can be used
                       // recur is tail recursive when in the tail position
}
5 factorial println
```

Anonymous functions (with anonymous recursion)

```
5 { 0 -- 1 ; n -- n n 1 - recur * } println
```

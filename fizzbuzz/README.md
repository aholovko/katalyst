# FizzBuzz

## Instructions

Write a function that takes positive integers and outputs their string representation.
Your function should comply with the following additional rules:

* If the number is a multiple of three, return the string "Fizz".
* If the number is a multiple of five, return the string "Buzz".
* If the number is a multiple of both three and five, return the string "FizzBuzz".

For example, given the numbers from 1 to 15 in order, the function would return:

```text
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
```

## TDD Principles

1. You are not allowed to write any production code unless it is for making a failing unit test pass.

## Implementation Strategies

* Faking It - fake it until you make it.
* Obvious Implementation - if the implementation is obvious, code it.
* Defer decision-making - the best decisions are made with the maximum possible information.
* Rule of Three - until we see three cases of obvious redundancy, we will defer refactoring it out.
* Triangulation - create a specific test case that forces the behavior of the code to change.

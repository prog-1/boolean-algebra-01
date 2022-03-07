# Boolean Algebra

## Truth Table

You are given the following claims:

> If Jones did not meet Smith last night (A), then either Smith was a
> murderer (B), or Jones is telling a lie (C). If Smith was not a murderer,
> then Jones did not meet Smith last night, and the murder happened after midnight (D). If the murder happened after midnight,
> then either Smith was a murderer, or Jones is telling a lie, but
> not both.

`A`, `B`, `C` and `D` are variables describing basic propositions in the claim. Write a Boolean proposition that describes the combination of claims using `A`, `B`, `C` and `D`.

Describe the claim using a function declared as `func proposition(a, b, c, d bool) bool`. Write tests for the function.

Write `main()` that calls the function `proposition` and prints the truth table, e.g.

```
    A      B      C      D        ?
  false  false  false  false      ?
  false  false  false   true      ?
  false  false   true  false      ?
  false  false   true   true      ?
  false   true  false  false      ?
  false   true  false   true      ?
  false   true   true  false      ?
  false   true   true   true      ?
   true  false  false  false      ?
   true  false  false   true      ?
   true  false   true  false      ?
   true  false   true   true      ?
   true   true  false  false      ?
   true   true  false   true      ?
   true   true   true  false      ?
   true   true   true   true      ?
```

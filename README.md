# Advent of Code 2022
These are my AOC solutions for 2022.

To run:

* `go run . all` from each day's directory.
* `go run . all sample` to load a file called "sample.txt" and execute

day0 can be used a template for new days.

Helpful utils, especially for parsing files, in `lib`

## main.go format

Each day has it's own struct, `Today`, that has three functions:

* `Init` - used to parse the puzzle input. The input is the file to read from the day's directory (like `sample.txt` or `input.txt`)
* `Part1` - Called to produce the answer for part 1 (in string format)
* `Part2` - Called to produce the answer for part 2 (in string format)

I generally try to make sure my solutions produce answers for both parts - even though it can often be faster to just edit the solution for part 1 to solve part 2.
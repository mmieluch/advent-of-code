# Day 8: Treetop Tree House

Link: https://adventofcode.com/2022/day/8

Unfortunately, there's a bit of broken logic in the second part of the puzzle, and therefore it expects an incorrect answer.

The second part calls to calculate "scenic scores" for all individual trees, based on the viewing distance from each respective tree. However, there are numerous row and column combinations, where the provided number is 0 - that means there is no tree to be counted into the viewing distance. Mind you - we're supposed to be counting trees, not plots. This affects the final answer, unfortunately. The correct solution for part 2 should be **492800**.

## Solution status

🟢 **WORKING**

## Instructions

Run with example input:

```shell
go run .
```

Run with production input:

```shell
go run . --production
```

Run unit tests:

```shell
go test ./...
```

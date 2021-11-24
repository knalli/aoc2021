# Advent of Code 2021 Solutions
Welcome to my solutions of [Advent Of Code](http://adventofcode.com) 2021 (AOC 2021).

A huge thanks to @topaz and his team for providing this great service.

Just like in 2018, 2019 and 2021, the solutions will be implemented using Go.

## Disclaimer
These are my personal solutions of the Advent Of Code (AOC). The code is
*not indented* to be perfect in any kind of area. This year, my personal
competition was to ~~learn~~ intensify Go handling. These snippets are here for everyone
learning more, too.

If you think, there is a piece of improvement: Go to the code,
fill a PR and we are all happy. Share the knowledge.

## Structure
The AOC contains 25 days with at least one puzzle/question per day (mostly there are two parts).

* Base path is the root folder.
* Each day has sub module named `day01`, `day02` until `day24` with a file `init.go` having 
  a function `Call`.
* The day `tpl` is for templating new days, invoked by the script line `./create_day.sh <day>`.
* Depending on content, a day could import (exported) symbols of a (previous) day.

## Usage

For running the day `day00`
* CLI: just enter `go run main.go 0`

## License / Copyright
Everything is free for all.

Licensed under MIT. Copyright Jan Philipp.
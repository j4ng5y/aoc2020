# AOC 2020

For full visibility and just so my leader board members know I'm doing this myself and not copying answers:

## Running the Go Code

If you want to run my Go code, you should probably have [Go](https://golang.org/dl) installed, but otherwise:

- `clone this repo`
- `cd <aoc2020>/day<N>/go`
- `go run main.go -input <path to the input file> [-part (1|2)]`

OR

you can just head on over the the releases, pick your day, and download the binaries if I happened to have compiled them for you.
You can easily grab my input files for whatever day using the following:

```bash
DAY=1 wget https://raw.githubusercontent.com/j4ng5y/aoc2020/day${DAY}/day${DAY}/input.txt
```

## Running the Rust Code

if you want to run my Rust code, you should probably have [Rust](https://rustup.rs) (with the standard toolchain such as Cargo and whatnot)

## Usage (I'll try to keep these as consistent as possible)

```bash
Usage of day<N>_<OS>_<ARCH>:
  -input string
        The input file to parse
  -part int
        The part of the challenge to display (1/2), all other integers (or no value), displays both
```

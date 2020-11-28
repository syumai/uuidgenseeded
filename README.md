# uuidgenseeded

* A CLI tool to generate UUID with seed text.

## Installation

```console
go get github.com/syumai/uuidgenseeded
```

## Usage

```sh
# Using with seed argument
$ uuidgenseeded example
475A87DF-68E7-4C04-A4F0-E11504AEC553

# Using with pipe
$ echo -n "example" | uuidgenseeded
475A87DF-68E7-4C04-A4F0-E11504AEC553

# Output in lower case (shorthand: -l)
$ uuidgenseeded -lower example
475a87df-68e7-4c04-a4f0-e11504aec553
```

## LICENSE

MIT

## Author

syumai

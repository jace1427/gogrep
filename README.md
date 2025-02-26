# GoGrep

## Usage

Usage of gogrep:

- c Print in color (default true)
- i Ignore case distinctions
- n Print line number with output lines
- v Select non-matching lines

## Installation

Download gogrep:

`git clone https://github.com/jace1427/gogrep.git`

Build:

`go build ./cmd/gogrep/`

Then run:

`$ ./gogrep -n Romeo shakespeare.txt`

```txt
shakespeare.txt:95854:  Romeo, son to Montague.
shakespeare.txt:95856:  Mercutio, kinsman to the Prince and friend to Romeo.
shakespeare.txt:95857:  Benvolio, nephew to Montague, and friend to Romeo
shakespeare.txt:95861:  Balthasar, servant to Romeo.
shakespeare.txt:96052:  M. Wife. O, where is Romeo? Saw you him to-day?
shakespeare.txt:96093:                       Enter Romeo.
shakespeare.txt:96105:  Ben. It was. What sadness lengthens Romeo's hours?
shakespeare.txt:96143:    This is not Romeo, he's some other where.
shakespeare.txt:96241:                   Enter Benvolio and Romeo.
shakespeare.txt:96252:  Ben. Why, Romeo, art thou mad?
...
...
```

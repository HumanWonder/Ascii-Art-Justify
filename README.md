# Ascii-Art-Justify

## About 

Same instructions as in the first subject but here the alignment can be changed.

To change the alignment of the output it must be possible to use a flag --align=<type>, in which type can be :

- center

- left

- right

- justify

The representation must be adapted to the terminal size. If the terminal window is redeuced the graphical representation should be adapted to the terminal size accordingly.

Only text that fits the terminal size will be tested.

## Getting Started

Clone this repository : https://zone01normandie.org/git/afouquem/ascii-art-justify.git

### Installing

The format to use the program is the one below :

```
go run . --align=<option> [string] [font]
```

-The flag must have exactly the same format as above, any other formats must return the following usage message:

```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard

```

## Usage <a name = "usage"></a>

The program still runs with only one STRING argument, or with a BANNER specified.

Example:

```
go run main.go --align=right "Hello There" shadow
```
It can also run with a color specified (if no letters are specified, the whole string is colored):

```
go run . --color=<color> <letters to be colored> "something"
```

Or a basic text to Ascii program :
```
go run . something [optional font]
```


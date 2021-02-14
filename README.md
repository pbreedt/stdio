# stdio
> A simple library to read input from standard IO (os.Stdin)

## General info
This project was born mainly to explore some Go features, such as:
- Go modules (with nothing in the main working folder)
- Reading user input from command line
- Unit testing command line input

## Technologies
* Go - version 1.15

## Setup
Like with all Go modules, you can simply "go get" it.

```go get github.com/pbreedt/stdio/input```

## Code Examples
Usage is simple and clear:

```Go
package main

import (
	"fmt"

	"github.com/pbreedt/stdio/input"
)

func main() {
    // Read user's name from os.Stdin
	userName, err := input.ReadString("What is your name? ")
	if err != nil {
		fmt.Printf("Encountered error %v\n", err)
		return
	}

    // Read user's age as int from os.Stdin
	userAge, err := input.ReadInt("What's your age? ")
	if err != nil {
		fmt.Printf("Encountered error %v\n", err)
		return
	}
	fmt.Printf("Hello %s. You are %d years old.\n", userName, userAge)

    /* Read user's confirmation from os.Stdin
	   2nd param 'true' will interpret Yes/No variants
	   as well as the standard Go strconv.ParseBool(str) values
    */
	confirm, err := input.ReadBool("Is this correct? ", true)
	if err != nil {
		fmt.Printf("Encountered error %v\n", err)
		return
	}
	if confirm {
		fmt.Printf("Great! Pleased to meet you, %s.\n", userName)
	} else {
		fmt.Printf("Sorry I got that wrong. :(\n")
	}
}
```

See [Random Number Guessing game](http://github.com/pbreedt/rndnumguess) for a silly implementation 🙂

## Features
Current features include:
* Reading an integer (int) value from cmd line
* Reading an float (float64) value from cmd line
* Reading an boolean (bool) value from cmd line.  Could also accept variants of Yes/No
* Reading an string (string) value from cmd line

Possible improvements:
* Provide io.Reader to read from / io.Writer to write to

## Status
Project is: _in progress_

## Credits
Thanks to [ritaly](https://github.com/ritaly/README-cheatsheet) for a quick readme template

## Contact
Created by [@pbreedt](mailto:petrus.breedt@gmail.com)
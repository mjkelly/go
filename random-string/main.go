// random-string generates a configurable random string, for passwords.
package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Flags.
var (
	passwordLength = flag.Int("length", 16, "Length of password to generate.")
	LowerAlphaNum  = flag.Bool("loweralphanum", false,
		"Only generate passwords with lowercase alphanumeric characters. "+
			"Mutually exclusive with --alphanum and --numeric.")
	AlphaNum = flag.Bool("alphanum", false,
		"Only generate passwords with alphanumeric characters. "+
			"Mutually exclusive with --loweralphanum and --numeric.")
	Numeric = flag.Bool("numeric", false,
		"Only generate passwords with numeric characters. "+
			"Mutually exclusive with --alphanum and --loweralphanum.")
	Quiet = flag.Bool("quiet", false, "Suppress unnecessary output.")
)

// Possible characters to be used in passwords. (These are all one byte long.
// Larger characters might cause problems.)
var (
	lowerAlphaChars = "abcdefghijklmnopqrstuvwxyz"
	upperAlphaChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericChars    = "0123456789"
	symbolChars     = "~!@#%^&*()-_=+[]{}|;:<>,./?"
)

func main() {
	flag.Parse()
	var chars string
	charsetFlags := 0

	if *LowerAlphaNum {
		chars = lowerAlphaChars + numericChars
		charsetFlags++
	}
	if *AlphaNum {
		chars = lowerAlphaChars + upperAlphaChars + numericChars
		charsetFlags++
	}
	if *Numeric {
		chars = numericChars
		charsetFlags++
	}
	if charsetFlags == 0 {
		chars = lowerAlphaChars + upperAlphaChars + numericChars + symbolChars
	} else if charsetFlags > 1 {
		log.Fatal("You can only specify one of --loweralphanum, --alphanum, and --numeric.")
	}

	numChars := len(chars)
	numCharsBig := big.NewInt(int64(numChars)) // used for crypto/rand
	bitsPerChar := math.Log2(float64(numChars))
	totalBits := bitsPerChar * float64(*passwordLength)

	if !*Quiet {
		fmt.Printf("%d characters long.\n", *passwordLength)
		fmt.Printf("Choosing from %d characters. %2.3f bits of entropy per character.\n",
			numChars, bitsPerChar)
		fmt.Printf("%2.3f total bits of entropy.\n\n", totalBits)
		fmt.Println("Password:")
	}

	var passBuffer bytes.Buffer
	for i := 0; i < *passwordLength; i++ {
		randBig, err := rand.Int(rand.Reader, numCharsBig)
		if err != nil {
			log.Fatalf("%s", err)
		}
		randChar := chars[randBig.Int64()]
		passBuffer.WriteByte(randChar)
	}

	fmt.Println(passBuffer.String())
}

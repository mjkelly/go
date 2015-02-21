// random-phrase generates a configurable random passphrase.
package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"regexp"
	"strings"
)

// Flags.
var (
	wordsPerPhrase = flag.Int("words", 5, "Number of words per phrase.")
	numPhrases     = flag.Int("phrases", 10, "Number of phrases to show.")
	dictionary     = flag.String("dictionary", "/usr/share/dict/words", "Dictionary to use.")
	simple         = flag.Bool("simple", false, "Simple mode: Use only lowercase words, without aprostrophes.")
	quiet          = flag.Bool("quiet", false, "Suppress unnecessary output.")
)

// Globals.
var (
	// wordRegexp restricts which words we use from the dictionary. Restrictions
	// are: (1) at least 3 chars long, (2) only contains A-Z (any case) or
	// apostrophe.
	//
	// The minimum word length restriction is to keep the total passphrase
	// length high, so it's much harder to brute-force the passphrase as a
	// random string than as a passphrase. (That ensures that our entropy
	// calculations, which are based on the string as a passphrase, stay
	// relevant.)
	//
	// TODO(mjkelly): Justify this with some math.
	wordRegexp = regexp.MustCompile(`^[a-zA-Z']{3,}$`)
	// simpleWordRegexp is like wordRegexp, but we disallow uppercase letters and
	// apostrophes.
	simpleWordRegexp = regexp.MustCompile(`^[a-z]{3,}$`)
)

// bitsToLength returns how long a random sequence of typeable characters would
// have to be in order to have 'bits' bits of entropy. We assume a character
// set of 89 characters, which is what our friend,
// github.com/mjkelly/go/random-string, uses.
func bitsToLength(bits float64) float64 {
	charset := float64(89)
	return bits / math.Log2(charset)
}

func main() {
	flag.Parse()

	file, err := os.Open(*dictionary)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalWords := 0
	words := make([]string, 0)
	r := wordRegexp
	if *simple {
		r = simpleWordRegexp
	}
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			words = append(words, scanner.Text())
		}
		totalWords++
	}

	numWords := len(words)
	numWordsBig := big.NewInt(int64(numWords))
	bitsPerWord := math.Log2(float64(numWords))
	bitsPerPhrase := bitsPerWord * float64(*wordsPerPhrase)
	totalBits := bitsPerPhrase - math.Log2(float64(*numPhrases))
	equivalentRandomLength := bitsToLength(bitsPerPhrase)

	if !*quiet {
		fmt.Printf("%d possible words (of %d in %s).\n", numWords, totalWords, *dictionary)
		fmt.Printf("%d random words per phrase.\n", *wordsPerPhrase)
		fmt.Printf("∴ %f bits of entropy per word.\n", bitsPerWord)
		fmt.Printf("∴ %f bits of entropy per phrase.\n", bitsPerPhrase)
		fmt.Printf("(approximately equivalent to %.0f char random password)\n", equivalentRandomLength)
		fmt.Printf("%d phrases to choose from.\n", *numPhrases)
		fmt.Printf("∴ %f bits if you pick one phrase from this list.\n", totalBits)
		fmt.Println("---------------------------------------------------")
	}

	for i := 0; i < *numPhrases; i++ {
		phrase := make([]string, 0, *numPhrases)
		for j := 0; j < *wordsPerPhrase; j++ {
			randBig, err := rand.Int(rand.Reader, numWordsBig)
			if err != nil {
				log.Fatal(err)
			}
			phrase = append(phrase, words[randBig.Int64()])
		}
		fmt.Println(strings.Join(phrase, " "))
	}
}

package main

import (
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/gin-gonic/gin"
)

func myloCorpus() string {
	bs, err := ioutil.ReadFile("corpus.txt")
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func loremIpsumCorpus() string {
	bs, err := ioutil.ReadFile("lorem.txt")
	if err != nil {
		panic(err)
	}
	return string(bs)
}


// This is a fixed, sized array
type FixedArrayOfInt [10]uint
// This is a pointer to the same
type PtrFxdRryFNt *[10]uint

func words(s string) []string {
	return strings.Fields(s)
}

// I am a function
// Called getRandomWords
// I take `wrds`, a slice of string
// and return `out`, a slice of string
func getRandomWords(wrds []string) []string {
	var out []string
	numberOfWords := rand.Intn(len(wrds))
	for i := 0; i < numberOfWords; i++ {
		rndWrd := wrds[rand.Intn(len(wrds))]
		out = append(out, rndWrd)
	}
	return out
}

func generateLoremIpsumText() string {
	// returns some words which are random
	// starting with "lorem ipsum"
	myloWords := words(myloCorpus())
	loremWords := words(loremIpsumCorpus())
	sliceOfMyloWords := getRandomWords(myloWords)
	sliceOfLoremWords := getRandomWords(loremWords)
	words := append(sliceOfLoremWords, sliceOfMyloWords...)
	for i := len(words) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	// return lorem ipsum text with some weekly words
	// in a web request/endpoint
	r := gin.Default()
	loremHandler := func(c *gin.Context) {
		c.String(200, generateLoremIpsumText())
	}
	r.GET("/lorem",loremHandler)

	r.Run()
}

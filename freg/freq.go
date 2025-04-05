package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// What is the most common word (ignoring case) in sherlock.txt?
// word frequency

func main() {
	file, err := os.Open("freg/sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	w, err := mostCommon(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(w)
	// mapDemo()

	/*
		// path := "C:\TO\NEW\REPORT.CSV"
		// `s` is "raw" string, \ is just a \
		path := `C:\to\new\report\csv`
		fmt.Println(path)

	*/
}

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(freqs)
}

var request = `GET / HTTP/1.1
HOST: httpbin.or
Connection: Close
`

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func mapDemo() {
	var stocks map[string]int // word -> count
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("%s -> $%d\n", sym, price)

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%d\n", sym, price)

	} else {
		fmt.Printf("%s not found \n", sym)
	}
	/*
		stocks = make(map[string]int)
		stocks[sym] = 136

	*/

	stocks = map[string]int{
		sym:    137,
		"AAPL": 173,
	}
	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%d\n", sym, price)

	} else {
		fmt.Printf("%s not found \n", sym)
	}

	for k := range stocks {
		fmt.Println(k)
	}

	for k, v := range stocks {
		fmt.Println(k, v)
	}

	for _, v := range stocks {
		fmt.Println(v)
	}

	delete(stocks, "AAPL")
	fmt.Println(stocks)
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""

	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}
	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int)
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}

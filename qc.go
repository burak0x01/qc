package main

import (
	"bufio"
	"flag"
	"net/url"
	"os"
)

func main() {

	changedValue := flag.String("p", "NEW_VALUE", "new query value")
	inputFile := flag.String("i", "", "domain list")
	outputFile := flag.String("o", "result.txt", "output file")
	flag.Parse()

	file, err := os.Open(*inputFile)
	check(err)
	defer file.Close()

	wfile, err := os.OpenFile(*outputFile, os.O_CREATE|os.O_WRONLY, 0666)
	check(err)
	defer wfile.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		u, err := url.Parse(scanner.Text())
		check(err)

		originQuery, _ := url.ParseQuery(u.RawQuery)

		tempQuery := make(map[string][]string)
		for k, v := range originQuery {
			tempQuery[k] = v
		}

		for k, _ := range originQuery {

			for k1, v := range tempQuery {
				originQuery[k1] = v
			}

			originQuery.Set(k, *changedValue)

			u.RawQuery = originQuery.Encode()

			newUrl := u.Scheme + "://" + u.Host + u.Path + "?" + u.RawQuery + "\n"
			wfile.WriteString(newUrl)
			check(err)
		}
	}
	check(scanner.Err())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

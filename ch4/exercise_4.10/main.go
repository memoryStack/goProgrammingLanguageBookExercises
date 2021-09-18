// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goProgrammingLanguageBook/ch4/github"
)

func getCategoryKey(days int) string {
	switch {
	case days < 30:
		return "less than a month"
	case days < 365:
		return "less than a year"
	case days > 365:
		return "more than a year"
	}
	return "invalid category"
}

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d total issues:\n", result.TotalCount)

	categoryMap := make(map[string][]*github.Issue)

	for _, item := range result.Items {
		days := item.CreatedAt.Day()
		categoryKey := getCategoryKey(days)

		categoryMap[categoryKey] = append(categoryMap[categoryKey], item)
	}

	for key, value := range categoryMap {
		fmt.Printf("issues for %s:\n", key)
		for _, item := range value {
			fmt.Printf("#%-5d %9.9s %.55s %d\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt.Day())
			// fmt.Printf("%v\n", issue)
		}
	}

}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"errors"
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"linear algebra":        {"calculus"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {

	if coursesOrder, err := topoSort(prereqs); err != nil {
		fmt.Println(err)
	} else {
		for i, course := range coursesOrder {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	var cycle bool

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}

			var found bool
			for _, course := range order {
				if course == item {
					found = true
					break
				}
			}
			if !found {
				cycle = true
				break
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// this is where all the topo sort is going on
	visitAll(keys)

	if cycle {
		return nil, errors.New("cycle found, invalid dependencies, aborting")
	}

	return order, nil
}

//!-main

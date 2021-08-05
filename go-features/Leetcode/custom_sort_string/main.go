package main

import "fmt"

/*
order and str are strings composed of lowercase letters. In order, no letter occurs more than once.
order was sorted in some custom order previously. We want to permute the characters of str so that
they match the order that order was sorted. More specifically, if x occurs before y in order, then x
should occur before y in the returned string. Return any permutation of str (as a string) that satisfies
this property.
 */

func main() {
	fmt.Println(customSortString("cbafgzzzz", "bc"))
}

// loop str, find any letters that match order
// if match, pop from order, pop from str
// if end of loop || order is empty, append copy of order to the front

func customSortString(o string, s string) string {


	return s
}

func moveToFront(str string, i int) string {
	str = string(str[i]) + str[:i] + str[i + 1:]
	return str
}

func pop(str string, index int) string {
	// pops letter at index
	str = str[:index] + str[index + 1:]
	return str
}

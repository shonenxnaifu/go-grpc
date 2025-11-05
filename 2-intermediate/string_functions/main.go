package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello Go!"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[1:7])

	// String Convertion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	// string splitting
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ",")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts)
	fmt.Println(parts1)

	countries := []string{"germany", "france", "italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go"))

	replaced := strings.Replace(str, "Go", "Universe", 1)
	fmt.Println(replaced)

	strwspace := " Hello Everyone!"
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	fmt.Println(strings.Repeat("foo ", 3))

	fmt.Println(strings.Count("Hello", "l"))

	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))
	fmt.Println(strings.HasSuffix("Hello", "la"))

	str5 := "Hel1lo, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	str6 := "Hello, 侍恩"
	fmt.Println(utf8.RuneCountInString(str6))

	// String Builder
	// for memory efficiency
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	//Convert builder to a string
	result1 := builder.String()
	fmt.Println(result1)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result1 = builder.String()
	fmt.Println(result1)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result1 = builder.String()
	fmt.Println(result1)
}

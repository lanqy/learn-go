package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("A", "B"))             // A < B
	fmt.Println(strings.Compare("B", "A"))             // B > A
	fmt.Println(strings.Compare("Japan", "Australia")) // J > A
	fmt.Println(strings.Compare("Australia", "Japan")) // A > J
	fmt.Println(strings.Compare("Germany", "Germany")) // G == G
	fmt.Println(strings.Compare("Germany", "GERMANY")) // GERMANY > Germany
	fmt.Println(strings.Compare("", ""))
	fmt.Println(strings.Compare("", " ")) // Space is less
}

func main1() {
	fmt.Println(strings.Contains("Australia", "Aus"))
	fmt.Println(strings.Contains("Australia", "Australian"))
	fmt.Println(strings.Contains("Japan", "JAPAN")) // Case sensitive
	fmt.Println(strings.Contains("Japan", "JAP"))   // Case sensitive
	fmt.Println(strings.Contains("Become inspired to travel to Australia.", "Australia"))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("12554", "1"))
}

func main2() {
	fmt.Println(strings.ContainsAny("Australia", "a"))
	fmt.Println(strings.ContainsAny("Australia", "r & a"))
	fmt.Println(strings.ContainsAny("JAPAN", "j"))
	fmt.Println(strings.ContainsAny("JAPAN", "J"))
	fmt.Println(strings.ContainsAny("JAPAN", "JAPAN"))
	fmt.Println(strings.ContainsAny("JAPAN", "japan"))
	fmt.Println(strings.ContainsAny("Shell-12541", "1"))

	// Contains vs ContainsAny
	fmt.Println(strings.ContainsAny("Shell-12541", "1-2")) // true
	fmt.Println(strings.Contains("Shell-12541", "1-2"))    // false
}

func main3() {
	fmt.Println(strings.Count("Australia", "a"))
	fmt.Println(strings.Count("Australia", "A"))
	fmt.Println(strings.Count("Australia", "M"))
	fmt.Println(strings.Count("Japanese", "Japan"))
	fmt.Println(strings.Count("Japan", "Japanese"))
	fmt.Println(strings.Count("Shell-25152", "-25"))
	fmt.Println(strings.Count("Shell-25152", "-21"))
	fmt.Println(strings.Count("test", "")) // length of string + 1
	fmt.Println(strings.Count("test", " "))
}

func main4() {
	fmt.Println(strings.EqualFold("Australia", "AUSTRALIA"))
	fmt.Println(strings.EqualFold("Australia", "aUSTRALIA"))
	fmt.Println(strings.EqualFold("Australia", "Australia"))
	fmt.Println(strings.EqualFold("Australia", "Aus"))
	fmt.Println(strings.EqualFold("Australia", "Australia & Japan"))
	fmt.Println(strings.EqualFold("JAPAN-1254", "japan-1254"))
	fmt.Println(strings.EqualFold(" ", " ")) // single space both side
	fmt.Println(strings.EqualFold(" ", " ")) // double space right side
}

func main5() {
	testString := "Australia is a country and continent surrounded by the Indian and Pacific oceans."
	testArray := strings.Fields(testString)
	for _, v := range testArray {
		fmt.Println(v)
	}
}

func main6() {
	x := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	strArray := strings.FieldsFunc(`Australia major cities – Sydney, Brisbane,
                                 Melbourne, Perth, Adelaide – are coastal`, x)
	for _, v := range strArray {
		fmt.Println(v)
	}

	fmt.Println("\n*****************Split by number*******************\n")

	y := func(c rune) bool {
		return unicode.IsNumber(c)
	}
	testArray := strings.FieldsFunc(`1 Sydney Opera House.2 Great Barrier Reef.3 Uluru-Kata Tjuta National Park.4 Sydney Harbour Bridge.5 Blue Mountains National Park.6 Melbourne.7 Bondi Beach`, y)
	for _, w := range testArray {
		fmt.Println(w)
	}
}

func main7() {
	fmt.Println(strings.HasPrefix("Australia", "Aus"))
	fmt.Println(strings.HasPrefix("Australia", "aus"))
	fmt.Println(strings.HasPrefix("Australia", "Jap"))
	fmt.Println(strings.HasPrefix("Australia", ""))
}

func main8() {
	fmt.Println(strings.HasSuffix("Australia", "lia"))
	fmt.Println(strings.HasSuffix("Australia", "A"))
	fmt.Println(strings.HasSuffix("Australia", "LIA"))
	fmt.Println(strings.HasSuffix("123456", "456"))
	fmt.Println(strings.HasSuffix("Australia", ""))
}

func main9() {
	fmt.Println(strings.Index("Australia", "Aus"))
	fmt.Println(strings.Index("Australia", "aus"))
	fmt.Println(strings.Index("Australia", "A"))
	fmt.Println(strings.Index("Australia", "a"))
	fmt.Println(strings.Index("Australia", "Jap"))
	fmt.Println(strings.Index("Japan-124", "-"))
	fmt.Println(strings.Index("Japan-124", ""))
}

func main10() {
	fmt.Println(strings.IndexAny("australia", "japan")) // a position
	fmt.Println(strings.IndexAny("japan", "pen"))       // p position
	fmt.Println(strings.IndexAny("mobile", "one"))      // o position
	fmt.Println(strings.IndexAny("123456789", "4"))     // 4 position
	fmt.Println(strings.IndexAny("123456789", "0"))     // 0 position
}

func main11() {
	var s, t, u byte
	t = '1'
	fmt.Println(strings.IndexByte("australia", t))
	fmt.Println(strings.IndexByte("LONDON", t))
	fmt.Println(strings.IndexByte("JAPAN", t))

	s = 1
	fmt.Println(string.IndexByte("5221-JAPAN", s))

	u = '1'
	fmt.Println(strings.IndexByte("5221-JAPAN", u))
}

func main12() {
	var s, t, u rune
	t = '1'
	fmt.Println(strings.IndexRune("australia", t))
	fmt.Println(strings.IndexRune("LONDON", t))
	fmt.Println(strings.IndexRune("JAPAN", t))

	s = 1
	fmt.Println(strings.IndexRune("5221-JAPAN", s))

	u = '1'
	fmt.Println(strings.IndexRune("5221-JAPAN", u))
}

func main13() {
	textString := []string{"Australia", "Japan", "Canada"}
	fmt.Println(strings.Join(textString, "-"))

	// Slice of strings
	textNum := []strings{"1", "2", "3", "4", "5"}
	fmt.Println(strings.Join(textNum, ""))
}

package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("contains: ", s.Contains("test", "es"))
	p("count: ", s.Count("test", "t"))
	p("hasPrefix: ", s.HasPrefix("test", "te"))
	p("hasSuffix: ", s.HasSuffix("test", "st"))
	p("index: ", s.Index("test", "e"))
	p("join: ", s.Join([]string{"a", "b"}, "-"))
	p("repeat: ", s.Repeat("a", 5))
	p("replace: ", s.Replace("foo", "o", "0", -1))
	p("replace: ", s.Replace("foo", "o", "0", 1))
	p("split: ", s.Split("a-b-c-d-e", "-"))
	p("toLower: ", s.ToLower("TEST"))
	p("toUpper", s.ToUpper("test"))
}

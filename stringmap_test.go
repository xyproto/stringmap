package stringmap

import (
	"testing"
	"fmt"
	"strings"
)

func TestDiv(t *testing.T) {
	fmt.Println("hi")
	m := New("cheese:bread", "coffee:tea")
	fmt.Println(*m)
	s := New("cheese:bread", "coffee:tea").Slice()
	fmt.Println(s)
	x := New(s...)
	fmt.Println(*x)
	fmt.Println(x.String())
	t2 := New("cheese:bread, coffee:tea")
	fmt.Println(t2.String())
	fmt.Println(t2.Slice())
	z := New("cheese:bread,  coffee:tea, apple:banana")
	fmt.Println(*z)
	fmt.Println(z.String())
	fmt.Println(z.Slice())
	z.Map(strings.ToUpper)
	fmt.Println(*z)
	z.Map(strings.ToLower)
	z.MapKeys(strings.ToUpper)
	fmt.Println(*z)
	fmt.Println(z.SortedKeys())
	fmt.Println("KSBV:", z.KeysSortedByValues())
	fmt.Println(z.GetMap())
	//t.Errorf("Error!\n")
}


package editlib_test

import (
	"bytes"
	"fmt"
	. "github.com/andresvia/editlib/editlib"
	"strings"
)

func ExampleAppend() {
	start := "# editlib start do not edit manually"
	end := "# editlib end do not edit manually"
	original := `this is my configuration file
and is full of things`
	insert := `this is my
brand new conf`
	outbuf := bytes.Buffer{}
	Edit(&outbuf, strings.NewReader(original), start, end, insert)
	out := outbuf.String()
	fmt.Print(out)
	// Output:
	// this is my configuration file
	// and is full of things
	// # editlib start do not edit manually
	// this is my
	// brand new conf
	// # editlib end do not edit manually
}

func ExampleEdit() {
	start := "# editlib start do not edit manually"
	end := "# editlib end do not edit manually"
	original := `this is my configuration file
# editlib start do not edit manually
and it have some configuration
# editlib end do not edit manually
and is full of things`
	insert := "but i want a new one"
	outbuf := bytes.Buffer{}
	Edit(&outbuf, strings.NewReader(original), start, end, insert)
	out := outbuf.String()
	fmt.Print(out)
	// Output:
	// this is my configuration file
	// # editlib start do not edit manually
	// but i want a new one
	// # editlib end do not edit manually
	// and is full of things
}

func ExampleRemove() {
	start := "# editlib start do not edit manually"
	end := "# editlib end do not edit manually"
	original := `this is my configuration file
and is full of things`
	insert := ""
	outbuf := bytes.Buffer{}
	Edit(&outbuf, strings.NewReader(original), start, end, insert)
	out := outbuf.String()
	fmt.Print(out)
	// Output:
	// this is my configuration file
	// and is full of things
}

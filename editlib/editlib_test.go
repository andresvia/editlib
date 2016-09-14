package editlib_test

import (
	"bytes"
	. "github.com/andresvia/editlib/editlib"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"strings"
	"testing"
)

func TestEdit(t *testing.T) {

	Convey("Given start, termination and some text to insert", t, func() {

		start_text := "# EDITLIB START"
		termination_text := "# EDITLIB TERMINATE"
		to_insert := "hello world"

		Convey("The final result should be the same", func() {

			final_text := "I want to be edited.\n# EDITLIB START\nhello world\n# EDITLIB TERMINATE"

			Convey("With an initial text", func() {

				buf := bytes.Buffer{}
				initial_text := "I want to be edited."

				Convey("The text is appended", func() {
					Edit(&buf, strings.NewReader(initial_text), start_text, termination_text, to_insert)
					buffer_bytes, _ := ioutil.ReadAll(&buf)
					So(string(buffer_bytes), ShouldEqual, final_text)
				})

			})

			Convey("Or with an existing text", func() {

				buf := bytes.Buffer{}
				existing_text := "I want to be edited.\n# EDITLIB START\nreplace me\n# EDITLIB TERMINATE"

				Convey("The text is edited in place", func() {
					Edit(&buf, strings.NewReader(existing_text), start_text, termination_text, to_insert)
					buffer_bytes, _ := ioutil.ReadAll(&buf)
					So(string(buffer_bytes), ShouldEqual, final_text)
				})

			})

			Convey("With an already edited text", func() {

				buf := bytes.Buffer{}

				Convey("Text is not changed", func() {
					Edit(&buf, strings.NewReader(final_text), start_text, termination_text, to_insert)
					buffer_bytes, _ := ioutil.ReadAll(&buf)
					So(string(buffer_bytes), ShouldEqual, final_text)
				})

			})
		})

	})

	Convey("Given start and termination texts", t, func() {

		start_text := "# EDITLIB START"
		termination_text := "# EDITLIB TERMINATE"

		Convey("The final result should be the same", func() {

			final_text := "I want to be edited."

			Convey("With an existing text", func() {

				buf := bytes.Buffer{}
				existing_text := "I want to be edited.\n# EDITLIB START\ndelete me\n# EDITLIB TERMINATE"

				Convey("Text is removed", func() {
					Edit(&buf, strings.NewReader(existing_text), start_text, termination_text, "")
					buffer_bytes, _ := ioutil.ReadAll(&buf)
					So(string(buffer_bytes), ShouldEqual, final_text)
				})

			})

			Convey("Or with an initial text", func() {

				buf := bytes.Buffer{}
				initial_text := "I want to be edited."

				Convey("Text is not changed", func() {
					Edit(&buf, strings.NewReader(initial_text), start_text, termination_text, "")
					buffer_bytes, _ := ioutil.ReadAll(&buf)
					So(string(buffer_bytes), ShouldEqual, final_text)
				})

			})
		})

	})

}

func TestSimpleEdit(t *testing.T) {
	Convey("Testing scenarios", t, func() {

		Convey("appending empty 1", func() {
			edit := ""
			expected := `#start
insert
#end`
			result := SimpleEdit(edit, "#start", "#end", "insert")
			So(result, ShouldEqual, expected)
		})

		Convey("appending not empty 1", func() {
			edit := `
`
			expected := `

#start
insert
#end`
			result := SimpleEdit(edit, "#start", "#end", "insert")
			So(result, ShouldEqual, expected)
		})

		Convey("appending not empty 2", func() {
			edit := "foo"
			expected := `foo
#start
insert
#end`
			result := SimpleEdit(edit, "#start", "#end", "insert")
			So(result, ShouldEqual, expected)
		})

		Convey("appending not empty 3", func() {
			edit := `
foo`
			expected := `
foo
#start
insert
#end`
			result := SimpleEdit(edit, "#start", "#end", "insert")
			So(result, ShouldEqual, expected)
		})

		Convey("appending not empty 4", func() {
			edit := `foo
`
			expected := `foo

#start
insert
#end`
			result := SimpleEdit(edit, "#start", "#end", "insert")
			So(result, ShouldEqual, expected)
		})

		Convey("edit not covered", func() {
			edit := `#start
edit me
#end`
			expected := `#start
edited
#end`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)

		})
		Convey("edit up covered 1", func() {
			edit := `
#start
edit me
#end`
			expected := `
#start
edited
#end`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)

		})
		Convey("edit up covered 2", func() {
			edit := `cover
#start
edit me
#end`
			expected := `cover
#start
edited
#end`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)

		})
		Convey("edit up covered 3", func() {
			edit := `cover

#start
edit me
#end`
			expected := `cover

#start
edited
#end`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)

		})
		Convey("edit down covered 1", func() {
			edit := `#start
edit me
#end
`
			expected := `#start
edited
#end
`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)
		})
		Convey("edit down covered 2", func() {
			edit := `#start
edit me
#end
cover`
			expected := `#start
edited
#end
cover`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)
		})
		Convey("edit down covered 3", func() {
			edit := `#start
edit me
#end

cover`
			expected := `#start
edited
#end

cover`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)
		})
		Convey("edit sourounded 1", func() {
			edit := `
#start
edit me
#end
`
			expected := `
#start
edited
#end
`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)
		})
		Convey("edit sourounded 2", func() {
			edit := `cover
#start
edit me
#end
cover`
			expected := `cover
#start
edited
#end
cover`
			result := SimpleEdit(edit, "#start", "#end", "edited")
			So(result, ShouldEqual, expected)
		})

		Convey("removing not covered", func() {
			edit := `#start
edit me
#end`
			expected := ""
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})
		Convey("removing up covered", func() {
			edit := `cover
#start
edit me
#end`
			expected := "cover"
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})
		Convey("removing down covered", func() {
			edit := `#start
edit me
#end
cover`
			expected := "cover"
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})

		Convey("removing sourounded 1", func() {
			edit := `cover
#start
edit me
#end
cover`
			expected := `cover
cover`
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})

		Convey("removing sourounded 2", func() {
			edit := `
cover

#start
edit me
#end

cover
`
			expected := `
cover


cover
`
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})

		Convey("already edited not covered", func() {
			edit := `#start
don't edit me
#end`
			expected := `#start
don't edit me
#end`
			result := SimpleEdit(edit, "#start", "#end", "don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("already edited up covered", func() {
			edit := `cover
#start
don't edit me
#end`
			expected := `cover
#start
don't edit me
#end`
			result := SimpleEdit(edit, "#start", "#end", "don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("already edited down covered", func() {
			edit := `#start
don't edit me
#end
cover`
			expected := `#start
don't edit me
#end
cover`
			result := SimpleEdit(edit, "#start", "#end", "don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("already edited sourounded 1", func() {
			edit := `cover
#start
don't edit me
#end
cover`
			expected := `cover
#start
don't edit me
#end
cover`
			result := SimpleEdit(edit, "#start", "#end", "don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("already edited sourounded 2", func() {
			edit := `
#start
don't edit me
#end
`
			expected := `
#start
don't edit me
#end
`
			result := SimpleEdit(edit, "#start", "#end", "don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("start substring 1", func() {
			edit := `cover
foo#start
don't edit me
#end
cover`
			expected := `cover
foo#start
don't edit me
#end
cover`
			result := SimpleEdit(edit, "#start", "#end", "i said don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("start substring 2", func() {
			edit := `foo#start
don't edit me
#end`
			expected := `foo#start
don't edit me
#end`
			result := SimpleEdit(edit, "#start", "#end", "i said don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("end substring 1", func() {
			edit := `#start
don't edit me
#ending`
			expected := `#start
don't edit me
#ending`
			result := SimpleEdit(edit, "#start", "#end", "i said don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("end substring 2", func() {
			edit := `#start
don't edit me
#ending
and more`
			expected := `#start
don't edit me
#ending
and more`
			result := SimpleEdit(edit, "#start", "#end", "i said don't edit me")
			So(result, ShouldEqual, expected)
		})
		Convey("substrings", func() {
			edit := `
#starting
don't edit me
a#ending
and more`
			expected := `
#starting
don't edit me
a#ending
and more
#start
i said don't edit me
#end`
			result := SimpleEdit(edit, "#start", "#end", "i said don't edit me")
			So(result, ShouldEqual, expected)
		})

		Convey("nothing on empty", func() {
			edit := ""
			expected := ""
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})
		Convey("nothing on not empty 1", func() {
			edit := "do nothing"
			expected := "do nothing"
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})
		Convey("nothing on not empty 2", func() {
			edit := `
do nothing`
			expected := `
do nothing`
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})
		Convey("nothing on not empty 3", func() {
			edit := `do nothing
`
			expected := `do nothing
`
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})
		Convey("nothing on not empty 4", func() {
			edit := `
do nothing
`
			expected := `
do nothing
`
			result := SimpleEdit(edit, "#start", "#end", "")
			So(result, ShouldEqual, expected)
		})
	})
}

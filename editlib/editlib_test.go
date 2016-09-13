package editlib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEdit(t *testing.T) {

	Convey("Given start, termination and some text to insert", t, func() {

		start_text := "# EDITLIB START"
		termination_text := "# EDITLIB TERMINATE"
		to_insert := "hello world"

		Convey("The final result should be the same", func() {
			final_text := `I want to be edited.
# EDITLIB START
hello world
# EDITLIB TERMINATE`

			Convey("With an initial text", func() {

				initial_text := "I want to be edited."

				Convey("The text is appended", func() {
					edited_text, edited := Edit(initial_text, start_text, termination_text, to_insert)
					So(edited_text, ShouldEqual, final_text)
					So(edited, ShouldBeFalse)
				})

			})

			Convey("Or with an existing text", func() {

				existing_text := `I want to be edited.
# EDITLIB START
replace me
# EDITLIB TERMINATE`

				Convey("The text is edited in place", func() {
					edited_text, edited := Edit(existing_text, start_text, termination_text, to_insert)
					So(edited_text, ShouldEqual, final_text)
					So(edited, ShouldBeFalse)
				})

			})
		})

		Convey("With an already edited text", func() {

			Convey("Should do nothing", nil)

		})

	})

	Convey("Given start and termination texts", t, func() {

		Convey("With an existing text", func() {

			Convey("Should remove start and termination texts and everything in between", nil)

		})

		Convey("With an initial text", func() {

			Convey("Should do nothing", nil)

		})

	})

}

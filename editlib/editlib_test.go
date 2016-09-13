package editlib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEdit(t *testing.T) {

	Convey("Given start, termination and some text to insert", t, func() {

		Convey("With an initial text", func() {

			Convey("Should append new text", nil)

		})

		Convey("With an existing text and some text to insert", func() {

			Convey("Should edit existing text in place", nil)

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

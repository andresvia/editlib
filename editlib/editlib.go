package editlib

import (
	"io"
	"io/ioutil"
	"strings"
)

func Edit(dst io.Writer, src io.Reader, start, end, insert string) (err error) {
	var source_bytes []byte
	source_bytes, err = ioutil.ReadAll(src)
	final_value := SimpleEdit(string(source_bytes), start, end, insert)
	_, err = io.WriteString(dst, final_value)
	return
}

func SimpleEdit(initial_value, start, end, insert string) (final_value string) {
	initial_parts := strings.SplitN(initial_value, start+"\n", 2)

	editing := false
	if len(initial_parts) == 1 {
		initial_parts = append(initial_parts, "")
	} else {
		editing = true
	}

	if !strings.HasSuffix(initial_parts[0], "\n") && initial_parts[0] != "" && editing {
		// start substring found
		final_value = initial_value
		return
	} else if !strings.HasSuffix(initial_parts[0], "\n") && initial_parts[0] != "" {
		start = "\n" + start
	} else if strings.HasSuffix(initial_parts[0], "\n") && !editing {
		start = "\n" + start
	}

	final_parts := strings.SplitN(initial_parts[1], "\n"+end, 2)

	if len(final_parts) == 1 {
		// end substring found
		final_parts = append(final_parts, "")
		editing = false
	}

	if !strings.HasPrefix(final_parts[1], "\n") && final_parts[1] != "" {
		final_value = initial_value
		return
	}

	if insert != "" {
		final_value = initial_parts[0] + start + "\n" + insert + "\n" + end + final_parts[1]
	} else if editing {
		nl := ""
		if initial_parts[0] != "" && final_parts[1] != "" {
			nl = "\n"
		}
		final_value = strings.TrimSuffix(initial_parts[0], "\n") + nl + strings.TrimPrefix(final_parts[1], "\n")
	} else {
		final_value = initial_parts[0] + final_parts[1]
	}

	return
}

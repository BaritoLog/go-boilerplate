package execkit

import (
	"fmt"
	"io"
)

// Run and print the output
func Print(w io.Writer, cmds ...Cmd) (err error) {
	for _, cmd := range cmds {
		var out []byte
		out, err = cmd.Output()
		if err != nil {
			return
		}

		fmt.Fprintln(w, string(out))
	}

	return
}

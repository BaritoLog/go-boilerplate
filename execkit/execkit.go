package execkit

// Run and print the output
func Print(printer Printer, cmds ...Cmd) (err error) {
	for _, cmd := range cmds {
		var out []byte
		out, err = cmd.Output()
		if err != nil {
			return
		}

		printer.Print(string(out))
	}

	return
}

// acmedot prints the location of dot in the current acme window.
//     usage: acmedot [format1 [format2]]
//         format1 - specifies the output format. It will be 
//                   used as if called by 
//                   fmt.Printf(format1, start, end).
//         format2 - specifies the output format where start and
//                   end are the same. It will be used as if called
//                   by fmt.Printf(format2, start)
// The default format is "%d,%d\n".
package main

import (
	"fmt"
	"os"
	"strconv"

	"9fans.net/go/acme"
)

func main() {
	if os.Getenv("winid") == "" {
		error("no winid, not acme window")
	}
	id, err := strconv.Atoi(os.Getenv("winid"))
	if err != nil {
		error("cannot parse winid %q: %s", os.Getenv("winid"), err)
	}
	w, err := acme.Open(id, nil)
	if err != nil {
		error("cannot open window %d: %s", id, err)
	}
	// Need the address file open to get it correctly
	if _, _, err := w.ReadAddr(); err != nil {
		error("cannot read addr: %s", err)
	}
	if err := w.Ctl("addr=dot"); err != nil {
		error("cannot write ctl: %s", err)
	}
	start, end, err := w.ReadAddr()
	if err != nil {
		error("cannot read addr: %s", err)
	}
	if len(os.Args) > 2 && start == end {
		fmt.Printf(os.Args[2], start)
	} else if len(os.Args) > 1 {
		fmt.Printf(os.Args[1], start, end)
	} else {
		fmt.Printf("%d,%d\n", start, end)
	}
}

func error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

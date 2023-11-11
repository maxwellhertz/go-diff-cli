package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	deltaFlagDefault := false
	deltaFlagDescription := "Describe the operations required to transform src into dest."
	deltaFlagLong := flag.Bool("delta", deltaFlagDefault, deltaFlagDescription)
	deltaFlagShort := flag.Bool("d", deltaFlagDefault, deltaFlagDescription)
	multiLineFlagDefault := false
	multiLineFlagDescription := "Whether the input strings are multi-line."
	multiLineFlagLong := flag.Bool("lines", multiLineFlagDefault, multiLineFlagDescription)
	multiLineFlagShort := flag.Bool("l", multiLineFlagDefault, multiLineFlagDescription)
	flag.Parse()

	texts := flag.Args()
	if len(texts) != 2 {
		fmt.Println("Please specify two strings to compare.")
		os.Exit(-1)
		return
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(texts[0], texts[1], *multiLineFlagShort || *multiLineFlagLong)
	if dmp.DiffLevenshtein(diffs) == 0 {
		os.Exit(0)
		return
	}
	
	fmt.Println(dmp.DiffPrettyText(diffs))
	if *deltaFlagShort || *deltaFlagLong {
		fmt.Println(dmp.DiffToDelta(diffs))
	}
}

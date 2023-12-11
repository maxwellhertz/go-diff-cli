package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/nsf/jsondiff"
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

	if isJsonString(texts[0]) && isJsonString(texts[1]) {
		diffJson(texts[0], texts[1])
		return
	}
	diffSimpleText(texts[0], texts[1], *multiLineFlagShort || *multiLineFlagLong, *deltaFlagShort || *deltaFlagLong)
}

func isJsonString(text string) bool {
	var raw json.RawMessage
	return json.Unmarshal([]byte(text), &raw) == nil
}

func diffSimpleText(src string, dest string, multiLine bool, delta bool) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(src, dest, multiLine)
	if dmp.DiffLevenshtein(diffs) == 0 {
		os.Exit(0)
		return
	}

	fmt.Println(dmp.DiffPrettyText(diffs))
	if delta {
		fmt.Println(dmp.DiffToDelta(diffs))
	}
}

func diffJson(src string, dest string) {
	opts := jsondiff.DefaultConsoleOptions()
	_, result := jsondiff.Compare([]byte(src), []byte(dest), &opts)
	fmt.Println(result)
}

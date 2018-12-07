package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl"
)

func main() {
	fail := false

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		search := arg
		if info, err := os.Stat(arg); err == nil && info.IsDir() {
			search = fmt.Sprintf("%s/*.tf", arg)
		}

		files, err := filepath.Glob(search)
		if err != nil {
			fmt.Println("error finding files:", err)
		}

		for _, filename := range files {
			fmt.Println("Checking \"" + filename + "\"")
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Println("error reading file:", err)
				fail = true
				continue
			}

			_, err = hcl.Parse(string(file))
			if err != nil {
				fmt.Println("error parsing file: ", err)
				fail = true
				continue
			}
		}
	}

	if fail {
		os.Exit(-1)
	}
}

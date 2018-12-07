package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl"
)

func main() {
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
			fmt.Printf("error finding files: %s", err)
		}

		for _, filename := range files {
			fmt.Printf("Checking %s ... ", filename)
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Printf("error reading file: %s\n", err)
				break
			}

			_, err = hcl.Parse(string(file))
			if err != nil {
				fmt.Printf("error parsing file: %s\n", err)
				break
			}
		}
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
	
	"gopkg.in/yaml.v2"
	devfilepkg "github.com/devfile/library/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/parser"
)

func main() {
	parserTest()
}

func parserTest() {
	var args parser.ParserArgs
	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "http") {
			args = parser.ParserArgs{
				URL: os.Args[1],
			}
		} else {
			args = parser.ParserArgs{
				Path: os.Args[1],
			}
		}
		fmt.Println("parsing devfile from " + os.Args[1])

	} else {
		args = parser.ParserArgs{
			Path: "devfile.yaml",
		}
		fmt.Println("parsing devfile from ./devfile.yaml")
	}
	devfile, warning, err := devfilepkg.ParseDevfileAndValidate(args)
	if err != nil {
		fmt.Println(err)
	} else {
		if len(warning.Commands) > 0 || len(warning.Components) > 0 || len(warning.Projects) > 0 || len(warning.StarterProjects) > 0 {
			fmt.Printf("top-level variables were not substituted successfully %+v\n", warning)
		}
		writeDevfile(devfile, args.Path)
	}
	
}

func writeDevfile(devfile parser.DevfileObj, filename string) {

    // Encode data into YAML format
	yamlData, err := yaml.Marshal(devfile.Data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

    fmt.Printf("Writing devfile: %v\n", "flattened-"+filename)
    
	fs := devfile.Ctx.GetFs()
	err = fs.WriteFile("flattened-"+filename, yamlData, 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
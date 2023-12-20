//go:build exclude
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/rbxl"
	. "github.com/robloxapi/rbxfile/declare"
)

func handleDir(name string, path string) *rbxfile.Instance {
	entries, _ := os.ReadDir(path)

	instances := make([]*rbxfile.Instance, 0, 0)

	for _, entry := range entries {
		name := entry.Name()

		if entry.IsDir() {
			//Handle recursive directory
			instances = append(instances, handleDir(name, filepath.Join(path, name)))
			continue
		}

		var insname string
		var found bool
		var class string

		if insname, found = strings.CutSuffix(name, ".init.lua"); insname != "" && found {
			class = "Script"
		} else if insname, found = strings.CutSuffix(name, ".lua"); insname != "" && found {
			class = "ModuleScript"
		}

		if !found {
			fmt.Fprintf(os.Stderr, "unknown class for %s", name)
			os.Exit(1)
		}

    		b, err := os.ReadFile(filepath.Join(path, name))
    		if err != nil {
       			fmt.Print(err)
       			os.Exit(1)
    		}

		instances = append(instances, Instance(class,
			Property("Name", String, insname),
			Property("Source", ProtectedString, b),
		).Declare())
	}

	folderinst := Instance("Folder",
		Property("Name", String, name),
	).Declare()

	folderinst.Children = instances

	return folderinst
}

func main() {
	rootinst := handleDir("VinegarBridge Plugin", "plugin")

	root := &rbxfile.Root{
		Instances: []*rbxfile.Instance{rootinst},
		Metadata:  map[string]string{},
	}

	encoder := rbxl.Encoder{
		Mode: rbxl.Model,
	}

	f, _ := os.OpenFile("vinegarbridge.rbxm", os.O_RDWR|os.O_CREATE, 0755)

	encoder.Encode(f, root)
	f.Close()


}
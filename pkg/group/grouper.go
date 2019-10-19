package group

import "os"

//Grouper is an interfacce that is required for file.Assigner to group files in directory
type Grouper interface {
	Group() (map[string][]os.FileInfo, error)
}

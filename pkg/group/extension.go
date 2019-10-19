package group

import (
	"os"
	"strings"
)

//ExtensionGroup holds file base on file extension
type ExtensionGroup struct {
	Datas []os.FileInfo
}

//NewExtensionGroup creates a new TypeGroup struct
func NewExtensionGroup(fis []os.FileInfo) *ExtensionGroup {
	return &ExtensionGroup{
		Datas: fis,
	}
}

//Group implements Grouper interface
//creates mapping between group and list of FilesInfo base on file extension
func (eg ExtensionGroup) Group() (map[string][]os.FileInfo, error) {
	mapping := make(map[string][]os.FileInfo)
	for _, fi := range eg.Datas {
		sep := strings.Split(fi.Name(), ".")
		if len(sep) > 1 {
			extension := sep[len(sep)-1]
			mapping[extension] = append(mapping[extension], fi)
		}
	}
	return mapping, nil
}

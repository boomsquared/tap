package group

import (
	"os"
	"strings"
)

//ExtensionGroup holds file base on file extension
type ExtensionGroup struct {
	Datas map[string][]os.FileInfo
}

//NewExtensionGroup creates a new TypeGroup struct
func NewExtensionGroup(fis []os.FileInfo) *ExtensionGroup {
	datas := make(map[string][]os.FileInfo)
	for _, fi := range fis {
		sep := strings.Split(fi.Name(), ".")
		if len(sep) > 1 {
			extension := sep[len(sep)-1]
			datas[extension] = append(datas[extension], fi)
		}
	}
	return &ExtensionGroup{
		Datas: datas,
	}
}

//Group creates mapping between group and list of files base on file extension
func (eg ExtensionGroup) Group() map[string][]os.FileInfo {
	return eg.Datas
}

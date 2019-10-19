package group

import (
	"os"
	"path"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

//EXIFGroup holds file base on Exif properties
type EXIFGroup struct {
	Datas    []os.FileInfo
	BasePath string
	By       exif.FieldName
}

//NewEXIFGroup creates a new EXIFGroup struct
func NewEXIFGroup(fis []os.FileInfo, basePath string, by exif.FieldName) *EXIFGroup {
	return &EXIFGroup{
		Datas:    fis,
		BasePath: basePath,
		By:       by,
	}
}

//Group implemented Gropper interface
//creates mapping between group and list of FilesInfo base on file exif fieldname
func (eg EXIFGroup) Group() (map[string][]os.FileInfo, error) {
	mapping := make(map[string][]os.FileInfo)
	for _, fi := range eg.Datas {
		if fi.IsDir() {
			continue
		}
		fPath := path.Join(eg.BasePath, fi.Name())
		x, err := getEXIF(fPath)
		if err != nil {
			continue
			// return nil, errors.Wrap(err, "unable to extract EXIF")
		}
		tag, err := x.Get(eg.By)
		if err != nil {
			continue
		}
		key := strings.ReplaceAll(tag.String(), "/", "|")
		key = strings.ReplaceAll(key, "\"", "")
		mapping[key] = append(mapping[key], fi)
	}
	return mapping, nil
}

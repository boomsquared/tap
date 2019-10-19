package file

import (
	"fmt"
	"path"
	"sort"
	"strings"

	filesort "github.com/boomsquared/tap/pkg/sort"
	"github.com/pkg/errors"
)

//Renamer rename files and directory
type Renamer struct {
	objects  filesort.FileSort
	basePath string
	fileOp   FileOp
}

//NewRenamer returns a new file renamer object
func NewRenamer(fo FileOp) *Renamer {
	return &Renamer{
		fileOp: fo,
	}
}

//Load files info
func (r *Renamer) Load(path, by string) error {
	fis, err := r.fileOp.ReadDir(path)
	if err != nil {
		return errors.Wrap(err, "unable to read from path")
	}
	var objs filesort.FileSort
	switch by {
	case "size":
		objs = filesort.NewSizeSort(fis)
	default:
		objs = filesort.NewSizeSort(fis)
	}

	r.objects = objs
	r.basePath = path
	return nil

}
func (r *Renamer) generateName(i int, prefix string) string {
	return path.Join(r.basePath, fmt.Sprintf("%s-%d", prefix, i+1))
}

//Rename file according to sort
func (r *Renamer) Rename(prefix string) error {
	sort.Sort(r.objects)
	for i := 0; i < r.objects.Len(); i++ {
		fname := r.objects.Iterate()[i].Name()
		src := path.Join(r.basePath, fname)
		p := strings.Split(fname, ".")
		dst := ""
		if len(p) > 1 {
			dst = fmt.Sprintf("%s.%s", r.generateName(i, prefix), p[len(p)-1])
		} else {
			dst = r.generateName(i, prefix)
		}
		err := r.fileOp.Rename(src, dst)
		if err != nil {
			return errors.Wrap(err, "unable to rename files")
		}
	}
	return nil
}

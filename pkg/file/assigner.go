package file

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/boomsquared/tap/pkg/group"
	"github.com/pkg/errors"
)

//Assigner groups files in directory
type Assigner struct {
	objects  group.Grouper
	basePath string
}

//NewAssigner returns a new file Assigner object
func NewAssigner() *Assigner {
	return &Assigner{}
}

//Load files from path
func (g *Assigner) Load(path string) error {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return errors.Wrap(err, "unable to read from path")
	}
	objs := group.NewExtensionGroup(fis)
	g.objects = objs
	g.basePath = path
	return nil
}

//Assign files base on key of Grouper.Group
func (g *Assigner) Assign() error {
	mapping := g.objects.Group()
	for gName, fInfos := range mapping {
		dirPath := path.Join(g.basePath, gName)
		if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
			return errors.Wrap(err, "unable to create directory")
		}
		for _, fInfo := range fInfos {
			src := path.Join(g.basePath, fInfo.Name())
			dst := path.Join(g.basePath, gName, fInfo.Name())
			if err := os.Rename(src, dst); err != nil {
				return errors.Wrap(err, "unable to rename files")
			}

		}

	}
	return nil
}

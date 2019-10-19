package file

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/boomsquared/tap/pkg/group"
	"github.com/pkg/errors"
	"github.com/rwcarlsen/goexif/exif"
)

//Assigner groups files in directory
type Assigner struct {
	objects  group.Grouper
	basePath string
	fileOp   FileOp
}

//NewAssigner returns a new file Assigner object
func NewAssigner(fo FileOp) *Assigner {
	return &Assigner{
		fileOp: fo,
	}
}

//Load files from path
func (g *Assigner) Load(path, by string) error {
	fis, err := ioutil.ReadDir(path)
	if err != nil || len(fis) == 0 {
		return errors.Wrap(err, "unable to read from path")
	}
	var objs group.Grouper
	switch by {
	case "device":
		objs = group.NewEXIFGroup(fis, path, exif.Model)
	case "extension":
		objs = group.NewExtensionGroup(fis)
	case "fnumber":
		objs = group.NewEXIFGroup(fis, path, exif.FNumber)
	case "iso":
		objs = group.NewEXIFGroup(fis, path, exif.ISOSpeedRatings)
	case "lens":
		objs = group.NewEXIFGroup(fis, path, exif.LensModel)
	case "shutterspeed":
		objs = group.NewEXIFGroup(fis, path, exif.ShutterSpeedValue)

	default:
		objs = group.NewEXIFGroup(fis, path, exif.Model)
	}
	g.objects = objs
	g.basePath = path
	return nil
}

//Assign files base on key of Grouper.Group
func (g *Assigner) Assign() error {
	mapping, err := g.objects.Group()
	if err != nil {
		return errors.Wrap(err, "unable to group files")
	}
	for gName, fInfos := range mapping {
		dirPath := path.Join(g.basePath, gName)
		if err := g.fileOp.MkdirAll(dirPath, os.ModePerm); err != nil {
			return errors.Wrap(err, "unable to create directory")
		}
		for _, fInfo := range fInfos {
			src := path.Join(g.basePath, fInfo.Name())
			dst := path.Join(g.basePath, gName, fInfo.Name())
			if err := g.fileOp.Rename(src, dst); err != nil {
				return errors.Wrap(err, "unable to rename files")
			}

		}

	}
	return nil
}

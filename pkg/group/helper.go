package group

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rwcarlsen/goexif/exif"
)

func getEXIF(filePath string) (*exif.Exif, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open file")
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode EXIF")
	}
	return x, nil

}

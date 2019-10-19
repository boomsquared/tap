package file

import (
	"io/ioutil"
	"os"
)

//FileOp is an interface that interacts with the file system
//This interface exist to make code unittestable
type FileOp interface {
	ReadDir(dirname string) ([]os.FileInfo, error)
	MkdirAll(name string, perm os.FileMode) error
	Rename(oldpath string, newpath string) error
}

//Operation implements the fileOp interface
type Operation struct {
}

//ReadDir proxy to ioutil.ReadDir
func (o Operation) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

//MkdirAll proxy to os.Mkdir
func (o Operation) MkdirAll(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

//Rename proxy to os.Rename
func (o Operation) Rename(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

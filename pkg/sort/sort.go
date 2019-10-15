package sort

import (
	"os"
	"sort"
)

//FileSort interface that specify method that is needed to use renames
type FileSort interface {
	Iterate() []os.FileInfo
	sort.Interface
}

package sort

import "os"

//SizeSort sort files base on image metadata
type SizeSort struct {
	Datas []os.FileInfo
}

//NewSizeSort create a new SizeSort struct
func NewSizeSort(fis []os.FileInfo) *SizeSort {
	return &SizeSort{
		fis,
	}
}

//Iterate over file info
func (ss SizeSort) Iterate() []os.FileInfo {
	return ss.Datas
}

// Len is the number of elements in the collection.
func (ss SizeSort) Len() int {
	return len(ss.Datas)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (ss SizeSort) Less(i, j int) bool {
	return ss.Datas[i].Size() < ss.Datas[j].Size()
}

// Swap swaps the elements with indexes i and j.
func (ss *SizeSort) Swap(i, j int) {
	ss.Datas[i], ss.Datas[j] = ss.Datas[j], ss.Datas[i]
}

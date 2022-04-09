package Chore

type Size int

const (
	Small   Size = iota //0
	Medium              //1
	Large               //2
	Unknown = -1
)

func createSize(s string) Size {
	if s == "Small" {
		return Small
	} else if s == "Medium" {
		return Medium
	} else if s == "Large" {
		return Large
	} else {
		return Unknown
	}
}

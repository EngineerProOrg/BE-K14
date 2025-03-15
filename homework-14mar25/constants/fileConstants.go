package constants

const FileTxtPersons = "persons.txt"
const FileJsonPersons = "persons.json"

type FileType int

const (
	FileTxt  FileType = iota // 0
	FileJson                 // 1
)

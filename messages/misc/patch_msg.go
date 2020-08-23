package misc

import "fmt"

/*
File contains Structure definitions for Patch protocol (8).
*/

/*
Patch Client sends this message to request the latest file list.
Patch Server sends this message to tell the client where to get it.
*/
type LatestFileList struct {
	LatestVersion uint32
	ListFileName  string
	ListFileType  uint32
	ListFileTime  uint32
	ListFileSize  uint32
	ListFileCRC   uint32
	ListFileURL   string
	URLPrefix     string
	URLSuffix     string
}

func (m LatestFileList) String() string {
	return fmt.Sprintf("%T", m)
}

/*
Patch Client sends this message to request the latest file list.
Patch Server sends this message to tell the client where to get it.
Same as above but with locale
*/
type LatestFileListV2 struct {
	LatestVersion uint32
	ListFileName  string
	ListFileType  uint32
	ListFileTime  uint32
	ListFileSize  uint32
	ListFileCRC   uint32
	ListFileURL   string
	URLPrefix     string
	URLSuffix     string
	Locale        string
}

func (m LatestFileListV2) String() string {
	return fmt.Sprintf("%T", m)
}

/*
Patch Client sends this message with the current version of a specific package.
Patch Server replies with the next version of the specific package.
*/

type NextVersion struct {
	PkgName   string
	Version   int32
	URLPrefix string
	Filename  string
	FileType  int32
}

func (m NextVersion) String() string {
	return fmt.Sprintf("%T", m)
}

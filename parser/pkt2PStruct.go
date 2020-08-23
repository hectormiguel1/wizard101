package parser

import (
	"wizard101/messages"
	"wizard101/messages/misc"
)

const (
	LatestFile   = 1
	LatestFileV2 = 2
	NextVersion  = 3
)

func BuildPatchMessages(pkt []byte, msgOrder uint8) messages.ProtocolMessage {
	switch msgOrder {
	case LatestFile:
		return buildLatestFile(pkt)
	case LatestFileV2:
		return buildLatestFileV2(pkt)
	case NextVersion:
		return buildNextVersion(pkt)
	default:
		return nil
	}
}

func buildNextVersion(pkt []byte) messages.ProtocolMessage {
	pkt, _ = ReadUShort(pkt)
	pkt, pkgName := readString(pkt)
	pkt, version := ReadInt(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, urlPrefix := readString(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, fileName := readString(pkt)
	pkt, fileType := ReadInt(pkt)

	return misc.NextVersion{
		PkgName:   pkgName,
		Version:   version,
		URLPrefix: urlPrefix,
		Filename:  fileName,
		FileType:  fileType,
	}
}

func buildLatestFileV2(pkt []byte) messages.ProtocolMessage {
	pkt, latestVersion := ReadUInt(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, listFileName := readString(pkt)
	pkt, listFileType := ReadUInt(pkt)
	pkt, listFileTime := ReadUInt(pkt)
	pkt, listFileSize := ReadUInt(pkt)
	pkt, listFileCRC := ReadUInt(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, listFileURL := readString(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, urlPrefix := readString(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, urlSuffix := readString(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, locale := readString(pkt)

	return misc.LatestFileListV2{
		LatestVersion: latestVersion,
		ListFileName:  listFileName,
		ListFileType:  listFileType,
		ListFileTime:  listFileTime,
		ListFileSize:  listFileSize,
		ListFileCRC:   listFileCRC,
		ListFileURL:   listFileURL,
		URLPrefix:     urlPrefix,
		URLSuffix:     urlSuffix,
		Locale:        locale,
	}
}

func buildLatestFile(pkt []byte) messages.ProtocolMessage {
	pkt, latestVersion := ReadUInt(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, listFileName := readString(pkt)
	pkt, listFileType := ReadUInt(pkt)
	pkt, listFileTime := ReadUInt(pkt)
	pkt, listFileSize := ReadUInt(pkt)
	pkt, listFileCRC := ReadUInt(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, listFileURL := readString(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, urlPrefix := readString(pkt)
	pkt, _ = ReadUShort(pkt)
	pkt, urlSuffix := readString(pkt)

	return misc.LatestFileList{
		LatestVersion: latestVersion,
		ListFileName:  listFileName,
		ListFileType:  listFileType,
		ListFileTime:  listFileTime,
		ListFileSize:  listFileSize,
		ListFileCRC:   listFileCRC,
		ListFileURL:   listFileURL,
		URLPrefix:     urlPrefix,
		URLSuffix:     urlSuffix,
	}
}

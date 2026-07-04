package dirops

import (
	"fmt"
	"os"
	"time"
)

type MockFileInfo struct { // Should implement os.FileInfo
	name    string
	isDir   bool
	modTime time.Time
}

func (aMockFileInfo MockFileInfo) Name() string {
	return aMockFileInfo.name
}

func (aMockFileInfo MockFileInfo) Size() int64 {
	return 1 // Doesn't matter, as we're not checking this
}

func (aMockFileInfo MockFileInfo) Mode() os.FileMode {
	if aMockFileInfo.isDir {
		return 1777 // dir bit is set
	}
	return 0o777 // dir bit is not set
}

func (aMockFileInfo MockFileInfo) ModTime() time.Time {
	return aMockFileInfo.modTime
}

func (aMockFileInfo MockFileInfo) IsDir() bool {
	return aMockFileInfo.isDir
}

func (aMockFileInfo MockFileInfo) Sys() any {
	return nil
}

type MockEntry struct { // Should implement os.DirEntry
	shouldInfoFail bool
	mockFileInfo   MockFileInfo
}

func (aMockEntry MockEntry) Name() string {
	return aMockEntry.mockFileInfo.Name()
}

func (aMockEntry MockEntry) IsDir() bool {
	return aMockEntry.mockFileInfo.IsDir()
}

func (aMockEntry MockEntry) Type() os.FileMode {
	return aMockEntry.mockFileInfo.Mode()
}

func (aMockEntry MockEntry) Info() (os.FileInfo, error) {
	if aMockEntry.shouldInfoFail {
		return nil, fmt.Errorf("Mock error for file %q", aMockEntry.mockFileInfo.Name())
	}

	return aMockEntry.mockFileInfo, nil
}

func createMockEntry(name string, isDir bool, shouldInfoFail bool, modTime time.Time) MockEntry {
	aMockFileInfo := MockFileInfo{
		name,
		isDir,
		modTime,
	}

	return MockEntry{
		shouldInfoFail: shouldInfoFail,
		mockFileInfo:   aMockFileInfo,
	}
}

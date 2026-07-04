package gopro

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"slices"
	"strings"
)

func GetGoProPath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error trying to get current user %w", err)
	}

	uid := currentUser.Uid
	mtpPath := filepath.Join("/", "run", "user", uid, "gvfs")

	mtpDirEntries, err := os.ReadDir(mtpPath)
	if err != nil {
		return "", fmt.Errorf("error trying to get list of directories in %q", mtpPath)
	}

	mtpDirEntries = slices.DeleteFunc(mtpDirEntries, func(mtpDirEntry os.DirEntry) bool {
		isDir := mtpDirEntry.IsDir()
		if !isDir {
			return true
		}

		dirName := mtpDirEntry.Name()
		if !strings.HasPrefix(dirName, "mtp") {
			return true
		}

		if !strings.Contains(dirName, "GoPro") {
			return true
		}

		return false
	})

	if len(mtpDirEntries) == 0 {
		fmt.Println("No matches found for criteria:")
		fmt.Println("- is directory?")
		fmt.Println("- directory starts with 'mtp'?")
		fmt.Println("- directory contains substring 'GoPro'?")
		return "", fmt.Errorf("no directories found in MTP path %q", mtpPath)
	}

	goProVideosPath := filepath.Join(mtpPath, mtpDirEntries[0].Name(), "GoPro MTP Client Disk Volume", "DCIM", "100GOPRO")

	goProVideos, err := os.ReadDir(goProVideosPath)
	if err != nil {
		return "", fmt.Errorf("error when trying to read directory %q: %w", goProVideosPath, err)
	}

	if len(goProVideos) == 0 {
		return "", fmt.Errorf("no files found in directory %q", goProVideosPath)
	}

	return goProVideosPath, nil
}

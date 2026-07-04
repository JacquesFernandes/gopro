package dirops

import (
	"os"
	"slices"
	"time"
)

type Entry = os.DirEntry

type DirDateGroup struct {
	Date    time.Time
	Entries []Entry
}

func GroupDirEntriesByDate(entries []Entry) ([]DirDateGroup, error) {
	dateGroupings := make([]DirDateGroup, 0)

	for _, dirEntry := range entries {
		fileInfo, err := dirEntry.Info()
		if err != nil {
			return make([]DirDateGroup, 0), err
		}

		fileModtime := fileInfo.ModTime()
		dateKey := time.Date(fileModtime.Year(), fileModtime.Month(), fileModtime.Day(), 0, 0, 0, 0, fileModtime.Location())
		possibleExistingGroupIndex := slices.IndexFunc(dateGroupings, func(dateGroup DirDateGroup) bool {
			return dateGroup.Date.Equal(dateKey)
		})

		if possibleExistingGroupIndex == -1 {
			dateGroup := DirDateGroup{
				Date: dateKey,
				Entries: []Entry{
					dirEntry,
				},
			}

			dateGroupings = append(dateGroupings, dateGroup)
		} else {
			dateGroupings[possibleExistingGroupIndex].Entries = append(dateGroupings[possibleExistingGroupIndex].Entries, dirEntry)
		}
	}

	return dateGroupings, nil
}

package dirops

import (
	"fmt"
	"os"
	"time"
)

type Entry = os.DirEntry

// TODO: Look at "pointers as values" or "Read-Modify-Write"
type (
	key          = string
	value        = []Entry
	DirDateGroup = map[key]value
)

func timeToKey(aTime time.Time) key {
	return fmt.Sprintf("%d-%d-%d", aTime.Year(), aTime.Month(), aTime.Day())
}

func GroupDirEntriesByDate(entries []Entry) (DirDateGroup, error) {
	var dirDateGroup DirDateGroup = make(DirDateGroup)

	// TODO: Write internals
	return dirDateGroup, nil
}

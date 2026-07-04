package dirops

import (
	"testing"
	"time"
)

func groupDateTester(baseTime time.Time, t *testing.T) func(groupName string, expectedDay int, group DirDateGroup) {
	return func(groupName string, expectedDay int, group DirDateGroup) {
		year, month, day := group.Date.Date()

		if year != baseTime.Year() {
			t.Errorf("%q group :: year :: expected %d, got %d", groupName, baseTime.Year(), year)
		}

		if month != baseTime.Month() {
			t.Errorf("%q group :: month :: expected %q, got %q", groupName, baseTime.Month(), month)
		}

		if day != expectedDay {
			t.Errorf("%q group :: day :: expected %d, got %d", groupName, baseTime.Day(), day)
		}
	}
}

func checkDateGroup(group DirDateGroup, baseTime time.Time, t *testing.T) {
}

func TestGroupDirEntriesByDate(t *testing.T) {
	baseTime := time.Now()

	t.Run("No errors, groups okay", func(t *testing.T) {
		mockEntries := []Entry{
			createMockEntry("uno", time.Date(baseTime.Year(), baseTime.Month(), 10, 0, 0, 0, 0, baseTime.Location()), false),
			createMockEntry("dos", time.Date(baseTime.Year(), baseTime.Month(), 10, 0, 0, 0, 0, baseTime.Location()), false),
			createMockEntry("foo", time.Date(baseTime.Year(), baseTime.Month(), 12, 0, 0, 0, 0, baseTime.Location()), false),
			createMockEntry("bar", time.Date(baseTime.Year(), baseTime.Month(), 12, 0, 0, 0, 0, baseTime.Location()), false),
			createMockEntry("nemo", time.Date(baseTime.Year(), baseTime.Month(), 14, 0, 0, 0, 0, baseTime.Location()), false),
		}

		dateGroups, err := GroupDirEntriesByDate(mockEntries)
		if err != nil {
			t.Errorf("expected err to be nil, got non-nil %q", err)
		}

		if len(dateGroups) != 3 {
			t.Errorf("Expcted 3 groups, got %d", len(dateGroups))
		}

		firstGroup := dateGroups[0]
		secondGroup := dateGroups[1]
		thirdGroup := dateGroups[2]

		testGroupDate := groupDateTester(baseTime, t)
		testGroupDate("first", 10, firstGroup)
		testGroupDate("second", 12, secondGroup)
		testGroupDate("third", 14, thirdGroup)

		if len(firstGroup.Entries) != 2 {
			t.Errorf("Expected firstGroup to have 2 entries, but got %d", len(firstGroup.Entries))
		}

		if len(secondGroup.Entries) != 2 {
			t.Errorf("Expected secondGroup to have 2 entries, but got %d", len(secondGroup.Entries))
		}

		if len(thirdGroup.Entries) != 1 {
			t.Errorf("Expected thirdGroup to have 1 entries, but got %d", len(thirdGroup.Entries))
		}
	})

	t.Run("Internal error, should return an empty slice", func(t *testing.T) {
		mockEntries := []Entry{
			createMockEntry("uno", time.Now(), false),
			createMockEntry("dos", time.Now(), true),
		}

		dateGroups, err := GroupDirEntriesByDate(mockEntries)

		if len(dateGroups) > 0 {
			t.Errorf("Expected dateGroups list to be empty, got length %d", len(dateGroups))
		}

		if err == nil {
			t.Error("Expected an error, but got none?")
		}
	})
}

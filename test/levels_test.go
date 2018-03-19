package meerkats__test

import (
	"testing"

	. "github.com/Tlantic/meerkats"
)

func TestLevel_String(t *testing.T) {

	if LevelTrace.String() != Levels[LevelTrace] {
		t.Fail()
	}

	if LevelDebug.String() != Levels[LevelDebug] {
		t.Fail()
	}

	if LevelInfo.String() != Levels[LevelInfo] {
		t.Fail()
	}

	if LevelWarning.String() != Levels[LevelWarning] {
		t.Fail()
	}

	if LevelError.String() != Levels[LevelError] {
		t.Fail()
	}

	if LevelFatal.String() != Levels[LevelFatal] {
		t.Fail()
	}

	if LevelPanic.String() != Levels[LevelPanic] {
		t.Fail()
	}
}

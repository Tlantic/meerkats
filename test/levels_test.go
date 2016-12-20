package meerkats__test

import (
	"testing"
	. "github.com/Tlantic/meerkats"
)

func TestLevel_String(t *testing.T) {

	if ( TRACE.String() != Levels[TRACE] ) {
		t.Fail()
	}

	if ( DEBUG.String() != Levels[DEBUG] ) {
		t.Fail()
	}

	if ( INFO.String() != Levels[INFO] ) {
		t.Fail()
	}

	if ( WARNING.String() != Levels[WARNING] ) {
		t.Fail()
	}

	if ( ERROR.String() != Levels[ERROR] ) {
		t.Fail()
	}

	if ( FATAL.String() != Levels[FATAL] ) {
		t.Fail()
	}

	if ( PANIC.String() != Levels[PANIC] ) {
		t.Fail()
	}
}

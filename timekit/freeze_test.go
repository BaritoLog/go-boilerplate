package timekit

import (
	"testing"
	"time"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestFreeze(t *testing.T) {
	err := Freeze("2017-10-02T10:00:00-05:00")
	defer Unfreeze()

	FatalIfError(t, err)
	FatalIf(t, time.Now().Format(time.RFC3339) != "2017-10-02T10:00:00-05:00", "wrong time.Now()")
}

func TestFreeze_Error(t *testing.T) {
	err := Freeze("wrong-formatted")
	FatalIfWrongError(t, err, `parsing time "wrong-formatted" as "2006-01-02T15:04:05Z07:00": cannot parse "wrong-formatted" as "2006"`)
}

func TestUnfreeze(t *testing.T) {
	err := Freeze("2017-10-02T10:00:00-05:00")
	Unfreeze()

	FatalIfError(t, err)
	FatalIf(t, time.Now().Format(time.RFC3339) == "2017-10-02T10:00:00-05:00", "wrong time.Now()")
}

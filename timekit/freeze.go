package timekit

import (
	"time"

	"github.com/bouk/monkey"
)

var (
	patch *monkey.PatchGuard
)

func Freeze(sometime string) (err error) {

	d, err := time.Parse(time.RFC3339, sometime)
	if err != nil {
		return
	}

	patch = monkey.Patch(
		time.Now,
		func() time.Time {
			return d
		},
	)

	return
}

func Unfreeze() {
	if patch != nil {
		patch.Unpatch()
	}

}

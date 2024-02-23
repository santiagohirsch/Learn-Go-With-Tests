package interactions_test

import (
	"Learn-Go-With-Tests/domain/interactions"
	"Learn-Go-With-Tests/specifications"
	"testing"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(interactions.Curse),
	)
}

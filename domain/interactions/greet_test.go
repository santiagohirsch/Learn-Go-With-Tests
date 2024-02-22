package interactions_test

import (
	"Learn-Go-With-Tests/domain/interactions"
	"Learn-Go-With-Tests/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}

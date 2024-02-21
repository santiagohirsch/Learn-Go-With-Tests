package go_specs_greet_test

import (
	go_specs_greet "Learn-Go-With-Tests/domain/interactions"
	"Learn-Go-With-Tests/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(go_specs_greet.Greet),
	)
}

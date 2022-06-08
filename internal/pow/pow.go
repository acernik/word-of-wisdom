package pow

import (
	"github.com/umahmood/hashcash"
)

// Generator interface defines all the methods for generating a Proof of Work solution.
type Generator interface {
	GetPowSolution() (string, error)
	Verify(solution string) (bool, error)
}

// generator is the type that implements the Generator interface.
type generator struct {
	hc *hashcash.Hashcash
}

// New returns the value of type that implements the Generator interface.
func New() Generator {
	hc, _ := hashcash.New(
		&hashcash.Resource{
			Data:          "someone@gmail.com",
			ValidatorFunc: validateResource,
		},
		nil,
	)

	return &generator{
		hc: hc,
	}
}

// validateResource is used when it is necessary to validate the resource. E.g. to check if email is in database.
func validateResource(resource string) bool {
	return true
}

// GetPowSolution generates the Proof of Work solution.
func (g *generator) GetPowSolution() (string, error) {
	var solution string

	hc, err := hashcash.New(
		&hashcash.Resource{
			Data:          "someone@gmail.com",
			ValidatorFunc: validateResource,
		},
		nil,
	)
	if err != nil {
		return solution, err
	}

	for {
		solution, err = hc.Compute()
		if err != nil {
			return solution, err
		}

		valid, err := hc.Verify(solution)
		if err != nil {
			return solution, err
		}

		if valid {
			break
		}
	}

	return solution, nil
}

// Verify verifies that PoW solution is valid.
func (g *generator) Verify(solution string) (bool, error) {
	valid, err := g.hc.Verify(solution)
	if err != nil {
		return valid, err
	}

	return valid, nil
}

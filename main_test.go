package understandingpointerreceivers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonPointerReceiversAllowsForAMoreFunctionalStyle(t *testing.T) {
	// Given
	integerContainer := IntegerContainer{10}
	runeContainer := RuneContainer{'a'}

	// When
	BoxOfThingContainers{}.
		AddThingContainer(&integerContainer).
		AddThingContainer(&runeContainer).
		RunThings()

	// Then
	assert.Equal(t, 2, integerContainer.thing)
	assert.Equal(t, 'z', runeContainer.thing)
}

func TestPointerReceiversLeadsToAMoreOOStyle(t *testing.T) {
	// Given
	integerContainer := IntegerContainer{10}
	runeContainer := RuneContainer{'a'}

	boxOfThingContainersUsingPointerReceivers := BoxOfThingContainersUsingPointerReceivers{}
	boxOfThingContainersUsingPointerReceivers.AddThingContainer(&integerContainer)
	boxOfThingContainersUsingPointerReceivers.AddThingContainer(&runeContainer)

	// When
	boxOfThingContainersUsingPointerReceivers.RunThings()

	// Then
	assert.Equal(t, 2, integerContainer.thing)
	assert.Equal(t, 'z', runeContainer.thing)
}

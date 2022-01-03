package chromosomes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBest(t *testing.T) {

	worst := Chromosome{
		Fitness: 3,
	}

	average := Chromosome{
		Fitness: 6,
	}

	best := Chromosome{
		Fitness: 9,
	}

	filter := &GetBest{}

	assert.True(t, filter.IsBetterFit(&average, &worst))
	assert.False(t, filter.IsBetterFit(&worst, &average))

	assert.True(t, filter.IsBetterFit(&best, &worst))
	assert.False(t, filter.IsBetterFit(&worst, &best))

	assert.True(t, filter.IsBetterFit(&best, &average))
	assert.False(t, filter.IsBetterFit(&average, &best))

}

func TestGetOldest(t *testing.T) {
	youngest := Chromosome{
		Age: 3,
	}

	average := Chromosome{
		Age: 6,
	}

	oldest := Chromosome{
		Age: 9,
	}

	filter := &GetOldest{}

	assert.True(t, filter.IsBetterFit(&average, &youngest))
	assert.False(t, filter.IsBetterFit(&youngest, &average))

	assert.True(t, filter.IsBetterFit(&oldest, &youngest))
	assert.False(t, filter.IsBetterFit(&youngest, &oldest))

	assert.True(t, filter.IsBetterFit(&oldest, &average))
	assert.False(t, filter.IsBetterFit(&average, &oldest))
}

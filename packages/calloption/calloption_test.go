package calloption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPeople(t *testing.T) {
	p1 := NewPeople("123456", "lc", WithAge(19))
	p2 := People{
		Id:   "123456",
		Name: "lc",
		Opt: peopleOption{
			Sex:     "female",
			Age:     19,
			Height:  170,
			Weight:  100,
			Color:   "yellow",
			Country: "China",
			Address: "HeNan",
		},
	}
	t.Logf("p1:%v, p2:%v", p1, p2)
	assert.Equal(t, *p1, p2)
}

// +build integration

package thing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoThisThingIT(t *testing.T) {

	thingtodo := "Cool thing inner integration test"
	ctd := CoolThing{thingtodo}
	r := DoThisThing(ctd)
	assert.Equal(t, thingtodo+" just got did", r)

}

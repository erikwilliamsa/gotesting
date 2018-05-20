package thing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockThing struct {
	mock.Mock
}

func (m *mockThing) ThingToDo() string {
	args := m.Called()
	return args.String(0)
}
func TestDoThisThing(t *testing.T) {

	mock := new(mockThing)
	thingtodo := "Make tacos"
	mock.On("ThingToDo").Return(thingtodo)
	r := DoThisThing(mock)

	assert.Equal(t, r, "Make tacos just got did")

}

func TestDoThisThingIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping itegration test")
	}
	thingtodo := "Cool thing inner integration test"
	ctd := CoolThing{thingtodo}
	r := DoThisThing(ctd)
	assert.Equal(t, thingtodo+" just got did", r)

}

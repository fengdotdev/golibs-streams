package streamcontroller_test

import (
	"testing"

	"github.com/fengdotdev/golibs-streams/obj/streamcontroller"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestStreamController_basis(t *testing.T) {

	controller := streamcontroller.NewStreamController[int]()

	controller.Add(1)
	controller.Add(2)
	controller.Add(3)

	controller.Stream().Listen(func(item int) {
		assert.True(t, item == 1 || item == 2 || item == 3)
	})

	//assert.Equal(t, 1, 1)

}

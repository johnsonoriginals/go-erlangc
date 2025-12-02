package erlangc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPressure(t *testing.T) {
	// Arrange.
	var rate int32 = 1
	var fulfillment int32 = 60
	var interval int32 = 60

	// Act.
	var result int32 = Pressure(rate, fulfillment, interval)

	// Assert.
	assert.Equal(t, int32(1), result, "The test expects that result will be 1.")
}

func TestUtilization(t *testing.T) {
	// Arrange.
	var pressure int32 = Pressure(60, 60, 60)
	var handlers int32 = 2

	// Act.
	var result int32 = Utilization(pressure, handlers)

	// Assert.
	assert.Equal(t, int32(30), result, "The test expects that result will be 30.")
}

func TestBNotBlocking(t *testing.T) {
	// Arrange.
	var pressure int32 = Pressure(2, 29)
	var handlers int32 = 1

	// Act.
	var result int32 = B(handlers, pressure)

	// Assert.
	assert.Equal(t, int32(0), result, "The test expects that result will be 0.")
}

func TestBIsBlocking(t *testing.T) {
	// Arrange.
	var pressure int32 = Pressure(60, 30)
	var handlers int32 = 1

	// Act.
	var result int32 = B(handlers, pressure)

	// Assert.
	assert.Equal(t, int32(1), result, "The test expects that result will be 1.")
}

func TestCNotQueuing(t *testing.T) {
	// Arrange.
	var handlers int32 = 10
	var pressure int32 = 5

	// Act.
	var result int32 = C(handlers, pressure)

	// Assert.
	assert.Equal(t, int32(0), result, "The test expects that the result will be 0.")
}

func TestCIsQueuing(t *testing.T) {
	// Arrange.
	var handlers int32 = 10
	var pressure int32 = 100

	// Act.
	var result int32 = C(handlers, pressure)

	// Assert.
	assert.Equal(t, int32(1), result, "The test expects that result will be 1.")
}

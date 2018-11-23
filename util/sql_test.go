package util_test

import (
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/spring2go/gravitee/util"
	"github.com/stretchr/testify/assert"
)

func TestInOrNull(t *testing.T) {
	nullInt := util.PositiveIntOrNull(1)
	assert.True(t, nullInt.Valid)

	value, err := nullInt.Value()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), value)
}

func TestPositiveIntOrNull(t *testing.T) {
	var (
		nullInt sql.NullInt64
		value   driver.Value
		err     error
	)

	// When the number is negative
	nullInt = util.PositiveIntOrNull(-1)

	// nullInt.Valid should be false
	assert.False(t, nullInt.Valid)

	// nullInt.Value() should return nil
	value, err = nullInt.Value()
	assert.Nil(t, err)
	assert.Nil(t, value)

	// When the number is greater than zero
	nullInt = util.PositiveIntOrNull(1)

	// nullInt.Valid should be true
	assert.True(t, nullInt.Valid)

	// nullInt.Value() should return the integer
	value, err = nullInt.Value()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), value)
}

func TestFloatOrNull(t *testing.T) {
	nullFloat := util.FloatOrNull(1.5)
	assert.True(t, nullFloat.Valid)

	value, err := nullFloat.Value()
	assert.Nil(t, err)
	assert.Equal(t, 1.5, value)
}

func TestPositiveFloatOrNull(t *testing.T) {
	var (
		nullFloat sql.NullFloat64
		value     driver.Value
		err       error
	)

	// When the number is negative
	nullFloat = util.PositiveFloatOrNull(-0.5)

	// nullFloat.Valid should be false
	assert.False(t, nullFloat.Valid)

	// nullFloat.Value() should return nil
	value, err = nullFloat.Value()
	assert.Nil(t, err)
	assert.Nil(t, value)

	// When the number is greater than zero
	nullFloat = util.PositiveFloatOrNull(1.5)

	// nullFloat.Valid should be true
	assert.True(t, nullFloat.Valid)

	// nullFloat.Value() should return the integer
	value, err = nullFloat.Value()
	assert.Nil(t, err)
	assert.Equal(t, 1.5, value)
}

func TestStringOrNull(t *testing.T) {
	var (
		nullString sql.NullString
		value      driver.Value
		err        error
	)

	// When the string is empty
	nullString = util.StringOrNull("")

	// nullString.Valid should be false
	assert.False(t, nullString.Valid)

	// nullString.Value() should return nil
	value, err = nullString.Value()
	assert.Nil(t, err)
	assert.Nil(t, value)

	// When the string is not empty
	nullString = util.StringOrNull("foo")

	// nullString.Valid should be true
	assert.True(t, nullString.Valid)

	// nullString.Value() should return the string
	value, err = nullString.Value()
	assert.Nil(t, err)
	assert.Equal(t, "foo", value)
}

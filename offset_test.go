package vmap

import (
	"testing"

	"github.com/rs/vast"
	"github.com/stretchr/testify/assert"
)

func TestOffsetMarshaler(t *testing.T) {
	b, err := Offset{}.MarshalText()
	if assert.NoError(t, err) {
		assert.Equal(t, "0%", string(b))
	}
	b, err = Offset{Percent: .1}.MarshalText()
	if assert.NoError(t, err) {
		assert.Equal(t, "10%", string(b))
	}
	d := vast.Duration(0)
	b, err = Offset{Duration: &d}.MarshalText()
	if assert.NoError(t, err) {
		assert.Equal(t, "00:00:00", string(b))
	}
	b, err = Offset{Position: 1}.MarshalText()
	if assert.NoError(t, err) {
		assert.Equal(t, "#1", string(b))
	}
	b, err = Offset{Position: OffsetStart}.MarshalText()
	if assert.NoError(t, err) {
		assert.Equal(t, "start", string(b))
	}
	b, err = Offset{Position: OffsetEnd}.MarshalText()
	if assert.NoError(t, err) {
		assert.Equal(t, "end", string(b))
	}
	b, err = Offset{Position: -4}.MarshalText()
	assert.EqualError(t, err, "Invalid position: -4")
}

func TestOffsetUnmarshaler(t *testing.T) {
	var o Offset
	if assert.NoError(t, o.UnmarshalText([]byte("0%"))) {
		assert.Nil(t, o.Duration)
		assert.Equal(t, float32(0.0), o.Percent)
	}
	o = Offset{}
	if assert.NoError(t, o.UnmarshalText([]byte("10%"))) {
		assert.Nil(t, o.Duration)
		assert.Equal(t, float32(0.1), o.Percent)
	}
	o = Offset{}
	if assert.NoError(t, o.UnmarshalText([]byte("00:00:00"))) {
		if assert.NotNil(t, o.Duration) {
			assert.Equal(t, vast.Duration(0), *o.Duration)
		}
		assert.Equal(t, float32(0), o.Percent)
	}
	o = Offset{}
	if assert.NoError(t, o.UnmarshalText([]byte("#1"))) {
		assert.Nil(t, o.Duration)
		assert.Equal(t, float32(0), o.Percent)
		assert.Equal(t, 1, o.Position)
	}
	o = Offset{}
	if assert.NoError(t, o.UnmarshalText([]byte("start"))) {
		assert.Nil(t, o.Duration)
		assert.Equal(t, float32(0), o.Percent)
		assert.Equal(t, OffsetStart, o.Position)
	}
	o = Offset{}
	if assert.NoError(t, o.UnmarshalText([]byte("end"))) {
		assert.Nil(t, o.Duration)
		assert.Equal(t, float32(0), o.Percent)
		assert.Equal(t, OffsetEnd, o.Position)
	}
	o = Offset{}
	assert.EqualError(t, o.UnmarshalText([]byte("#0")), "invalid offset: #0")
	o = Offset{}
	assert.EqualError(t, o.UnmarshalText([]byte("#-10")), "invalid offset: #-10")
	o = Offset{}
	assert.EqualError(t, o.UnmarshalText([]byte("#1.1")), "invalid offset: #1.1")
	o = Offset{}
	assert.EqualError(t, o.UnmarshalText([]byte("abc%")), "invalid offset: abc%")
}

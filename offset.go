package vmap

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/vast"
)

// Offset represents either a vast.Duration or a percentage of the video duration.
type Offset struct {
	// If not nil, the Offset is duration based
	Duration *vast.Duration
	// If not 0 and Duration is nil, the Offset is a position number.
	// Some numbers are reserved like -1 (vmap.OffsetStart) or -2 (vmap.OffsetEnd).
	Position int
	// If Duration is nil and Position 0, the Offset is percent based
	Percent float32
}

const (
	// OffsetStart defines an Offset.Position for pre-rolls
	OffsetStart = -1
	// OffsetEnd defines an Offset.Position for post-rolls
	OffsetEnd = -2
)

// MarshalText implements the encoding.TextMarshaler interface.
func (o Offset) MarshalText() ([]byte, error) {
	if o.Duration != nil {
		return o.Duration.MarshalText()
	} else if o.Position != 0 {
		switch o.Position {
		case OffsetStart:
			return []byte("start"), nil
		case OffsetEnd:
			return []byte("end"), nil
		default:
			if o.Position < 0 {
				return nil, fmt.Errorf("Invalid position: %d", o.Position)
			}
			return []byte(fmt.Sprintf("#%d", o.Position)), nil
		}
	}
	return []byte(fmt.Sprintf("%d%%", int(o.Percent*100))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (o *Offset) UnmarshalText(data []byte) error {
	switch string(data) {
	case "start":
		o.Position = OffsetStart
		return nil
	case "end":
		o.Position = OffsetEnd
		return nil
	}
	if strings.HasSuffix(string(data), "%") {
		p, err := strconv.ParseInt(string(data[:len(data)-1]), 10, 8)
		if err != nil {
			return fmt.Errorf("invalid offset: %s", data)
		}
		o.Percent = float32(p) / 100
		return nil
	}
	if strings.HasPrefix(string(data), "#") {
		p, err := strconv.ParseInt(string(data[1:]), 10, 32)
		if err != nil || p <= 0 {
			return fmt.Errorf("invalid offset: %s", data)
		}
		o.Position = int(p)
		return nil
	}
	var d vast.Duration
	o.Duration = &d
	return o.Duration.UnmarshalText(data)
}

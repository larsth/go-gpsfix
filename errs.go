package gpsfix

import "errors"

var (
	//ErrUnknownFixMode is the error for an unknown FixMode value
	ErrUnknownFixMode = errors.New("Unknown FixMode value: valid values " +
		"are FixNotSeen, 0 or FixNone, 1 or Fix2D, 2 or Fix3D, 3.")

	//ErrNilSlice is the error for a nil input byte slice to FixMode's
	//UnmarshalJSON method
	ErrNilSlice = errors.New("Nil input slice to 'func (f *FixMode) " +
		"UnmarshalJSON(p []byte) error'.")
)

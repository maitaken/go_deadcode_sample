package os

import (
	"github.com/maitaken/go_deadcode_sample/internal/enum"
)

type OSInterface interface {
	Type() enum.OSType
}

type Windows struct {
}

func (w Windows) Type() enum.OSType {
	return enum.Windows
}

// unused
type Mac struct {
}

func (w Mac) Type() enum.OSType {
	return enum.Windows
}

// unused (from deadcode func)
type Android struct {
}

func (w Android) Type() enum.OSType {
	return enum.Android
}

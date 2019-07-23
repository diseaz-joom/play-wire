package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyPrinter(t *testing.T) {
	p := NewPrinter("Empty")
	require.Equal(t, "Empty{}", p.String())
}

func TestPrinterWithChildren(t *testing.T) {
	p := NewPrinter("Root",
		NewPrinter("First"),
		NewPrinter("Second"),
	)
	require.Equal(t, "Root{First{}, Second{}}", p.String())
}

func TestPrinterHierarchy(t *testing.T) {
	p := NewPrinter("Root",
		NewPrinter("First",
			NewPrinter("Second"),
		),
	)
	require.Equal(t, "Root{First{Second{}}}", p.String())
}

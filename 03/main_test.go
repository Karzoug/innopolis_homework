package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_chainFormatter_Format(t *testing.T) {
	chf := chainFormatter{}
	chf.AddFormatter(bold{})
	assert.Equal(t, "**Hello, World!**", chf.Format("Hello, World!"))

	chf.AddFormatter(plainText{})
	chf.AddFormatter(italic{})
	assert.Equal(t, "_**Hello, World!**_", chf.Format("Hello, World!"))
}

func Test_plainText_Format(t *testing.T) {
	var fmt Formatter = plainText{}
	assert.Equal(t, "Hello, World!", fmt.Format("Hello, World!"))
	assert.Equal(t, "**Hello, World!**", fmt.Format("**Hello, World!**"))
}

func Test_bold_Format(t *testing.T) {
	var fmt Formatter = bold{}
	assert.Equal(t, "**Hello, World!**", fmt.Format("Hello, World!"))
	assert.Equal(t, "**_Hello, World!_**", fmt.Format("_Hello, World!_"))
}

func Test_italic_Format(t *testing.T) {
	var fmt Formatter = italic{}
	assert.Equal(t, "_Hello, World!_", fmt.Format("Hello, World!"))
	assert.Equal(t, "_**Hello, World!**_", fmt.Format("**Hello, World!**"))
}

func Test_code_Format(t *testing.T) {
	var fmt Formatter = code{}
	assert.Equal(t, "`Hello, World!`", fmt.Format("Hello, World!"))
	assert.Equal(t, "`**Hello, World!**`", fmt.Format("**Hello, World!**"))
}

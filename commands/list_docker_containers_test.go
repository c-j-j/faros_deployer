package commands_test

import (
	"github.com/c-j-j/faros_deployer/commands"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemp(t *testing.T) {
	assert.Equal(t, 0, commands.ListDockerContainers{}.Execute())
}

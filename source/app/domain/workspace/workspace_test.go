package workspace_test

import (
	"testing"
	"yoyaku/app/core"
	"yoyaku/app/domain/workspace"

	"github.com/stretchr/testify/assert"
)

func TestNewEntityOk(t *testing.T) {
	var name string = "test.workspace"

	entity, err := workspace.NewEntity(name)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, 0, entity.GetId())
	assert.Equal(t, name, entity.GetName())
}

func TestNewEntityNgValidationError(t *testing.T) {
	var name string = "test.test.test.test.test.test.test.test.test.test.workspace"

	entity, err := workspace.NewEntity(name)
	if entity != nil {
		t.Errorf("Error")
	}

	dstErr := core.AsAppError(err)
	assert.Equal(t, core.ValidationError, dstErr.Code())
}

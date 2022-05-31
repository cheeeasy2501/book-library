package relationships

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelations_MarshalJSON(t *testing.T) {
	r := Relations{"author", "test1", "test2"}

	json, err := r.MarshalJSON()

	assert.NoError(t, err)
	assert.NotNil(t, json)
	assert.Equal(t, "author,test1,test2", string(json))
}

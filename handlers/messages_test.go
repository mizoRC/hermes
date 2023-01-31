package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getSpaceMessages(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}

	assert := assert.New(t)

	// assert equality
	assert.Equal(123, 123, "they should be equal")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getSpaceMessages(tt.args.c)
		})
	}
}

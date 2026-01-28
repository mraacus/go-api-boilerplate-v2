package handler

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name   string
		param  NewHandlerParam
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "success",
			param:  NewHandlerParam{},
			wantErr: assert.NoError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler, err := New(test.param)
			if test.wantErr(t, err) {
				return
			}
			assert.NotNil(t, handler)
		})
	}
}

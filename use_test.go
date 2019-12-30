package doer

import (
	"testing"

	"github.com/nthlongtma/go-gomock/mocks"

	"github.com/golang/mock/gomock"
)

func TestDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	user := User{
		mockDoer,
	}

	// mockDoer.EXPECT().Do(2, "one").Return(nil)
	user.Use()

}

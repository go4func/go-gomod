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

	mockDoer.EXPECT().Do(1, "two").Return(nil)
	user.Use()

}

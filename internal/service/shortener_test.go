package service

import (
	"context"
	"errors"
	"go.uber.org/mock/gomock"
	"linkSh/internal/service/mocks"
	"testing"
)

func TestFindLongURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStorage := mocks.NewMockShortenerService(ctrl)

	mockStorage.EXPECT().LongLink(gomock.Any(), "http://Y12.ru").Return("", errors.New("not find long link"))

	result, err := mockStorage.LongLink(context.Background(), "http://Y12.ru")
	if err == nil || err.Error() != "not find long link" {
		t.Errorf("Expected error: 'not find long link', got: %v", err)
	}
	if result != "" {
		t.Errorf("Expected result: '', got: %v", result)
	}
}

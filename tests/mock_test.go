package tests

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"hospital-rest/mocks"
	"testing"
)

var RESULT = "I got mocked"

func TestRadarWithGoMock(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	radar_mock := mocks.NewMockRadar(controller)

	radar_mock.EXPECT().Scan(1).Return(RESULT)

	result := radar_mock.Scan(1)

	assert.Equal(t, result, RESULT)
}

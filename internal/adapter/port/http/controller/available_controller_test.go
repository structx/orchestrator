package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/structx/orchestrator/internal/adapter/port/http/controller"
	"github.com/structx/orchestrator/internal/core/domain"
	"github.com/structx/orchestrator/internal/core/domain/mocks"
)

type AvailableControllerSuite struct {
	suite.Suite
	cc *controller.Available
}

func (suite *AvailableControllerSuite) SetupTest() {

	mockAvail := mocks.NewAvailableService(suite.T())
	mockAvail.EXPECT().ListRegions(mock.Anything).Return([]domain.Region{}, nil).Once()
	suite.cc = controller.NewAvailable(mockAvail)
}

func (suite *AvailableControllerSuite) TestGetRegions() {

	assert := suite.Assert()

	testcases := []struct {
		expected int
	}{
		{
			expected: http.StatusAccepted,
		},
	}

	for _, tt := range testcases {

		rr := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/api/v1/avail/regions", nil)
		assert.NoError(err)

		suite.cc.ListRegions(rr, request)
		assert.Equal(tt.expected, rr.Code)
	}
}

func TestAvailableControllerSuite(t *testing.T) {
	suite.Run(t, new(AvailableControllerSuite))
}

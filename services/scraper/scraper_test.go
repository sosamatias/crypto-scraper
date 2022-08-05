package scraper

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sosamatias/crypto-scraper/gateway"
	"github.com/sosamatias/crypto-scraper/mocks"
	"github.com/sosamatias/crypto-scraper/model"
	"github.com/stretchr/testify/assert"
)

func Test_Execute_OK(t *testing.T) {
	t.Parallel()

	gatewayResponse := model.GatewayResponseMock()
	cryptoSnapshots := gatewayResponse.ToCryptoSnapshot()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mocks.NewMockGateway(ctrl)
	mockGateway.EXPECT().List(gateway.SortMarketCap).Return(&gatewayResponse, nil).Times(1)

	mockRepository := mocks.NewMockRepository(ctrl)
	mockRepository.EXPECT().CreateInBatches(cryptoSnapshots, 50).Return(nil).Times(1)

	err := Execute(mockGateway, mockRepository)
	assert.NoError(t, err)
}

func Test_Execute_Error(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockErr := errors.New("mock error")
	mockGateway := mocks.NewMockGateway(ctrl)
	mockGateway.EXPECT().List(gateway.SortMarketCap).Return(nil, mockErr).Times(1)

	mockRepository := mocks.NewMockRepository(ctrl)

	err := Execute(mockGateway, mockRepository)
	assert.ErrorIs(t, err, mockErr)
}

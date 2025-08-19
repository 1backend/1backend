package endpoint_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTokenExchangeCaching(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClientFactory := client.NewMockClientFactory(ctrl)

	mockUserSvc := test.MockUserSvc(context.Background(), ctrl)

	exchanged := []string{"first-exchange", "second-exchange"}
	count := 0

	mockClientFactory.EXPECT().
		Client(gomock.Any()).
		Return(&openapi.APIClient{
			UserSvcAPI: mockUserSvc,
		}).Times(2)

	mockExchangeTokenRequest := openapi.ApiExchangeTokenRequest{
		ApiService: mockUserSvc,
	}

	mockUserSvc.EXPECT().
		ExchangeToken(gomock.Any()).
		Return(mockExchangeTokenRequest).
		Times(2)

	mockUserSvc.EXPECT().
		ExchangeTokenExecute(gomock.Any()).
		DoAndReturn(func(r openapi.ApiExchangeTokenRequest) (*openapi.UserSvcExchangeTokenResponse, *http.Response, error) {
			count++
			return &openapi.UserSvcExchangeTokenResponse{
				Token: openapi.UserSvcAuthToken{
					Token:     exchanged[count-1],
					ExpiresAt: time.Now().Add(1 * time.Second).Format(time.RFC3339),
				},
			}, nil, nil
		}).Times(2)

	// Use a very short cache duration to test expiry
	pc := endpoint.NewTokenExchanger(mockClientFactory)
	pc.(*endpoint.TokenExchangerImpl).Testing = true

	// First call - should hit service
	newToken, err1 := pc.ExchangeToken(context.TODO(), "original-token", "new-app")
	assert.NoError(t, err1)
	assert.Equal(t, exchanged[0], newToken)

	// Second call - should hit cache, not service
	_, err1 = pc.ExchangeToken(context.TODO(), "original-token", "new-app")
	assert.NoError(t, err1)
	assert.Equal(t, exchanged[0], newToken)

	// Wait for cache to expire
	time.Sleep(time.Second)

	// Third call - should hit service again
	resp3, err3 := pc.ExchangeToken(context.TODO(), "original-token", "new-app")
	assert.NoError(t, err3)
	assert.Equal(t, exchanged[1], resp3)
}

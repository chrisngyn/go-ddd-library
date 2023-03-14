package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/common/client/lending"
)

type LendingHTTPClient struct {
	client *lending.ClientWithResponses
}

func NewLendingHTTPClient(t *testing.T, address, parentRoute string) LendingHTTPClient {
	ok := WaitForPort(address)
	require.True(t, ok, "Lending HTTP Server timeout")

	url := fmt.Sprintf("http://%s/%s", address, strings.TrimLeft(parentRoute, "/"))

	client, err := lending.NewClientWithResponses(url)
	require.NoError(t, err)

	return LendingHTTPClient{client: client}
}

func (c LendingHTTPClient) PlaceOnHold(t *testing.T, patronID, bookID string, numOfDays int) int {
	response, err := c.client.PlaceHold(context.Background(), patronID, lending.PlaceHoldJSONRequestBody{
		BookId:    bookID,
		NumOfDays: numOfDays,
	})
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())

	return response.StatusCode
}

package api

import (
	"context"
	"ecommerce/common/response"
	"net/http"

	"github.com/cardinalby/hureg"
	"github.com/danielgtaylor/huma/v2"
)

type API_Test struct{}

func TestAPI(api hureg.APIGen) {
	handler := &API_Test{}
	apiGroup := api.AddBasePath("/v1")

	hureg.Register(apiGroup, huma.Operation{
		Tags:        []string{"TestTag"},
		OperationID: "TestInfo",
		Method:      http.MethodGet,
		Path:        "/test-api",
	}, handler.GetTestInfo)
}

func (h *API_Test) GetTestInfo(ctx context.Context, req *struct {
}) (*response.TestMessageResponse, error) {
	result := &response.TestMessageResponse{}

	result.Body.Message1 = "Hế lô mấy cưng"
	result.Body.Message2 = "Cuộc đời này ngắn ngủi, sống sao cho người ta quý mến!"
	result.Body.Message3 = "Sống mà chó má quá kiểu gì trời cũng trả báu không sớm thì muộn"
	return result, nil
}

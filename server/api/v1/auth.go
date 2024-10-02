package api

import (
	"context"
	"ecommerce/common/model"
	"ecommerce/common/response"
	"net/http"

	"github.com/cardinalby/hureg"
	"github.com/danielgtaylor/huma/v2"
)

type ApiAuth struct{}

func RegisterApiAuth(api hureg.APIGen) {
	handler := &ApiAuth{}

	apiGroup := api.AddBasePath("/v1/auth")

	hureg.Register(apiGroup, huma.Operation{
		Tags:        []string{"Authentication"},
		OperationID: "Sign In",
		Method:      http.MethodPost,
		Path:        "/sign-in",
	}, handler.SignIn)

	hureg.Register(apiGroup, huma.Operation{
		Tags:        []string{"Authentication"},
		OperationID: "Sign In",
		Method:      http.MethodPost,
		Path:        "/sign-in",
	}, handler.SignUp)

	hureg.Register(apiGroup, huma.Operation{
		Tags:        []string{"Authentication"},
		OperationID: "Sign In",
		Method:      http.MethodPost,
		Path:        "/sign-in",
	}, handler.SignOut)
}

func (h *ApiAuth) SignIn(c context.Context, req *struct {
	Form model.LoginForm
}) (res *response.GenericResponse[model.LoginResponse], err error) {

	return
}

func (h *ApiAuth) SignUp(c context.Context, req *struct {
	Form model.LoginForm
}) (res *response.GenericResponse[model.LoginResponse], err error) {

	return
}

func (h *ApiAuth) SignOut(c context.Context, req *struct {
	Form model.LoginForm
}) (res *response.GenericResponse[model.LoginResponse], err error) {

	return
}

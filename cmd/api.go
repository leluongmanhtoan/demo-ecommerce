package cmd

import (
	"ecommerce/common/constant"

	"ecommerce/common/log"
	"ecommerce/server"

	APIv1 "ecommerce/server/api/v1"

	"github.com/cardinalby/hureg"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/spf13/cobra"
)

var cmdAPI = &cobra.Command{
	Use:   "api",
	Short: "Ecommerce API server",
	Long:  "Ecommerce API server",
	Run: func(cmd *cobra.Command, args []string) {
		RunCommandAPI()
	},
}

func RunCommandAPI() {
	log.Debug("Starting service...")

	// Create HTTP server
	server := server.NewServer()
	humaAPI := humagin.New(server.Engine, huma.Config{
		OpenAPI: &huma.OpenAPI{
			OpenAPI: "3.1.0",
			Info: &huma.Info{
				Title:   "DEMO - ECOMMERCE PROJECT API DOCUMENT",
				Version: constant.SERVICE_VERSION,
			},
			Servers: []*huma.Server{
				{
					URL:         "https://api-ecommerce.codelo.life/",
					Description: "AWS Environment",
					Variables:   map[string]*huma.ServerVariable{},
				},
				{
					URL:         "http://localhost:8686/",
					Description: "Local Environment",
					Variables:   map[string]*huma.ServerVariable{},
				},
			},
		},
		OpenAPIPath:   "/openapi",
		DocsPath:      "/docs",
		Formats:       huma.DefaultFormats,
		DefaultFormat: "application/json",
	})

	// huma
	api := hureg.NewAPIGen(humaAPI)

	api.GetHumaAPI()
	APIv1.TestAPI(api)
	log.Debug("Starting api server")
	server.Start(config.Port)
}

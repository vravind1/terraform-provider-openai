package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func New() provider.Provider {
	return &openaiProvider{}
}

var _ provider.Provider = (*openaiProvider)(nil)

type openaiProvider struct{}

func (p *openaiProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "openai"
}

func (p *openaiProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {
}

func (p *openaiProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *openaiProvider) Resources(context.Context) []func() resource.Resource {
	return nil
}

func (p *openaiProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCompletionDataSource,
	}
}

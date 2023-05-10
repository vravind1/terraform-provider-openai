package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"os"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openai "github.com/sashabaranov/go-openai"
)

type completionDataSource struct{}

type completionDataSourceModel struct {
	Prompt       types.String `tfsdk:"prompt"`
	Model        types.String `tfsdk:"model"`
	Result       types.String `tfsdk:"result"`
	FinishReason types.String `tfsdk:"finish_reason"`
	MaxTokens    types.Int64  `tfsdk:"max_tokens"`
	TotalTokens  types.Int64  `tfsdk:"total_tokens"`
	ID           types.String `tfsdk:"id"`
}

var _ datasource.DataSource = (*completionDataSource)(nil)

func NewCompletionDataSource() datasource.DataSource {
	return &completionDataSource{}
}

func (d *completionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_completion"
}

func (d *completionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `This is Completion Data Source`,
		Attributes: map[string]schema.Attribute{

			"prompt": schema.StringAttribute{
				Description: "Prompt text for Completion",
				Required:    true,
			},
			"model": schema.StringAttribute{
				Description: "Open AI model for this request",
				Required:    true,
			},
			"result": schema.StringAttribute{
				Description: "Responded Completion Text",
				Computed:    true,
			},
			"max_tokens": schema.Int64Attribute{
				Description: "The maximum number of tokens to generate in the completion",
				Required:    true,
			},
			"total_tokens": schema.Int64Attribute{
				Description: "Total tokens used for this request",
				Computed:    true,
			},
			"finish_reason": schema.StringAttribute{
				Description: "Reason of completion this request",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "ID of this request",
				Computed:    true,
			},
		},
	}
}

func (e *completionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data completionDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	c := openai.NewClient(os.Getenv("OPENAI_APIKEY"))

	query := openai.CompletionRequest{
		Model:     data.Model.ValueString(),
		MaxTokens: int(data.MaxTokens.ValueInt64()),
		Prompt:    data.Prompt.ValueString(),
	}

	res, err := c.CreateCompletion(ctx, query)
	if err != nil {

		tflog.Trace(ctx, fmt.Sprintf("Completion Error %s", err))
		resp.Diagnostics.AddError("Completion Error Summary", err.Error())
		return
	}
	tflog.Trace(ctx, res.Choices[0].FinishReason)

	data.Result = types.StringValue(res.Choices[0].Text)
	data.FinishReason = types.StringValue(res.Choices[0].FinishReason)
	data.TotalTokens = types.Int64Value(int64(res.Usage.TotalTokens))
	data.ID = types.StringValue(res.ID)
	diags := resp.State.Set(ctx, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(diags...)
}

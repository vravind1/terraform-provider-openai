# Using Terraform Cloud's Continuous Validation feature

-> This can be used with Terraform 1.5 and later versions

The [Continuous Validation](https://developer.hashicorp.com/terraform/cloud-docs/workspaces/health#continuous-validation) feature in [Terraform Cloud (TFC)](https://developer.hashicorp.com/terraform/cloud-docs) allows users to make assertions about their infrastructure between applied runs. This helps users to identify issues at the time they first appear and avoid situations where a change is only identified during a future terraform plan/apply or once it causes a user-facing problem.

Users can add checks to their Terraform configuration using an HCL language feature called check{} blocks. Check blocks contain assertions that are defined with a custom condition expression and an error message. When the condition expression evaluates to true the check passes, but when the expression evaluates to false Terraform will show a warning message that includes the user-defined error message.

## Example - To check the reason of completion
The Completion API provides a finish_reason field in its response, allowing users to verify if the API has returned a complete message successfully. This example demonstrates how to utilize the finish_reason field to ensure the completeness of the API response.
```hcl
terraform {

  required_providers {
    openai = {
      source  = "registry.terraform.io/vravind1/openai"
      version = "0.1.0"
    }
  }
}

check "finish_status" {
  data "openai_completion" "example" {
    model      = "text-davinci-003"
    max_tokens = 30
    prompt     = "write me a essay about terraform"
  }
  assert {
    condition     = data.openai_completion.example.finish_reason == "stop"
    error_message = format("Completion ended abnormally due to %s reason", data.openai_completion.example.finish_reason)
  }
}
```
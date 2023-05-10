---
page_title: "OpenAI Provider"
subcategory: ""
description: |-
  
---

# OpenAI Provider

The OpenAI Terraform provider enables users to interact with OpenAI APIs and perform various actions through Terraform configuration files.

Use the navigation to the left to read about the available data sources.

## Example Usage

```terraform
data "openai_completion" "example" {
  model      = "text-davinci-003"
  max_tokens = 16
  prompt     = "which is the best cloud provider?"
}

output "example_response" {
  value = data.openai_completion.example.result
}
```
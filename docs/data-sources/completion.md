---
page_title: "openai_completion Data Source - terraform-provider-openai"
subcategory: ""
description: |-
  Completion API
---

# openai_completion (Data Source)

~> **NOTE:** To use the OpenAI provider, set your OpenAI API key as the`OPENAI_APIKEY`environment variable.
## Example Usage

### URL Usage
```terraform
data "openai_completion" "example" {
  model      = "text-davinci-003"
  max_tokens = 16
  prompt     = "which is the best cloud provider?"
}

output "example_response" {
  value = data.openai_completion.example.example_result
}

output "example_tokens" {
  value = data.openai_completion.example.example_total_tokens
}

output "example_reason" {
  value = data.openai_completion.example.example_finish_reason
}
```

## Schema

### Required

- `model` (String) ID of the model to use
- `max_tokens` (Number) The maximum number of tokens to generate in the completion
- `prompt` (String) Prompt to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.

### Read-Only

- `result` (String) The completion text.
- `total_tokens` (Number) The total number of tokens used for this completion request.
- `finish_reason` (String) reason to complete this request
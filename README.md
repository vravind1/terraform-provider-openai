# Terraform Provider: OpenAI
[The Terraform Provider for OpenAI](https://registry.terraform.io/providers/vravind1/openai/latest) is a plugin built using the [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework). 

## Requirements
- [Terraform](https://www.terraform.io/) 1.0+
- API key from [OpenAI](https://platform.openai.com/signup)

## Usage


**NOTE:** To use the OpenAI provider, set your OpenAI API key as the `OPENAI_APIKEY` environment variable.


```terraform
terraform {
required_providers {
    openai = {
      source  = "registry.terraform.io/vravind1/openai"    
      version = "0.1.0".  // use the appropriate version
    }
  }
}

data "openai_completion" "example" {
  model      = "text-davinci-003"
  max_tokens = 16
  prompt     = "which is the best cloud provider?"
}

output "example_response" {
  value = data.openai_completion.example.result
}

output "example_tokens" {
  value = data.openai_completion.example.total_tokens
}

output "example_reason" {
  value = data.openai_completion.example.finish_reason
}
```
More details on how to use this provider can be found on the [Terraform Registry](https://registry.terraform.io/providers/vravind1/openai/latest). 

## License

[MIT License](https://github.com/vravind1/terraform-provider-openai/blob/main/LICENSE)

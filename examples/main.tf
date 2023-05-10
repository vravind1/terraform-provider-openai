terraform {

  required_providers {
    openai = {
      source  = "registry.terraform.io/vravind1/openai"
      version = "0.0.1"
    }
  }
}

data "openai_completion" "example" {
  model      = "text-davinci-003"
  max_tokens = 30
  prompt     = "say test"
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

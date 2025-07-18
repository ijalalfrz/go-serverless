provider "aws" {
  region = "ap-southeast-1"

  endpoints {
    dynamodb = "http://dynamodb:8000"
  }

  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}

module "dynamodb" {
  source = "../../modules/dynamodb"

  environment = "local"
  table_name  = var.table_name
  tags = {
    Project = "go-serverless"
  }
}
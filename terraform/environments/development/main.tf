provider "aws" {
  region = var.aws_region
}

data "aws_dynamodb_table" "existing_table" {
  name = "${var.table_name}_${var.environment}"
}

data "aws_iam_role" "existing_role" {
  name = "${var.app_name}-${var.environment}-lambda-role"
}

module "dynamodb" {
  source = "../../modules/dynamodb"

  environment = var.environment
  table_name  = var.table_name
  tags = {
    Environment = var.environment
    Project     = var.app_name
  }
}

module "lambda" {
  source = "../../modules/lambda"

  app_name        = var.app_name
  environment     = var.environment
  lambda_zip_path = var.lambda_zip_path
  memory_size     = 128
  timeout         = 30

  environment_variables = {
    DYNAMODB_TABLE_NAME = module.dynamodb.table_name
    LOG_LEVEL           = "debug"
    DYNAMODB_REGION     = var.aws_region
    PROFILING_ENABLED   = "false"
  }

  dynamodb_table_arn = module.dynamodb.table_arn
}

module "api_gateway" {
  source = "../../modules/api_gateway"

  app_name                   = var.app_name
  environment                = var.environment
  lambda_function_name       = module.lambda.function_name
  lambda_function_invoke_arn = module.lambda.function_invoke_arn
}
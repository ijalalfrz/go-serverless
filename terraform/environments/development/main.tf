provider "aws" {
  region = var.aws_region
}

# Create ECR Repository
resource "aws_ecr_repository" "app" {
  name         = "${var.app_name}-${var.environment}"
  force_delete = true
}

# Create ECR Lifecycle Policy
resource "aws_ecr_lifecycle_policy" "app" {
  repository = aws_ecr_repository.app.name

  policy = jsonencode({
    rules = [{
      rulePriority = 1
      description  = "Keep last 5 images"
      selection = {
        tagStatus   = "any"
        countType   = "imageCountMoreThan"
        countNumber = 5
      }
      action = {
        type = "expire"
      }
    }]
  })
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

  app_name    = var.app_name
  environment = var.environment
  image_uri   = "${aws_ecr_repository.app.repository_url}:latest"
  memory_size = 128
  timeout     = 30

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
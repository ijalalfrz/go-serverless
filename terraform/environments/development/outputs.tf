output "api_endpoint" {
  value = module.api_gateway.api_endpoint
}

output "lambda_function_name" {
  value = module.lambda.function_name
}

output "dynamodb_table_name" {
  value = module.dynamodb.table_name
}
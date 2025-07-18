output "function_name" {
  value       = aws_lambda_function.function.function_name
  description = "Name of the Lambda function"
}

output "function_arn" {
  value       = aws_lambda_function.function.arn
  description = "ARN of the Lambda function"
}

output "function_invoke_arn" {
  value       = aws_lambda_function.function.invoke_arn
  description = "Invoke ARN of the Lambda function"
}

output "role_arn" {
  value       = aws_iam_role.lambda_role.arn
  description = "ARN of the Lambda IAM role"
}
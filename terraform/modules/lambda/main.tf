resource "aws_lambda_function" "function" {
  filename      = var.lambda_zip_path
  function_name = "${var.app_name}-${var.environment}"
  role          = aws_iam_role.lambda_role.arn
  handler       = "app"
  runtime       = "provided.al2023"
  memory_size   = var.memory_size
  timeout       = var.timeout

  environment {
    variables = var.environment_variables
  }
}

resource "random_string" "lambda_suffix" {
  length  = 8
  special = false
  upper   = false
}


resource "aws_iam_role" "lambda_role" {
  name = "${var.app_name}-${var.environment}-lambda-role-${random_string.lambda_suffix.result}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

  lifecycle {
    create_before_destroy = false
    prevent_destroy       = true
  }
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Add DynamoDB permissions
resource "aws_iam_role_policy" "dynamodb_policy" {
  name = "${var.app_name}-${var.environment}-dynamodb-policy-${random_string.lambda_suffix.result}"
  role = aws_iam_role.lambda_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:DeleteItem",
          "dynamodb:Query",
          "dynamodb:Scan"
        ]
        Resource = [var.dynamodb_table_arn]
      }
    ]
  })
}
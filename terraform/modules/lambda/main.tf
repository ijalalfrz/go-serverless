resource "aws_lambda_function" "function" {
  function_name = "${var.app_name}-${var.environment}"
  role         = aws_iam_role.lambda_role.arn
  package_type = "Image"
  image_uri    = var.image_uri
  memory_size  = var.memory_size
  timeout      = var.timeout

  environment {
    variables = var.environment_variables
  }
}

resource "aws_iam_role" "lambda_role" {
  name = "${var.app_name}-${var.environment}-lambda-role"

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
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Add DynamoDB permissions
resource "aws_iam_role_policy" "dynamodb_policy" {
  name = "${var.app_name}-${var.environment}-dynamodb-policy"
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
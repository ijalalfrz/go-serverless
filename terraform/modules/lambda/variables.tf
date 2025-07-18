variable "app_name" {
  type = string
}

variable "environment" {
  type = string
}

variable "memory_size" {
  type    = number
  default = 128
}

variable "timeout" {
  type    = number
  default = 30
}

variable "environment_variables" {
  type    = map(string)
  default = {}
}

variable "dynamodb_table_arn" {
  type = string
}

variable "lambda_zip_path" {
  type        = string
  description = "Path to the Lambda function zip file"
}
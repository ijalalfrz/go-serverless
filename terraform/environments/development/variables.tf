variable "aws_region" {
  type    = string
  default = "ap-southeast-1"
}

variable "environment" {
  type    = string
  default = "development"
}

variable "app_name" {
  type    = string
  default = "go-serverless"
}

variable "table_name" {
  type    = string
  default = "devices_rizal_alfarizi"
}

variable "lambda_zip_path" {
  type    = string
  default = "../../bin/app.zip"
}
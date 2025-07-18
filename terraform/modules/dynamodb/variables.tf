variable "environment" {
  type        = string
  description = "Environment name"
}

variable "table_name" {
  type        = string
  description = "Base table name"
}

variable "read_capacity" {
  type        = number
  description = "Read capacity units"
  default     = 5
}

variable "write_capacity" {
  type        = number
  description = "Write capacity units"
  default     = 5
}

variable "tags" {
  type        = map(string)
  description = "Additional resource tags"
  default     = {}
}
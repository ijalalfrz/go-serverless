resource "aws_dynamodb_table" "table" {
  name           = "${var.table_name}_${var.environment}"
  billing_mode   = "PROVISIONED"
  read_capacity  = var.read_capacity
  write_capacity = var.write_capacity
  hash_key       = "PK"

  attribute {
    name = "PK"
    type = "S"
  }

  tags = merge(
    var.tags,
    {
      Environment = var.environment
      ManagedBy   = "terraform"
    }
  )

  lifecycle {
    prevent_destroy = true
    ignore_changes  = all # Ignores all changes to existing table
  }

}
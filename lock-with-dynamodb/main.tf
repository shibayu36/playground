terraform {
  cloud {
    organization = "shibayu36"
    workspaces {
      name = "lock-with-dynamodb"
    }
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_dynamodb_table" "dynamodb_lock_table" {
  name         = "lock-table"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  #   attribute {
  #     name = "ExpiredAt"
  #     type = "S" # RFC3339
  #   }

  #   attribute {
  #     name = "ReleaseID"
  #     type = "S"
  #   }

  #   attribute {
  #     name = "Time"
  #     type = "S" # RFC3339
  #   }
}

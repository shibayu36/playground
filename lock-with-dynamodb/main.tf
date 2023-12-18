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

resource "aws_iam_user" "lock_with_dynamodb_user" {
  name = "lock-with-dynamodb-user"
}

resource "aws_iam_access_key" "lock_with_dynamodb_user_key" {
  user = aws_iam_user.lock_with_dynamodb_user.name
}

resource "aws_iam_user_policy" "lock_with_dynamodb_user_policy" {
  name = "lock-with-dynamodb-user-policy"
  user = aws_iam_user.lock_with_dynamodb_user.name

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "dynamodb:GetItem",
        "dynamodb:PutItem",
        "dynamodb:UpdateItem",
        "dynamodb:DeleteItem",
        "dynamodb:Query",
        "dynamodb:Scan"
      ],
      "Resource": "${aws_dynamodb_table.dynamodb_lock_table.arn}"
    }
  ]
}
EOF
}

output "access_key" {
  value = aws_iam_access_key.lock_with_dynamodb_user_key.id
}

output "secret_key" {
  sensitive = true
  value     = aws_iam_access_key.lock_with_dynamodb_user_key.secret
}

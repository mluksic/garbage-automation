# Configure the AWS Provider
provider "aws" {
  region = "eu-central-1"
}

# Configure security roles
resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

# Define Lambda function
resource "aws_lambda_function" "garbageAutomation" {
  # If the file is not in the current working directory you will need to include a 
  # path.module in the filename.
  filename      = "../garbage.zip"
  function_name = "garbageAutomation"
  handler       = "garbage"
  runtime       = "go1.x"
  role          = aws_iam_role.iam_for_lambda.arn

  # The filebase64sha256() function is available in Terraform 0.11.12 and later
  # For Terraform 0.11.11 and earlier, use the base64sha256() function and the file() function:
  # source_code_hash = "${base64sha256(file("lambda_function_payload.zip"))}"
  source_code_hash = filebase64sha256("../garbage.zip")

  environment {
    variables = {
      APP_ENV = var.APP_ENV
      APP_PASSWORD=var.APP_PASSWORD
      FROM_EMAIL=var.FROM_EMAIL
      EMAIL_RECEIVERS=var.EMAIL_RECEIVERS
    }
  }
}


#####################
## EXTRA RESOURCES ##
#####################

# Create cloudwatch event rule
resource "aws_cloudwatch_event_rule" "every_workday_evening" {
  name                = "runGarbageCheck"
  description         = "Fires every workday evening at 6pm"
  schedule_expression = "cron(0 16 ? * MON-FRI *)"
}

# Create cloudwatch event target
resource "aws_cloudwatch_event_target" "check_every_workday_evening" {
  rule      = "${aws_cloudwatch_event_rule.every_workday_evening.name}"
  target_id = "lambda"
  arn       = "${aws_lambda_function.garbageAutomation.arn}"
}

# Create cloudwatch event rule
resource "aws_lambda_permission" "allow_cloudwatch_to_call_check_for_garbage" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.garbageAutomation.function_name}"
  principal     = "events.amazonaws.com"
  source_arn    = "${aws_cloudwatch_event_rule.every_workday_evening.arn}"
}

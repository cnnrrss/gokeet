data "aws_region" "current" {
}

# Define a Lambda function.
#
# The handler is the name of the executable for go1.x runtime.
resource "aws_lambda_function" "hello" {
  function_name    = "hello"
  filename         = "hello.zip"
  handler          = "hello"
  source_code_hash = "${base64sha256(file("hello.zip"))}"
  role             = "${aws_iam_role.hello.arn}"
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 1
}

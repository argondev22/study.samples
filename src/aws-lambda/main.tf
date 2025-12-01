# resource "aws_lambda_function" "python" {
#   function_name = "python-function"
#   role          = "arn:aws:iam::520196876033:role/s3-event-lambda-dev"
#   handler       = "main.lambda_handler"
#   runtime       = "python3.10"
#   filename      = "${path.module}/src/python/lambda_function.zip"
#   source_code_hash = filebase64sha256("${path.module}/src/python/lambda_function.zip")
# }

resource "aws_lambda_function" "go" {
  function_name = "go-function"
  role          = "arn:aws:iam::520196876033:role/s3-event-lambda-dev"
  handler       = "main.lambda_handler"
  runtime       = "provided.al2"
  filename      = "${path.module}/src/go/lambda_function.zip"
  source_code_hash = filebase64sha256("${path.module}/src/go/lambda_function.zip")
}
resource "aws_s3_bucket" "example" {
  bucket = "tfmap-example-bucket"
  acl    = "private"

  tags = {
    Name        = "ExampleBucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket" "ok" {
  bucket = "example-ok"
  acl    = "private"

  tags = {
    Name        = "TFMap"
    Environment = "dev"
  }
}

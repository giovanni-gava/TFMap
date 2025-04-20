resource "aws_s3_bucket" "broken" {
  bucket = "example-broken"
  acl    = "private"
}

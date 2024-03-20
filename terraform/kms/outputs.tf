output "kms_key" {
  description = "The values of the KMS key."
  value       = aws_kms_key.s3_buckets
}

output "kms_alias" {
  description = "The values of the KMS alias."
  value       = aws_kms_alias.s3_buckets
}

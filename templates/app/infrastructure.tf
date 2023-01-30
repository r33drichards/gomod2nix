terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.20.0"
    }
  }

  required_version = "~> 1.0"
  backend "http" {

  }
}

variable "aws_region" {
  description = "region to create aws app runner service & ecr repo"
  default = "us-east-1"
}
variable "image_tag"  {
  description = "image tag to be deployed to aws app runner"
  default = "latest"
}

resource "random_string" "bucket_name" {
  length  = 8
  special = false
  upper   = false
}

variable "bucket_name" {
    description = "name of the bucket to be created, defaults to random string"
    default = "go-cli-bin" + random_string.bucket_name.result
}

provider "aws" {
  region = var.aws_region
}

# s3 bucket 
resource "aws_s3_bucket" "myapp" {
  acl = "public-read"
  bucket = var.bucket_name
}

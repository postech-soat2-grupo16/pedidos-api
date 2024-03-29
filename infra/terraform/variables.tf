variable "aws_region" {
  type    = string
  default = "us-east-1"
}

variable "aws_az_a" {
  type    = string
  default = "us-east-1a"
}

variable "aws_az_b" {
  type    = string
  default = "us-east-1b"
}

variable "vpc_id" {
  type    = string
  default = "value"
}

variable "subnet_a" {
  type    = string
  default = "value"
}

variable "subnet_b" {
  type    = string
  default = "value"
}

variable "ecr_image" {
  description = "ECR Image"
  type        = string
  sensitive   = true
  default     = "null"
}

variable "execution_role_ecs" {
  description = "Execution role ECS"
  type        = string
  sensitive   = true
}

variable "desired_tasks" {
  description = "Mininum executing tasks"
  type        = number
  default     = 1
}

variable "ecs_cluster" {
  description = "Cluster ECS ARN"
  type        = string
  sensitive   = true
}

variable "sg_cluster_ecs" {
  description = "Cluster ECS Security group"
  type        = string
  default     = "value"
}

variable "lb_arn" {
  description = "Load Balancer ARN"
  type        = string
  sensitive   = true
}

variable "alb_fastfood_listener_arn" {
  description = "Default Listener ALB"
  type        = string
  sensitive   = true
}

variable "sqs_url" {
  description = "SQS Pedidos URL"
  type        = string
  sensitive   = true
  default     = ""
}

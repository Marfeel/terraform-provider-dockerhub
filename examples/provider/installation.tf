terraform {
  required_version = ">= 0.13"

  required_providers {
    dockerhub = {
      source  = "Marfeel/dockerhub"
      version = ">= 0.2.0"
    }
  }
}

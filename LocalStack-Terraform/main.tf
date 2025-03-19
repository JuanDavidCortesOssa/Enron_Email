terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

resource "docker_network" "app_network" {
  name = "app_network"
}

provider "aws" {
  access_key                  = "mock_access_key"
  secret_key                  = "mock_secret_key"
  region                      = "us-east-1"

  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    s3 = "http://localhost:4566"
  }
}

provider "docker" {}

resource "docker_image" "zincsearch" {
  name = "public.ecr.aws/zinclabs/zincsearch:latest"
}

resource "docker_container" "zincsearch" {
  name  = "zincsearch"
  image = docker_image.zincsearch.name
  ports {
    internal = 4080
    external = 4080
  }
  volumes {
    host_path      = "C:/Users/JUAN DAVID/Desktop/TechnicalTest/data"
    container_path = "/data"
  }
  env = [
    "ZINC_DATA_PATH=/data",
    "ZINC_FIRST_ADMIN_USER=admin",
    "ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123"
  ]
  networks_advanced {
    name = docker_network.app_network.name
  }
}

resource "docker_image" "email_server" {
  name         = "email-server:latest"
  build {
    context    = "../API" 
    dockerfile = "Dockerfile"
  }
  keep_locally = true
}

resource "docker_container" "email_server" {
  name  = "email-server"
  image = docker_image.email_server.name
  ports {
    internal = 8080
    external = 8080
  }
  depends_on = [docker_container.zincsearch]
  networks_advanced {
    name = docker_network.app_network.name
  }
}

# resource "docker_image" "vue_app" {
#   name         = "vue-app:latest"
#   build {
#     context    = ".."
#     dockerfile = "Dockerfile"
#   }
#   keep_locally = true
#   triggers = {
#     rebuild = filemd5("../Dockerfile")
#   }
# }

resource "docker_image" "vue_app" {
  name         = "vue-app:latest"
  keep_locally = true
}

resource "docker_container" "vue_app" {
  name  = "vue-app"
  image = docker_image.vue_app.name
  ports {
    internal = 80   
    external = 5173  
  }
  networks_advanced {
    name = docker_network.app_network.name
  }
}
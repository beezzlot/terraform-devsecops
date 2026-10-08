terraform {
  required_version = ">= 1.5.0, < 2.0.0"

  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = "~> 0.230"
    }
  }
}

provider "yandex" {
  zone = "ru-central1-a"
}

data "yandex_vpc_subnet" "default_subnet" {
  name = var.subnet_name
}

resource "yandex_compute_instance" "api" {
  name        = "api"
  zone        = var.zone
  platform_id = "standard-v3"

  resources {
    cores  = var.vm_specification.api.cores
    memory = var.vm_specification.api.memory
  }

  boot_disk {
    initialize_params {
      image_id = var.vm_specification.api.boot_disk
      size     = var.vm_specification.api.disk_size
    }
  }

  network_interface {
    subnet_id = data.yandex_vpc_subnet.default_subnet.id
    nat       = false
  }

  dynamic "secondary_disk" {
    for_each = var.extra_disk_ids

    content {
      disk_id     = secondary_disk.value
      device_name = secondary_disk.key
    }
  }

  metadata = {
    "ssh-keys" = "ubuntu:${var.ssh_public_key}"
  }
}

resource "yandex_compute_instance" "web" {
  name        = "web"
  zone        = var.zone
  platform_id = "standard-v3"

  resources {
    cores  = var.vm_specification.web.cores
    memory = var.vm_specification.web.memory
  }

  boot_disk {
    initialize_params {
      image_id = var.vm_specification.web.boot_disk
      size     = var.vm_specification.web.disk_size
    }
  }

  network_interface {
    subnet_id = data.yandex_vpc_subnet.default_subnet.id
    nat       = false
  }

  metadata = {
    "ssh-keys" = "ubuntu:${var.ssh_public_key}"
  }
}

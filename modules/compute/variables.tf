variable "vm_specification" {
  type = map(object({
    cores     = number
    memory    = number
    disk_size = number
    boot_disk = string
  }))
}

variable "ssh_public_key" {
  type = string
}

variable "extra_disk_ids" {
  type = map(string)
}

variable "subnet_name" {
  type = string
}

variable "zone" {
  type = string
}

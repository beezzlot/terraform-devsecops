variable "extra_disk_names" {
  description = "Names of additional disks attached to the API VM"
  type        = list(string)
  default     = ["logs", "backup"]

  validation {
    condition     = length(var.extra_disk_names) > 0 && alltrue([for name in var.extra_disk_names : trimspace(name) != ""])
    error_message = "extra_disk_names must contain at least one non-empty name."
  }
}

variable "extra_disk_type" {
  description = "Type of additional disks"
  type        = string
  default     = "network-hdd"

  validation {
    condition     = contains(["network-hdd", "network-ssd"], var.extra_disk_type)
    error_message = "extra_disk_type must be network-hdd or network-ssd."
  }
}

variable "extra_disk_size" {
  description = "Size of every additional disk in GB"
  type        = number
  default     = 20

  validation {
    condition     = var.extra_disk_size > 1 && floor(var.extra_disk_size) == var.extra_disk_size
    error_message = "extra_disk_size must be an integer greater than 1."
  }
}

variable "vm_specification" {
  description = "Specifications for web and api VMs"

  type = map(object({
    cores     = number
    memory    = number
    disk_size = number
    boot_disk = string
  }))

  default = {
    web = {
      cores     = 2
      memory    = 2
      disk_size = 10
      boot_disk = "fd845dr9j4h2aaq1m6ko"
    }

    api = {
      cores     = 2
      memory    = 2
      disk_size = 10
      boot_disk = "fd845dr9j4h2aaq1m6ko"
    }
  }

  validation {
    condition = (
      toset(keys(var.vm_specification)) == toset(["web", "api"]) &&
      alltrue([
        for vm in values(var.vm_specification) :
        vm.cores == 2 &&
        contains([2, 4], vm.memory) &&
        vm.disk_size > 1 &&
        trimspace(vm.boot_disk) != ""
      ])
    )

    error_message = "vm_specification must contain exactly web and api; cores must be 2, memory must be 2 or 4, disk_size must be greater than 1, and boot_disk must be non-empty."
  }
}

variable "ssh_public_key" {
  description = "SSH public key used on both VMs"
  type        = string

  validation {
    condition     = trimspace(var.ssh_public_key) != ""
    error_message = "ssh_public_key must not be empty."
  }
}

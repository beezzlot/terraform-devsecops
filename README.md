# Terraform homework

Модуль создаёт в Yandex Cloud:

- две виртуальные машины: `web` и `api`;
- дополнительные диски и подключает их к VM `api`;
- внутренние IP-адреса VM и ID дисков.

Используются два дочерних модуля:

```text
modules/compute
modules/disks
```

## Подготовка

Создайте SSH-ключ, если его ещё нет:

```bash
ssh-keygen -t ed25519 -C "yandex-cloud"
```

Покажите публичный ключ:

```bash
cat ~/.ssh/id_ed25519.pub
```

Скопируйте его в `terraform.tfvars.dev` и `terraform.tfvars.prod` в переменную `ssh_public_key`.

Пример:

```hcl
ssh_public_key = "ssh-ed25519 AAAA... user@host"
```

Приватный ключ `~/.ssh/id_ed25519` нельзя добавлять в Git.

## Инициализация

Выполняйте команды из корня проекта:

```bash
terraform fmt -recursive
terraform init
terraform validate
```

## Окружение dev

Показать план:

```bash
terraform plan -var-file=terraform.tfvars.dev
```

Создать инфраструктуру:

```bash
terraform apply -var-file=terraform.tfvars.dev
```

Показать outputs:

```bash
terraform output
```

Удалить инфраструктуру после проверки:

```bash
terraform destroy -var-file=terraform.tfvars.dev
```

## Окружение prod

Показать план:

```bash
terraform plan -var-file=terraform.tfvars.prod
```

Создать инфраструктуру:

```bash
terraform apply -var-file=terraform.tfvars.prod
```

Удалить инфраструктуру:

```bash
terraform destroy -var-file=terraform.tfvars.prod
```

## Важно

Перед запуском замените тестовый SSH-ключ в `.tfvars` на свой публичный ключ.

После проверки обязательно выполните `terraform destroy`, чтобы удалить созданные ресурсы.

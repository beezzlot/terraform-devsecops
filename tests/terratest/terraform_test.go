// terraform_test.go реализован приемущественно с помощью ИИ, опыта программирования на GO хватило не намного...
// В этом честно признаюсь!!!

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformDevInfrastructure(t *testing.T) {
	terraformDir := "../.."

	validateOptions := &terraform.Options{
		TerraformDir: terraformDir,
		NoColor:      true,
	}

	planOptions := &terraform.Options{
		TerraformDir: terraformDir,
		VarFiles: []string{
			"terraform.tfvars.dev",
		},
		NoColor: true,
	}

	t.Run("TerraformValidate", func(t *testing.T) {
		terraform.Init(t, validateOptions)
		terraform.Validate(t, validateOptions)
	})

	t.Run("TerraformPlan", func(t *testing.T) {
		terraform.InitAndPlan(t, planOptions)
	})

	t.Run("ExpectedVMs", func(t *testing.T) {
		terraform.InitAndPlan(t, planOptions)

		plan := terraform.Plan(t, planOptions)

		assert.Contains(
			t,
			plan,
			`yandex_compute_instance" "api`,
			"План должен содержать VM api",
		)

		assert.Contains(
			t,
			plan,
			`yandex_compute_instance" "web`,
			"План должен содержать VM web",
		)
	})

	t.Run("NoPublicIPs", func(t *testing.T) {
		terraform.InitAndPlan(t, planOptions)

		plan := terraform.Plan(t, planOptions)

		assert.Contains(
			t,
			plan,
			"nat            = false",
			"Сетевой интерфейс api/web не должен иметь NAT",
		)
	})

	t.Run("APIDisks", func(t *testing.T) {
		terraform.InitAndPlan(t, planOptions)

		plan := terraform.Plan(t, planOptions)

		assert.Contains(
			t,
			plan,
			`device_name = "backup"`,
			"План должен содержать диск backup",
		)

		assert.Contains(
			t,
			plan,
			`device_name = "logs"`,
			"План должен содержать диск logs",
		)

		assert.Contains(
			t,
			plan,
			"size        = 20",
			"Дополнительные диски должны иметь размер 20 ГБ",
		)
	})

	t.Run("DevEnvironmentParameters", func(t *testing.T) {
		terraform.InitAndPlan(t, planOptions)

		plan := terraform.Plan(t, planOptions)

		assert.Contains(
			t,
			plan,
			`zone                      = "ru-central1-a"`,
			"VM должны создаваться в зоне ru-central1-a",
		)

		assert.Contains(
			t,
			plan,
			`platform_id               = "standard-v3"`,
			"VM должны использовать платформу standard-v3",
		)
	})
}
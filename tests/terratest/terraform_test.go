package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformDevInfrastructure(t *testing.T) {
	t.Parallel()

	terraformDir := "../.."

	tfOptions := &terraform.Options{
		TerraformDir: terraformDir,
		VarFiles: []string{
			"terraform.tfvars.dev",
		},
		NoColor: true,
	}

	defer terraform.Destroy(t, tfOptions)

	t.Run("TerraformValidate", func(t *testing.T) {
		terraform.Init(t, tfOptions)
		terraform.Validate(t, tfOptions)
	})

	t.Run("TerraformPlan", func(t *testing.T) {
		terraform.Init(t, tfOptions)
		plan := terraform.Plan(t, tfOptions)

		assert.NotEmpty(t, plan, "Terraform plan не должен быть пустым")
	})

	t.Run("ExpectedVMs", func(t *testing.T) {
		terraform.InitAndPlan(t, tfOptions)

		webName := terraform.Output(t, tfOptions, "web_instance_name")
		apiName := terraform.Output(t, tfOptions, "api_instance_name")

		assert.NotEmpty(t, webName, "Должна существовать web VM")
		assert.NotEmpty(t, apiName, "Должна существовать api VM")
		assert.NotEqual(t, webName, apiName, "web и api должны быть разными VM")
	})

	t.Run("NoPublicIPs", func(t *testing.T) {
		terraform.InitAndPlan(t, tfOptions)

		webPublicIP := terraform.Output(t, tfOptions, "web_public_ip")
		apiPublicIP := terraform.Output(t, tfOptions, "api_public_ip")

		assert.Empty(t, webPublicIP, "web VM не должна иметь публичный IP")
		assert.Empty(t, apiPublicIP, "api VM не должна иметь публичный IP")
	})

	t.Run("APIDisks", func(t *testing.T) {
		terraform.InitAndPlan(t, tfOptions)

		apiDiskIDs := terraform.OutputList(t, tfOptions, "api_disk_ids")

		assert.NotEmpty(t, apiDiskIDs, "У api должны быть дополнительные диски")
		assert.GreaterOrEqual(
			t,
			len(apiDiskIDs),
			2,
			"У api должно быть минимум два дополнительных диска",
		)
	})

	t.Run("DevEnvironmentParameters", func(t *testing.T) {
		terraform.InitAndPlan(t, tfOptions)

		environment := terraform.Output(t, tfOptions, "environment")

		assert.Equal(t, "dev", environment, "Должно использоваться окружение dev")
	})
}
package terratest

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformValidateAndPlan(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../..",
		VarFiles:     []string{"terraform.tfvars.dev"},
	})

	// Test 1: Terraform validate
	t.Run("TerraformValidate", func(t *testing.T) {
		terraform.Validate(t, terraformOptions)
	})

	// Test 2: Terraform plan
	t.Run("TerraformPlan", func(t *testing.T) {
		terraform.InitAndPlan(t, terraformOptions)
	})
}

func TestNoPublicIP(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../..",
		VarFiles:     []string{"terraform.tfvars.dev"},
	})

	// Test 3: Check no public IP assigned
	t.Run("NoPublicIP", func(t *testing.T) {
		planOutput := terraform.InitAndPlan(t, terraformOptions)
		
		// Verify public_ip is not explicitly assigned in the plan
		assert.NotContains(t, planOutput, "public_ip = \"", 
			"VM should not have public IP assigned")
	})
}

func TestInfrastructureComposition(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../..",
		VarFiles:     []string{"terraform.tfvars.dev"},
	})

	// Test 4: Check infrastructure composition
	t.Run("InfrastructureComposition", func(t *testing.T) {
		terraform.InitAndPlan(t, terraformOptions)
		
		// Check that outputs exist
		outputs := terraform.OutputAll(t, terraformOptions)
		
		// Verify disk_ids output exists for api VM
		assert.Contains(t, outputs, "disk_ids", 
			"Should have disk_ids output for api VM")
	})
}

func TestDiskProperties(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../..",
		VarFiles:     []string{"terraform.tfvars.dev"},
	})

	// Test 5: Additional check - disk resources exist
	t.Run("DiskProperties", func(t *testing.T) {
		planOutput := terraform.InitAndPlan(t, terraformOptions)
		
		// Verify compute disks are created
		assert.Contains(t, planOutput, "yandex_compute_disk",
			"Should create compute disks")
	})
}
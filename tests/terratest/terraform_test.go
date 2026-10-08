package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func terraformOptions(t *testing.T) *terraform.Options {
	t.Helper()

	rootDir, err := filepath.Abs("../..")
	require.NoError(t, err)

	return &terraform.Options{
		TerraformDir: rootDir,
		VarFiles: []string{
			filepath.Join(rootDir, "terraform.tfvars-4.dev"),
		},
		EnvVars: map[string]string{
			"TF_IN_AUTOMATION":   "true",
			"TF_INPUT":           "false",
			"TF_VAR_ssh_public_key": os.Getenv("SSH_PUBLIC_KEY"),
		},
	}
}

func TestTerraformValidate(t *testing.T) {
	t.Parallel()

	options := terraformOptions(t)
	terraform.InitAndValidate(t, options)
}

func TestTerraformPlan(t *testing.T) {
	t.Parallel()

	options := terraformOptions(t)
	terraform.InitAndValidate(t, options)
	terraform.Plan(t, options)
}

func TestExpectedVMs(t *testing.T) {
	t.Parallel()

	options := terraformOptions(t)
	terraform.InitAndValidate(t, options)

	plan := terraform.InitAndPlan(t, options)

	planJSON := terraform.ShowPlanJson(t, options, plan)

	require.Contains(t, string(planJSON), "module.compute.yandex_compute_instance.api")
	require.Contains(t, string(planJSON), "module.compute.yandex_compute_instance.web")
}

func TestVMsHaveNoPublicIP(t *testing.T) {
	t.Parallel()

	options := terraformOptions(t)
	terraform.InitAndValidate(t, options)

	plan := terraform.InitAndPlan(t, options)
	planJSON := terraform.ShowPlanJson(t, options, plan)

	planText := string(planJSON)

	require.Contains(t, planText, `"nat":false`)
}

func TestAdditionalDisks(t *testing.T) {
	t.Parallel()

	options := terraformOptions(t)
	terraform.InitAndValidate(t, options)

	plan := terraform.InitAndPlan(t, options)
	planJSON := terraform.ShowPlanJson(t, options, plan)

	planText := string(planJSON)

	require.Contains(t, planText, "module.disks.yandex_compute_disk.extra_disks")
	require.Contains(t, planText, "logs")
	require.Contains(t, planText, "backup")
}

func TestAPIDisks(t *testing.T) {
	t.Parallel()

	options := terraformOptions(t)
	terraform.InitAndValidate(t, options)

	plan := terraform.InitAndPlan(t, options)
	planJSON := terraform.ShowPlanJson(t, options, plan)

	planText := string(planJSON)

	require.Contains(t, planText, "secondary_disk")
}

func TestDevVMParameters(t *testing.T) {
	t.Parallel()

	options := terraformOptions(t)
	terraform.InitAndValidate(t, options)

	plan := terraform.InitAndPlan(t, options)
	planJSON := terraform.ShowPlanJson(t, options, plan)

	planText := string(planJSON)

	require.Contains(t, planText, `"cores":2`)
	require.Contains(t, planText, `"memory":2`)
}
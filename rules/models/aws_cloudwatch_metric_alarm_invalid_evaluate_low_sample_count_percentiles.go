// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule() *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule {
	return &AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "evaluate_low_sample_count_percentiles",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_evaluate_low_sample_count_percentiles"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"evaluate_low_sample_count_percentiles must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"evaluate_low_sample_count_percentiles must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

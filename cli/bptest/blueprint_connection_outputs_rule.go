package bptest

import (
	"fmt"
	"github.com/hashicorp/go-version"
)

type BlueprintConnectionOutputsRule struct{}

func (r *BlueprintConnectionOutputsRule) name() string {
	return "blueprint_connection_outputs_rule"
}

func (r *BlueprintConnectionOutputsRule) enabled() bool {
	return true
}

func (r *BlueprintConnectionOutputsRule) check(ctx lintContext) error {
	// Check if Spec or Interfaces is nil to avoid null pointer dereference
	if ctx.metadata == nil || ctx.metadata.Spec == nil || ctx.metadata.Spec.Interfaces == nil {
		fmt.Println("metadata, spec, or interfaces are nil")
		return nil
	}

	// build list of outputs content

	for _, variable := range ctx.metadata.Spec.Interfaces.Variables {
		if variable == nil {
			continue // Skip if variable is nil
		}

		for _, conn := range variable.Connections {
			if conn == nil || conn.Spec == nil {
				continue // Skip if connection or Spec is nil
			}
			// Implement the output checks
			// - validate the output exist
			// - connection output type compatibility with referenced input variables
			if conn.Spec.OutputExpr == nil || conn.Spec.OutputExpr == "" {
				return fmt.Errorf("invalid output_expr: %w", conn.Spec.OutputExpr)
			}
		}
	}
	return nil
}

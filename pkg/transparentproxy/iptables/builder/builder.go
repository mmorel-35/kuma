package builder

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/kumahq/kuma/v2/pkg/transparentproxy/config"
	"github.com/kumahq/kuma/v2/pkg/transparentproxy/iptables/tables"
)

func BuildIPTablesForRestore(cfg config.InitializedConfigIPvX) string {
	// filter out any empty sets
	result := slices.DeleteFunc([]string{
		tables.BuildRulesForRestore(cfg, buildRawTable(cfg)),
		tables.BuildRulesForRestore(cfg, buildNatTable(cfg)),
		tables.BuildRulesForRestore(cfg, buildMangleTable(cfg)),
	}, func(s string) bool { return s == "" })

	separator := "\n"
	if cfg.Verbose {
		separator = "\n\n"
	}

	return strings.Join(result, separator) + "\n"
}

func RestoreIPTables(ctx context.Context, cfg config.InitializedConfig) (string, error) {
	cfg.Logger.Info("kumactl is about to apply the iptables rules that will enable transparent proxying on the machine. The SSH connection may drop. If that happens, just reconnect again")

	output, err := cfg.IPv4.Executables.Restore(
		ctx,
		BuildIPTablesForRestore(cfg.IPv4),
		false,
	)
	if err != nil {
		return "", fmt.Errorf("unable to restore iptables rules: %w", err)
	}

	if cfg.IPv6.Enabled() {
		ipv6Output, err := cfg.IPv6.Executables.Restore(
			ctx,
			BuildIPTablesForRestore(cfg.IPv6),
			false,
		)
		if err != nil {
			return "", fmt.Errorf("unable to restore ip6tables rules: %w", err)
		}

		output += ipv6Output
	}

	return output, nil
}

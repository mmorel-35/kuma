package dnsserver

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	kuma_dp "github.com/kumahq/kuma/v2/pkg/config/app/kuma-dp"
)

//go:embed Corefile
var config embed.FS

func WriteCorefile(cfg kuma_dp.DNS, config []byte) (string, error) {
	configFile := filepath.Join(cfg.ConfigDir, "Corefile")
	if err := writeFile(configFile, config, 0o600); err != nil {
		return "", fmt.Errorf("failed to persist coredns Corefile on disk: %w", err)
	}
	return configFile, nil
}

func writeFile(filename string, data []byte, perm os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		return err
	}
	return os.WriteFile(filename, data, perm)
}

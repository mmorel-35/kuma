package defaults

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"

	kuma_cp "github.com/kumahq/kuma/v2/pkg/config/app/kuma-cp"
	config_core "github.com/kumahq/kuma/v2/pkg/config/core"
	"github.com/kumahq/kuma/v2/pkg/core/resources/manager"
	"github.com/kumahq/kuma/v2/pkg/core/resources/store"
	"github.com/kumahq/kuma/v2/pkg/envoy/admin/tls"
)

func EnsureEnvoyAdminCaExists(
	ctx context.Context,
	resManager manager.ResourceManager,
	logger logr.Logger,
	cfg kuma_cp.Config,
) error {
	if cfg.Mode == config_core.Global {
		return nil // Envoy Admin CA is not synced in multizone env and not needed in Global CP.
	}
	_, err := tls.LoadCA(ctx, resManager)
	if err == nil {
		logger.V(1).Info("Envoy Admin CA already exists. Skip creating Envoy Admin CA.")
		return nil
	}
	if !store.IsNotFound(err) {
		return fmt.Errorf("error while loading envoy admin CA: %w", err)
	}
	pair, err := tls.GenerateCA()
	if err != nil {
		return fmt.Errorf("could not generate envoy admin CA: %w", err)
	}
	if err := tls.CreateCA(ctx, *pair, resManager); err != nil {
		return fmt.Errorf("could not create envoy admin CA: %w", err)
	}
	logger.Info("Envoy Admin CA created")
	return nil
}

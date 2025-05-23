package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.globalso.dev/x/telemetry/collector"
	"go.globalso.dev/x/telemetry/common"
	"go.globalso.dev/x/telemetry/config"
	"go.globalso.dev/x/telemetry/internal"
)

func main() {
	cmd := &cobra.Command{
		Use:   "collector",
		Short: "Telemetry Collector Agent",
		Long: `Telemetry Collector Agent is a service that collects, processes,
and exports telemetry data. It is a component of the OpenTelemetry project.`,
		PreRunE: preRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			initial := config.Default()

			if err := viper.UnmarshalKey("telemetry", initial); err != nil {
				return fmt.Errorf("failed to unmarshal telemetry configuration: %w", err)
			}

			initial.Resource = common.NewResource(
				common.WithName("collector"),
				common.WithVersion(internal.Version),
			)

			a, err := collector.Initialize(cmd.Context(), initial)
			if err != nil {
				return err
			}

			return a.Start(cmd.Context())
		},
	}

	cmd.PersistentFlags().String("config", "", "config file (default is agent.yaml)")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

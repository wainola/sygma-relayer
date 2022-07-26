// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package cmd

import (
	deploy2 "github.com/ChainSafe/sygma/chains/evm/cli/deploy"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/ChainSafe/sygma-core/chains/evm/cli/deploy"
	"github.com/ChainSafe/sygma-core/flags"

	"github.com/ChainSafe/sygma/example/app"
)

var (
	rootCMD = &cobra.Command{
		Use: "",
	}
	runCMD = &cobra.Command{
		Use:   "run",
		Short: "Run example app",
		Long:  "Run example app",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := app.Run(); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	flags.BindFlags(runCMD)
}

func Execute() {
	rootCMD.AddCommand(runCMD, deploy.DeployEVM, deploy2.LocalSetupCmd)
	if err := rootCMD.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to execute root cmd")
	}
}

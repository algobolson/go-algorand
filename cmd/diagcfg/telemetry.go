// Copyright (C) 2019 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/algorand/go-algorand/logging"
)

var (
	nodeName string
	uri      string
)

func init() {
	telemetryCmd.AddCommand(telemetryStatusCmd)
	telemetryCmd.AddCommand(telemetryEnableCmd)
	telemetryCmd.AddCommand(telemetryDisableCmd)
	telemetryCmd.AddCommand(telemetryNameCmd)
	telemetryCmd.AddCommand(telemetryEndpointCmd)

	// Enable Logging : node name
	telemetryNameCmd.Flags().StringVarP(&nodeName, "name", "n", "", "Friendly-name to use for node")
	telemetryEndpointCmd.Flags().StringVarP(&uri, "endpoint", "e", "", "Endpoint's URI")
}

var telemetryCmd = &cobra.Command{
	Use:   "telemetry",
	Short: "Control and manage Algorand logging",
	Long:  `Enable/disable and configure Algorand remote logging`,
	Run:   telemetryStatusCmd.Run,
}

var telemetryStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the node's telemetry status",
	Long:  `Print the node's telemetry status`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := logging.EnsureTelemetryConfig(nil, "")

		// If error loading config, can't disable / no need to disable
		if err != nil {
			fmt.Println(err)
			fmt.Println(loggingNotConfigured)
		} else if cfg.Enable == false {
			fmt.Println(loggingNotEnabled)
		} else {
			fmt.Printf(loggingEnabled, cfg.Name, cfg.GUID)
		}
	},
}

var telemetryEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable Algorand remote logging",
	Long:  `Enable Algorand remote logging`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := logging.EnsureTelemetryConfig(nil, "")

		// If error loading config, can't disable / no need to disable
		if err != nil {
			return
		}

		cfg.Enable = true
		cfg.Save(cfg.FilePath)
		fmt.Printf("Telemetry logging enabled: Name = %s, Guid = %s\n", cfg.Name, cfg.GUID)
	},
}

var telemetryDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable Algorand remote logging",
	Long:  `Disable Algorand remote logging`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := logging.EnsureTelemetryConfig(nil, "")

		// If error loading config, can't disable / no need to disable
		if err != nil {
			return
		}

		cfg.Enable = false
		cfg.Save(cfg.FilePath)
		fmt.Printf("Telemetry logging disabled: Name = %s, Guid = %s\n", cfg.Name, cfg.GUID)
	},
}

var telemetryNameCmd = &cobra.Command{
	Use:   "name -n nodeName",
	Short: "Enable Algorand remote logging",
	Long:  `Enable Algorand remote logging with specified node name`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := logging.EnsureTelemetryConfig(nil, "")
		if err != nil {
			fmt.Println(err)
			return
		}
		cfg.Enable = true
		if len(nodeName) > 0 {
			cfg.Name = nodeName
		}
		cfg.Save(cfg.FilePath)
		fmt.Printf("Telemetry logging: Name = %s, Guid = %s\n", cfg.Name, cfg.GUID)
	},
}

var telemetryEndpointCmd = &cobra.Command{
	Use:   "endpoint -e <url>",
	Short: "sets the \"URI\" property",
	Long:  `Sets the "URI" property in the telemetry configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := logging.EnsureTelemetryConfig(nil, "")
		if err != nil {
			fmt.Println(err)
			return
		}
		cfg.URI = uri
		cfg.Save(cfg.FilePath)
		fmt.Printf("Telemetry logging: Name = %s, Guid = %s, URI = %s\n", cfg.Name, cfg.GUID, cfg.URI)
	},
}

package cmd

import (
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/config"
	"github.com/spf13/pflag"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "eigenlayer-rewards-updater",
	Short: "Proof generation for rewards and claims",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {

	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eigenlayer-rewards-updater/config.yaml)")
	rootCmd.PersistentFlags().Bool("debug", false, "'true' or 'false'")
	rootCmd.PersistentFlags().String("dd-statsd-url", "", "URL to use for DataDog StatsD. If empty, DD_DOGSTATSD_URL will be used")
	rootCmd.PersistentFlags().Bool("enable-statsd", true, "Enable/disable statsd metrics collection")
	rootCmd.PersistentFlags().Bool("enable-tracing", true, "Enable/disable tracing")
	rootCmd.PersistentFlags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	rootCmd.PersistentFlags().String("environment", "dev", "The environment to use")
	rootCmd.PersistentFlags().String("network", "localnet", "Which network to use")
	rootCmd.PersistentFlags().String("private-key", "", "An ethereum private key")
	rootCmd.PersistentFlags().String("rewards-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the rewards coordinator contract")
	rootCmd.PersistentFlags().String("proof-store-base-url", "", "HTTP base url where data is stored")
	rootCmd.PersistentFlags().String("sidecar-rpc-url", "", "Sidecar RPC URL")
	rootCmd.PersistentFlags().Bool("sidecar-insecure-rpc", false, "Use insecure gRPC connection to sidecar")
	rootCmd.PersistentFlags().Bool("pushgateway-enabled", false, "Enable/disable pushgateway metrics collection")
	rootCmd.PersistentFlags().String("pushgateway-url", "", "URL to use for Pushgateway. This option is ignored if pushgateway-enable is not set.")

	rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(config.KebabToSnakeCase(f.Name), f)
		viper.BindEnv(f.Name)
	})
	initConfig(rootCmd)
}

func initConfig(cmd *cobra.Command) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in following directories
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$HOME/.eigenlayer-rewards-updater/")
		viper.AddConfigPath("/etc/eigenlayer-rewards-updater/")
		viper.AddConfigPath(".")
	}

	viper.SetEnvPrefix("eigenlayer")

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
	// https://github.com/carolynvs/stingoftheviper/blob/main/main.go#L96-L98
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Error loading config file: %s\n", err)
	}
}

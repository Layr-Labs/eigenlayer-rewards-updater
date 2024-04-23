package cmd

import (
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/spf13/pflag"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "eigenlayer-payment-updater",
	Short: "Proof generation for payments and claims",
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eigenlayer-payment-updater/config.yaml)")
	rootCmd.PersistentFlags().Bool("debug", false, "'true' or 'false'")

	rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(config.KebabToSnakeCase(f.Name), f)
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
		viper.AddConfigPath("$HOME/.eigenlayer-payment-updater/")
		viper.AddConfigPath("/etc/eigenlayer-payment-updater/")
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
	}
}

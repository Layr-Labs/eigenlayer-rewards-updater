package cmd

import (
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// updaterCmd represents the updater command
var updaterCmd = &cobra.Command{
	Use:   "updater",
	Short: "Generate and update payments merkle tree",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewUpdaterConfig()

	},
}

func init() {
	fmt.Println("Updater init")
	rootCmd.AddCommand(updaterCmd)

	updaterCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	updaterCmd.Flags().String("private-key", "", "An ethereum private key")
	updaterCmd.Flags().String("aws-access-key-id", "", "AWS access key ID")
	updaterCmd.Flags().String("aws-secret-access-key", "", "AWS secret access key")
	updaterCmd.Flags().String("aws-region", "us-east-1", "us-east-1")
	updaterCmd.Flags().String("s3-output-bucket", "", "s3://<bucket name and path>")

	updaterCmd.Flags().VisitAll(func(f *pflag.Flag) {
		fmt.Printf("flag: %v\n", config.KebabToSnakeCase(f.Name))
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
	})

}

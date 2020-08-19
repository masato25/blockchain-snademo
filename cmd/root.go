package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	seed          bool
	transformName string
	input         string
	parameter     string
	debug         bool
	rootCmd       = &cobra.Command{
		Use: "cobra",
		//Short: "Blockchain transforms of Maltego",
		//Long: `Blockchain transforms`,
	}
)

// Execute executes the root command.
func Execute() error {
	err := rootCmd.Execute()
	viper.Set("seed", seed)
	viper.Set("transformName", transformName)
	viper.Set("input", input)
	viper.Set("debug", debug)
	return err
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&seed, "seed", false, "產生資料庫檔案")
	rootCmd.PersistentFlags().StringVar(&transformName, "transform_name", "", "transform 函式名稱")
	rootCmd.PersistentFlags().StringVar(&input, "input", "", "查詢key")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
}

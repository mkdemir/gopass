package cmd

import (
	"fmt"
	"gopass/pkg/password"

	"github.com/spf13/cobra"
)

var length int
var digits, symbols bool

var rootCmd = &cobra.Command{
	Use:   "gopass",
	Short: "Generates a secure password",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := password.GeneratePassword(length, digits, symbols)

		if err != nil {
			fmt.Println("Error generating password: ", err)
			return
		}
		fmt.Println("Generated Password: ", p)
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&length, "length", "l", 8, "Length of the password")
	rootCmd.PersistentFlags().BoolVarP(&digits, "digits", "d", false, "Include digits in the password")
	rootCmd.PersistentFlags().BoolVarP(&symbols, "symbols", "s", false, "Include symbols in the password")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

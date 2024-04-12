package cmd

import (
	"fmt"

	"github.com/mkdemir/gopass/pkg/password"

	"github.com/spf13/cobra"
)

var length int
var digits, symbols, upper bool

var rootCmd = &cobra.Command{
	Use:   "gopass",
	Short: "Generates a secure password",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := password.GeneratePassword(length, digits, symbols, upper)

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
	rootCmd.PersistentFlags().BoolVarP(&upper, "upper", "u", false, "Include upper letter in the password")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

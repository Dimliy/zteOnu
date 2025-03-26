package cmd

import (
	
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print current version of zteOnu",
		Run: func(cmd *cobra.Command, args []string) {
			version.Show()
		},
		Args: cobra.NoArgs,
	})
}

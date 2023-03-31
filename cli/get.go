package cli

import "github.com/spf13/cobra"

var (
	url string
)

var GetCmd = &cobra.Command{
	Use:     "get url",
	Short:   "get file",
	Example: "get -i url",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	f := GetCmd.PersistentFlags()
	f.StringVarP(&url, "url", "i", "", "get url from http")
	RootCmd.AddCommand(GetCmd)
}

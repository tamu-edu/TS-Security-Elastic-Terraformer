package cmd

import (
	myrasec_terraforming "github.com/tamu-edu/TS-Security-Elastic-Terraformer/providers/myrasec"
	"github.com/tamu-edu/TS-Security-Elastic-Terraformer/terraformutils"

	"github.com/spf13/cobra"
)

// newCmdMyrasecImporter
func newCmdMyrasecImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "myrasec",
		Short: "Import current state to Terraform configuration from Myra Security",
		Long:  "Import current state to Terraform configuration from Myra Security",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newMyrasecProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newMyrasecProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "domain", "")
	return cmd
}

// newMyrasecProvider
func newMyrasecProvider() terraformutils.ProviderGenerator {
	return &myrasec_terraforming.MyrasecProvider{}
}

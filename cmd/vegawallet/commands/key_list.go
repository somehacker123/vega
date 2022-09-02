package cmd

import (
	"fmt"
	"io"

	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/cli"
	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/flags"
	"code.vegaprotocol.io/vega/cmd/vegawallet/commands/printer"
	"code.vegaprotocol.io/vega/wallet/wallet"
	"code.vegaprotocol.io/vega/wallet/wallets"

	"github.com/spf13/cobra"
)

var (
	listKeysLong = cli.LongDesc(`
		List the keys of a given wallet.
	`)

	listKeysExample = cli.Examples(`
		# List all keys
		{{.Software}} key list --wallet WALLET
	`)
)

type ListKeysHandler func(*wallet.ListKeysRequest) (*wallet.ListKeysResponse, error)

func NewCmdListKeys(w io.Writer, rf *RootFlags) *cobra.Command {
	h := func(req *wallet.ListKeysRequest) (*wallet.ListKeysResponse, error) {
		s, err := wallets.InitialiseStore(rf.Home)
		if err != nil {
			return nil, fmt.Errorf("couldn't initialise wallets store: %w", err)
		}

		return wallet.ListKeys(s, req)
	}

	return BuildCmdListKeys(w, h, rf)
}

func BuildCmdListKeys(w io.Writer, handler ListKeysHandler, rf *RootFlags) *cobra.Command {
	f := &ListKeysFlags{}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List the keys of a given wallet",
		Long:    listKeysLong,
		Example: listKeysExample,
		RunE: func(_ *cobra.Command, _ []string) error {
			req, err := f.Validate()
			if err != nil {
				return err
			}

			resp, err := handler(req)
			if err != nil {
				return err
			}

			switch rf.Output {
			case flags.InteractiveOutput:
				PrintListKeysResponse(w, resp)
			case flags.JSONOutput:
				return printer.FprintJSON(w, resp)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&f.Wallet,
		"wallet", "w",
		"",
		"Name of the wallet to use",
	)
	cmd.Flags().StringVarP(&f.PassphraseFile,
		"passphrase-file", "p",
		"",
		"Path to the file containing the wallet's passphrase",
	)

	autoCompleteWallet(cmd, rf.Home)

	return cmd
}

type ListKeysFlags struct {
	Wallet         string
	PassphraseFile string
}

func (f *ListKeysFlags) Validate() (*wallet.ListKeysRequest, error) {
	req := &wallet.ListKeysRequest{}

	if len(f.Wallet) == 0 {
		return nil, flags.MustBeSpecifiedError("wallet")
	}
	req.Wallet = f.Wallet

	passphrase, err := flags.GetPassphrase(f.PassphraseFile)
	if err != nil {
		return nil, err
	}
	req.Passphrase = passphrase

	return req, nil
}

func PrintListKeysResponse(w io.Writer, resp *wallet.ListKeysResponse) {
	p := printer.NewInteractivePrinter(w)

	str := p.String()
	defer p.Print(str)

	for i, key := range resp.Keys {
		if i != 0 {
			str.NextLine()
		}
		str.Text("Name:       ").WarningText(key.Name).NextLine()
		str.Text("Public key: ").WarningText(key.PublicKey).NextLine()
	}
}

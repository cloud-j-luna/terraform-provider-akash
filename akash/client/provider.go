package client

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"terraform-provider-akash/akash/client/cli"
	"terraform-provider-akash/akash/client/types"
)

func (ak *AkashClient) SendManifest(dseq string, provider string, manifestLocation string) (string, error) {

	cmd := cli.AkashCli(ak).SendManifest(manifestLocation).
		SetDseq(dseq).SetProvider(provider).SetHome(ak.Config.Home).
		SetKeyringBackend(ak.Config.KeyringBackend).SetFrom(ak.Config.KeyName).
		SetNode(ak.Config.Node).OutputJson()

	out, err := cmd.Raw()
	if err != nil {
		return "", err
	}

	tflog.Debug(ak.ctx, fmt.Sprintf("Response content: %s", out))

	return string(out), nil
}

func (ak *AkashClient) GetLeaseStatus(seqs Seqs, provider string) (*types.LeaseStatus, error) {

	cmd := cli.AkashCli(ak).LeaseStatus().
		SetHome(ak.Config.Home).SetDseq(seqs.Dseq).SetGseq(seqs.Gseq).SetOseq(seqs.Oseq).
		SetNode(ak.Config.Node).SetProvider(provider).SetFrom(ak.Config.KeyName)

	leaseStatus := types.LeaseStatus{}
	err := cmd.DecodeJson(&leaseStatus)
	if err != nil {
		return nil, err
	}

	return &leaseStatus, nil
}

package secret

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"os"
	"strings"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// Encrypt [key] [value]
func Encrypt(cmd *cobra.Command, args []string) error {

	key := args[0]
	value := ""

	if len(args) > 1 {
		value = args[1]
	} else {
		// Prompt for secret value securely if not already provided
		fmt.Print("Enter secret value: ")
		bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return errors.Wrap(err, "failed to read secret value")
		}
		value = strings.TrimSpace(string(bytePassword))
		fmt.Println() // Move to the next line after password input
	}

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	ctx := cmd.Context()

	cf := client.NewApiClientFactory(url)
	secretApi := cf.Client(client.WithToken(token)).SecretSvcAPI

	isSecureRsp, _, err := secretApi.IsSecure(ctx).Execute()

	// this will mess up the yaml structure but that is intentional
	var returnErr error
	if err != nil {
		returnErr = errors.Wrap(err, "warning: cannot identify if the server is secure")
	} else if !isSecureRsp.IsSecure {
		returnErr = fmt.Errorf("warning: secret service is not secure")
	}

	rsp, _, err := secretApi.EncryptValue(ctx).
		Body(openapi.SecretSvcEncryptValueRequest{
			Value: openapi.PtrString(value),
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to encrypt value")
	}

	h := crc32.ChecksumIEEE([]byte(value))
	hash := hex.EncodeToString([]byte{
		byte(h >> 24), byte(h >> 16), byte(h >> 8), byte(h),
	})

	secret := openapi.SecretSvcSecret{
		Id:                openapi.PtrString(sdk.Id("secr")),
		Key:               openapi.PtrString(key),
		Encrypted:         openapi.PtrBool(true),
		Value:             rsp.Value,
		Checksum:          openapi.PtrString(hash),
		ChecksumAlgorithm: openapi.ChecksumAlgorithmCRC32.Ptr(),
	}

	bs, err := yaml.Marshal(secret)
	if err != nil {
		return errors.Wrap(err, "failed to encode user info")
	}

	fmt.Print(string(bs))

	if returnErr != nil {
		fmt.Println("# " + returnErr.Error())
	}

	return nil
}

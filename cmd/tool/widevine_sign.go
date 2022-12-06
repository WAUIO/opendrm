package tool

import (
	"encoding/hex"
	"fmt"

	"github.com/alfg/widevine"
	"github.com/spf13/cobra"
)

var (
	aesKey string
	aesIV  string
)

func init() {
	WidevineSign.PersistentFlags().StringVar(&aesKey, "aes-key", "", "Widevine AES Signing Key")
	WidevineSign.PersistentFlags().StringVar(&aesIV, "aes-iv", "", "Widevine AES Signing IV")
}

var WidevineSign = &cobra.Command{
	Use:   "widevine:sign",
	Short: "Sign a base64 string",
	Run: func(cmd *cobra.Command, args []string) {
		// decoding keys
		k, _ := hex.DecodeString(aesKey)
		iv, _ := hex.DecodeString(aesIV)

		// crypto := widevine.NewCrypto(k, iv)
		// sig := crypto.GenerateSignature([]byte(args[0]))
		// fmt.Println(sig)

		// Set Widevine options and create instance.
		options := widevine.Options{
			Key:      k,
			IV:       iv,
			Provider: "widevine_test",
		}
		wv := widevine.New(options)

		// Your video content ID, usually a GUID.
		contentID := "9F0712D9F0824EA2A0DE8991B63EBBCD"

		// Set policy options.
		policy := widevine.Policy{
			ContentID: contentID,
			Tracks:    []string{"HD"},
			DRMTypes:  []string{"WIDEVINE"},
			Policy:    "default",
		}

		// Make the request to generate or get a content key.
		resp := wv.GetContentKey(contentID, policy)

		// Response data from Widevine Cloud.
		fmt.Println("status: ", resp.Status)
		fmt.Println("drm: ", resp.DRM)
		for i, v := range resp.Tracks {
			fmt.Println("Track #", i+1)
			fmt.Println("key: ", v.Key)
			fmt.Println("key_id: ", v.KeyID)
			fmt.Println("type: ", v.Type)
			fmt.Println("drm_type: ", v.PSSH[0].DRMType)
			fmt.Println("data: ", v.PSSH[0].Data)
		}
		fmt.Println("already_used: ", resp.AlreadyUsed)
	},
}

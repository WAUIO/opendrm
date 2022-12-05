package tool

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/alfg/widevine"
	"github.com/spf13/cobra"

	"github.com/wauio/opendrm/src/core/key"
)

var (
	tracks   []string
	drmTypes []string
	policy   string
)

func init() {
	WidevineGenerate.PersistentFlags().StringVar(&aesKey, "aes-key", "", "Widevine AES Signing Key")
	WidevineGenerate.PersistentFlags().StringVar(&aesIV, "aes-iv", "", "Widevine AES Signing IV")
	WidevineGenerate.PersistentFlags().StringSliceVar(&tracks, "track", []string{}, "List of tracks in the media")
	WidevineGenerate.PersistentFlags().StringSliceVar(&drmTypes, "drm-type", []string{"WIDEVINE"}, "DRM Type to setup license for")
	WidevineGenerate.PersistentFlags().StringVar(&policy, "policy", "default", "Policy to set into the license")
}

var WidevineGenerate = &cobra.Command{
	Use:   "widevine:gen",
	Short: "Generate new key pair and related pssh box information",
	Run: func(cmd *cobra.Command, args []string) {
		// decoding keys
		k, _ := hex.DecodeString(aesKey)
		iv, _ := hex.DecodeString(aesIV)

		// Set Widevine options and create instance.
		options := widevine.Options{
			Key:      k,
			IV:       iv,
			Provider: "widevine_test",
		}
		wv := widevine.New(options)

		// Your video content ID, usually a GUID.
		var contentID string
		if len(args) == 0 {
			contentID = strings.ToLower(strings.ReplaceAll(key.GenerateUUID(), "-", ""))
		} else {
			contentID = args[0]
		}

		// Set policy options.
		policy := widevine.Policy{
			ContentID: contentID,
			Tracks:    tracks,
			DRMTypes:  drmTypes,
			Policy:    policy,
		}

		// Make the request to generate or get a content key.
		resp := wv.GetContentKey(contentID, policy)

		// Response data from Widevine Cloud.
		fmt.Println("ContentID: ", contentID)
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

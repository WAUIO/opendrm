/*
	Opendrm, an open source implementation of industry-grade DRM
	(Digital Rights Management) or Key System.
	Copyright (C) 2018  wilkk

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package application

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wauio/opendrm/src/core/key"
	"github.com/wauio/opendrm/src/core/license"
	"github.com/wauio/opendrm/src/core/server"
	"github.com/wauio/opendrm/src/handler"
)

type KeyResp struct {
	Key []byte `json:"key"`
	Kid string `json:"kid"`
}

func GenKey(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//kid := r.Form["kid"]

	kengen := key.NewKeyGenerator(nil)
	key, kid := kengen.GenRandKey()
	resp := KeyResp{
		Key: key,
		Kid: kid,
	}
	data, _ := json.Marshal(&resp)
	w.Write(data)
}

type LicenseRequest struct {
	// required
	DeviceId string `json:"device_id"`
	// required
	Kids []string `json:"kids"`
	// optional
	ClientId *string
	// optional
	ContentId *string `json:"content_id"`
}

type LicenseResp struct {
	DeviceId string
	Licenses []string
}

func AcquireLicense(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request failed. err=%s", err)
		return
	}
	log.Printf("data:%s", string(reqData))

	req := &LicenseRequest{}
	test := ""
	err = json.Unmarshal(reqData, &test)
	if err != nil {
		w.Header().Add("X-AxDRM-ErrorMessage", err.Error())
		log.Printf("Failed to parse request, LicenseRequest expected. err=%s", err)

		//simulate
		req.Kids = []string{"171E7392-D6CF-4755-986F-A11330659179"}
		req.DeviceId = "simulator"

		// return
	}
	log.Printf("kids:%v", req)

	// Query objects to be authorized
	objs := []string{"07fba7c4-a5d3-43b2-973b-0b474a0b9edf"}
	certId := "6C37D30C-881B-1AE7-9F44-C0B25715AF36"
	// Generate license
	lic := license.NewCommonLicense(req.Kids, objs, certId)
	licenseStr := lic.Base64String()

	resp := &LicenseResp{
		DeviceId: req.DeviceId,
		Licenses: []string{licenseStr},
	}

	// respData, err := json.Marshal(resp)
	// if err != nil {
	// 	log.Printf("Marshal response failed. Err=%s", err)
	// 	w.Header().Add("X-AxDRM-ErrorMessage", err.Error())
	// 	return
	// }
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("X-AxDRM-Identity", resp.DeviceId)
	w.Header().Add("X-AxDRM-Server", "elacity_drm")
	w.Header().Add("X-AxDRM-Version", "0.1.0")
	w.Write(lic.Serialize(true, true))
}

func Bootstrap(addr string) {
	log.Printf("running on %s", addr)
	keyServer := server.NewKeyServer(addr)
	keyServer.HandleFunc("/genkey", GenKey)
	keyServer.HandleFunc("/acquirelicense", AcquireLicense)
	keyServer.HandleFunc("/acquirelicense/clearkey", handler.AcquireLicenseForClearKey)
	keyServer.HandleFunc("/acquirelicense/widevine", handler.AcquireLicenseForWidevive)
	if err := keyServer.Start(); err != nil {
		log.Fatal(err)
	}
}

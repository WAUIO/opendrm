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

package server

import (
	"net/http"

	"github.com/rs/cors"
)

type KeyServer struct {
	*http.ServeMux
	server *http.Server
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func NewKeyServer(addr string) *KeyServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type", "X-AxDRM-Message"},
		Debug:          true,
	}).Handler(mux)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &KeyServer{
		mux,
		server,
	}
}

func (this *KeyServer) Start() error {
	return this.server.ListenAndServe()
}

// package sensorstation
package main

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"net/http"
)

func web() error {
	// web stuff
	http.Handle("/ws", wserv) // websocket
	http.Handle("/api/config", config)
	http.Handle("/api/quote", quote)
	http.Handle("/", http.FileServer(http.Dir(config.App))) // home app

	return http.ListenAndServe(config.Addr, nil)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "index.html")
}

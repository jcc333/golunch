// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (p *Page) save() error {
	fmt.Printf("Saving post %+v\n", *p)
	filename := p.Slug + ".json"
	bytes, err := json.Marshal(*p)
	if err != nil {
		return err
	}
	fmt.Println("Writing the file", filename)
	return ioutil.WriteFile(filename, bytes, 0600)
}

func loadPage(slug string) (*Page, error) {
	filename := slug + ".json"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var p Page
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/posts/"):]
	p, err := loadPage(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	var p Page
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bytes, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/" + p.Slug, http.StatusFound)
}

func muxGetAndPost(onGet, onPost http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			onGet(w, r)
		}
		if r.Method == http.MethodPost {
			onPost(w, r)
		}
	}
}

func main() {
	http.HandleFunc("/posts/", muxGetAndPost(viewHandler, saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

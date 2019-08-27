# Notes on Week One
- We're only using modules to avoid gopath stuff. If that means nothing to you, rejoice: you missed some terrible stuff.

## Main
- If Main Returns Without an os.Exit(1), it returns 0
```
func main() {
	http.HandleFunc("/posts/", muxGetAndPost(viewHandler, saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Functions
- Functions are first class values
- Functions enclose their environments
- Functions are all over the dang place
- The type of a function might look like: `func(int, string) string`
```
// This is a function called viewHandler which takes w, an http.ResponseWriter and r, a *http.Request and returns nothing
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
```

```
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
```


## Structs and Types
```
// This is a type called 'Page' which is a struct
type Page struct {
	// This is a field called 'Slug' which is a string. It has an annotation called 'json' set to "slug"
	Slug string `json:"slug"`
	// This is a field called 'Title' which is a string. It has an annotation called 'json' set to "title"
	Title string `json:"title"`
	// This is a field called 'Body' which is a string. It has an annotation called 'json' set to "body"
	Body  string `json:"body"`
}
```
- `struct`s are order-specific, potentially named tuples
- `struct` does not mean `class`

## Errors
- `error` is an `interface` specified [here](https://golang.org/ref/spec#Errors) and reproduced below:
```
// This is a type called 'error' which is an interface
// that specifies a function called 'Error' which takes no arguments and returns a string
type error interface { Error() string }
```
```
// This is a function on *Pages called save that takes no arguments and returns an error
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
```

## Multiple Returns
- go functions can return multiple arguments
- some of those arguments can be "zero-values" (every type has a zero value)
```
// This is a function called loadPage which takes an argument called slug which is a string and returns a pointer to a Page and an error
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
```

## Higher-Order/First-Class Functions
```
// This is a function called muxGetAndPost which takes onGet and onPost, both of which are http.HandlerFuncs and returns an http.HandlerFunc
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
```


## Usage examples

A few usage examples can be found below. See the documentation for the full list of supported functions.

### Validation Single Image
```go
var SingleImage http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20)

	magic := magicimage.New(r.MultipartForm)

  // magic.Required = false
  // magic.MaxFileSize = 4 << 20 (4MB)

	if err := magic.ValidateSingleImage("file"); err != nil {
		fmt.Fprint(rw, err)
		return
	}

  magic.SaveImages(200, 200, "out/this-is-slug", true)
  // all filename
  fmt.Println(magic.FileNames)

	fmt.Fprint(rw, "success")
}
```

### Validation Multiple Image
```go
var MultipleImage http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20)

	magic := magicimage.New(r.MultipartForm)

  // magic.MinFileInSlice = 1
  // magic.MaxFileInSlice = 10

	if err := magic.ValidateMultipleImage("files"); err != nil {
		fmt.Fprint(rw, err)
		return
	}

  magic.SaveImages(200, 200, "out/this-is-slug", true)
  // all filename
  fmt.Println(magic.FileNames)

	fmt.Fprint(rw, "success")
}
```

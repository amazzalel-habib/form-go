# FormGo

FormGo is a Go (Golang) package that simplifies the process of encoding a struct into URL form values by extracting tagged fields. It's a convenient tool for working with web forms, making it easy to map Go structs to URL form data.

## Features

- Encode Go structs into URL form values with minimal code.
- Supports string and integer fields.
- Easily customize field names using struct tags.

## Installation

To use FormGo in your Go project, you need to install it:

```bash
go get github.com/amazzalel-habib/form-go
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/amazzalel-habib/form-go"
	"net/url"
)

type MyForm struct {
	Name  string `form:"name"`
	Age   int    `form:"age"`
	Email string `form:"email"`
}

func main() {
	data := MyForm{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	formValues, err := formgo.EncodeFormValues(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(formValues.Encode())
}
````

In this example, we have a struct `MyForm`, and we use the `form` tag to specify the field names for the URL form values. The `formgo.EncodeFormValues` function encodes the struct into URL form values.

## Contributing

We welcome contributions from the community. If you find a bug or have an idea for an improvement, please open an issue or create a pull request.

## License

FormGo is licensed under the MIT License.
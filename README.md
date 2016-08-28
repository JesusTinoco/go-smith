# go-smith [![Build Status](https://travis-ci.org/JesusTinoco/go-smith.svg?branch=master)](https://travis-ci.org/JesusTinoco/go-smith) [![Go Report Card](https://goreportcard.com/badge/github.com/JesusTinoco/go-smith)](https://goreportcard.com/report/github.com/JesusTinoco/go-smith) [![GoDoc](https://godoc.org/github.com/JesusTinoco/go-smith/stacksmith?status.svg)](https://godoc.org/github.com/JesusTinoco/go-smith/stacksmith) 

go-smith is a Go client library for the [Stacksmith API](https://stacksmith.bitnami.com/api/v1). Check the folder [example](example) to learn how to use the Stacksmith API.

## Install
```
go get github.com/JesusTinoco/go-smith/stacksmith
```

## Documentation

Read [GoDoc](https://godoc.org/github.com/JesusTinoco/go-smith/stacksmith)

## Usage

```
APIKey := "<API_KEY_STACKSMITH>"
client := stacksmith.NewClient(APIKey, nil)

pag := &stacksmith.PaginationParams{Page: 1, PerPage: 100}

// Get all the stacks created
stacksList, _, _ := client.Stacks.List(pag)
fmt.Println(fmt.Sprintf("You have %d stacks.", len(stacksList.Items)))
```

## Contributing

Bug reports and pull requests are welcome.

## License

[MIT License](LICENSE)

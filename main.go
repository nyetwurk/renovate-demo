package main

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/twmb/franz-go/pkg/kadm"
)

func main() {
	kadm.NewClient(nil)
	validation.ValidateStruct(nil)
}

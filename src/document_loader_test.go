package main

import (
	"github.com/scalm/rbac/pkg/json_io"
	"github.com/scalm/rbac/pkg/rbac_document"
	"github.com/scalm/rbac/pkg/rbac_simple"
	"github.com/scalm/rbac/pkg/rbac_simple_document"
	"testing"
)

func TestLoader1(t *testing.T) {
	documentLoader := rbac_document.NewDocumentReader(
		json_io.NewStringReader(
			`{"subjects":[{"id":"A","roles":["R1"]}],"roles":[{"id":"R1","permissions":["edit"]}],"permissions":[{"id":"edit"}]}`),
		rbac_simple_document.NewVisitorAdapter(
			rbac_simple.NewFactory(),
			&rbac_simple.PrintVisitor{}))

	err := documentLoader.Read()
	if err != nil {
		t.Error(err)
		return
	}
}
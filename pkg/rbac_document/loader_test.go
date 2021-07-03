package rbac_document

import (
	"github.com/scalm/rbac/pkg/json_io"
	"testing"
)

func TestReader1(t *testing.T) {
	documentStr := `{"subjects":[{"id":"A","roles":["R1"]}],"roles":[{"id":"R1","permissions":["edit"]}],"permissions":[{"id":"edit"}]}`
	documentReader := NewDocumentReader(json_io.NewStringReader(documentStr))
	err := documentReader.Accept(&PrintVisitor{})
	if err != nil {
		t.Error(err)
	}
}

func TestWriter1(t *testing.T) {
	writer := json_io.NewStringWriter()
	dw := NewDocumentWriter(writer)
	if err := dw.VisitDocumentStart(); err != nil {
		t.Error(err)
	}

	subjectA := SubjectNode{Id: "A"}
	if err := dw.VisitSubject(&subjectA); err != nil {
		t.Error(err)
	}

	roleR1 := RoleNode{Id: "R1"}
	if err := dw.VisitRole(&roleR1); err != nil {
		t.Error(err)
	}

	permissionEdit := PermissionNode{Id: "edit"}
	if err := dw.VisitPermission(&permissionEdit); err != nil {
		t.Error(err)
	}

	if err := dw.VisitSubjectAssignment(&subjectA, &roleR1); err != nil {
		t.Error(err)
	}

	if err := dw.VisitPermissionAssignment(&roleR1, &permissionEdit); err != nil {
		t.Error(err)
	}

	if err := dw.VisitDocumentEnd(); err != nil {
		t.Error(err)
	}

	documentStr := `{"subjects":[{"id":"A","roles":["R1"]}],"roles":[{"id":"R1","permissions":["edit"]}],"permissions":[{"id":"edit"}]}`

	if *writer.String != documentStr {
		t.Errorf("String not matched %v", *writer.String)
	}
}
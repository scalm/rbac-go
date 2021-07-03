package rbac_simple

import (
	"fmt"
	"github.com/scalm/rbac/pkg/rbac"
)

type PrintVisitor struct {
}

func (v PrintVisitor) VisitDocumentStart() error {
	_, err := fmt.Println("DocumentStart")
	return err
}

func (v PrintVisitor) VisitSubject(subject rbac.Subject) error {
	_, err := fmt.Println("Subject", subject)
	return err
}

func (v PrintVisitor) VisitRole(role rbac.Role) error {
	_, err := fmt.Println("Role", role)
	return err
}

func (v PrintVisitor) VisitPermission(permission rbac.Permission) error {
	_, err := fmt.Println("Permission", permission)
	return err
}

func (v PrintVisitor) VisitSubjectAssignment(sa rbac.SubjectAssignment) error {
	_, err := fmt.Printf("SubjectAssignment {Subject: %v, Role: %v}\n", sa.GetSubjectRef(), sa.GetRoleRef())
	return err
}

func (v PrintVisitor) VisitPermissionAssignment(pa rbac.PermissionAssignment) error {
	_, err := fmt.Printf("PermissionAssignment {Role: %v, Permission: %v}\n", pa.GetRoleRef(), pa.GetPermissionRef())
	return err
}

func (v PrintVisitor) VisitDocumentEnd() error {
	_, err := fmt.Println("DocumentEnd")
	return err
}

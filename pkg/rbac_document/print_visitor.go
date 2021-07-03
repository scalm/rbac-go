package rbac_document

import (
	"fmt"
)

type PrintVisitor struct {
}

func (v PrintVisitor) VisitDocumentStart() error {
	_, err := fmt.Println("DocumentStart")
	return err
}

func (v PrintVisitor) VisitSubject(subject *SubjectNode) error {
	_, err := fmt.Println("Subject", subject)
	return err
}

func (v PrintVisitor) VisitRole(role *RoleNode) error {
	_, err := fmt.Println("Role", role)
	return err
}

func (v PrintVisitor) VisitPermission(permission *PermissionNode) error {
	_, err := fmt.Println("Permission", permission)
	return err
}

func (v PrintVisitor) VisitSubjectAssignment(subject *SubjectNode, role *RoleNode) error {
	_, err := fmt.Printf("SubjectAssignment {Subject: %v, Role: %v}\n", subject, role)
	return err
}

func (v PrintVisitor) VisitPermissionAssignment(role *RoleNode, permission *PermissionNode) error {
	_, err := fmt.Printf("PermissionAssignment {Role: %v, Permission: %v}\n", role, permission)
	return err
}

func (v PrintVisitor) VisitDocumentEnd() error {
	_, err := fmt.Println("DocumentEnd")
	return err
}

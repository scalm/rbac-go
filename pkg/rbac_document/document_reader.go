package rbac_document

import (
	"errors"
	"fmt"
	"github.com/scalm/rbac/pkg/json_io"
)

type DocumentReader struct {
	upstreamReader json_io.AbstractReader
}

func (dr *DocumentReader) Accept(visitor AbstractDocumentVisitor) error {
	var document Document
	if err := dr.upstreamReader.Read(&document); err != nil {
		return err
	}

	return dr.acceptDocument(&document, visitor)
}

func NewDocumentReader(downstreamReader json_io.AbstractReader) *DocumentReader {
	return &DocumentReader{downstreamReader}
}

func (dr *DocumentReader) acceptDocument(document *Document, visitor AbstractDocumentVisitor) error {
	if err := visitor.VisitDocumentStart(); err != nil {
		return err
	}
	subjectsMap := make(map[string]*SubjectNode)
	rolesMap := make(map[string]*RoleNode)
	permMap := make(map[string]*PermissionNode)

	for _, subjectNode := range document.Subjects {
		subjectsMap[subjectNode.Id] = subjectNode
		if err := visitor.VisitSubject(subjectNode); err != nil {
			return err
		}
	}

	for _, roleNode := range document.Roles {
		rolesMap[roleNode.Id] = roleNode
		if err := visitor.VisitRole(roleNode); err != nil {
			return err
		}
	}

	for _, permissionNode := range document.Permissions {
		permMap[permissionNode.Id] = permissionNode
		if err := visitor.VisitPermission(permissionNode); err != nil {
			return err
		}
	}

	for _, subjectNode := range document.Subjects {
		subject := subjectsMap[subjectNode.Id]
		if subject == nil {
			return errors.New(fmt.Sprintf("Subject %v is not defined", subjectNode.Id))
		}
		for _, roleId := range subjectNode.Roles {
			role := rolesMap[roleId]
			if role == nil {
				return errors.New(fmt.Sprintf("Role %v is not defined", roleId))
			}
			if err := visitor.VisitSubjectAssignment(subject, role); err != nil {
				return err
			}
		}
	}

	for _, roleNode := range document.Roles {
		role := rolesMap[roleNode.Id]
		if role == nil {
			return errors.New(fmt.Sprintf("Role %v is not defined", roleNode.Id))
		}
		for _, permissionId := range roleNode.Permissions {
			permission := permMap[permissionId]
			if permission == nil {
				return errors.New(fmt.Sprintf("Permission %v is not defined", permissionId))
			}
			if err := visitor.VisitPermissionAssignment(role, permission); err != nil {
				return err
			}
		}
	}
	if err := visitor.VisitDocumentEnd(); err != nil {
		return err
	}

	return nil
}

package rbac_simple_document

import (
	"github.com/scalm/rbac/pkg/rbac"
	"github.com/scalm/rbac/pkg/rbac_document"
	"github.com/scalm/rbac/pkg/rbac_simple"
)

type VisitorAdapter struct {
	factory            *rbac_simple.Factory
	upstreamVisitor    rbac.AbstractVisitor
	subjectBySource    map[*rbac_document.SubjectNode]rbac.Subject
	rolesBySource      map[*rbac_document.RoleNode]rbac.Role
	permissionBySource map[*rbac_document.PermissionNode]rbac.Permission
}

func NewVisitorAdapter(factory *rbac_simple.Factory, upstreamVisitor rbac.AbstractVisitor) *VisitorAdapter {
	return &VisitorAdapter{factory, upstreamVisitor, nil, nil, nil}
}

func (v *VisitorAdapter) VisitDocumentStart() error {
	v.subjectBySource = make(map[*rbac_document.SubjectNode]rbac.Subject)
	v.rolesBySource = make(map[*rbac_document.RoleNode]rbac.Role)
	v.permissionBySource = make(map[*rbac_document.PermissionNode]rbac.Permission)
	return v.upstreamVisitor.VisitDocumentStart()
}

func (v *VisitorAdapter) VisitSubject(subject *rbac_document.SubjectNode) error {
	rbacSubject := v.factory.CreateSubject(&rbac_simple.CreateSubjectParams{Id: subject.Id})
	v.subjectBySource[subject] = rbacSubject
	return v.upstreamVisitor.VisitSubject(rbacSubject)
}

func (v *VisitorAdapter) VisitRole(role *rbac_document.RoleNode) error {
	rbacRole := v.factory.CreateRole(&rbac_simple.CreateRoleParams{Id: role.Id})
	v.rolesBySource[role] = rbacRole
	return v.upstreamVisitor.VisitRole(rbacRole)
}

func (v *VisitorAdapter) VisitPermission(permission *rbac_document.PermissionNode) error {
	rbacPermission := v.factory.CreatePermission(&rbac_simple.CreatePermissionParams{Id: permission.Id})
	v.permissionBySource[permission] = rbacPermission
	return v.upstreamVisitor.VisitPermission(rbacPermission)
}

func (v *VisitorAdapter) VisitSubjectAssignment(subject *rbac_document.SubjectNode, role *rbac_document.RoleNode) error {
	assignment := v.factory.CreateSubjectAssignment(v.subjectBySource[subject], v.rolesBySource[role])
	return v.upstreamVisitor.VisitSubjectAssignment(assignment)
}

func (v *VisitorAdapter) VisitPermissionAssignment(role *rbac_document.RoleNode, permission *rbac_document.PermissionNode) error {
	assignment := v.factory.CreatePermissionAssignment(v.permissionBySource[permission], v.rolesBySource[role])
	return v.upstreamVisitor.VisitPermissionAssignment(assignment)
}

func (v *VisitorAdapter) VisitDocumentEnd() error {
	v.subjectBySource = nil
	v.rolesBySource = nil
	v.permissionBySource = nil
	return nil
}

package rbac_document

type AbstractDocumentVisitor interface {
	VisitDocumentStart() error
	VisitSubject(subject *SubjectNode) error
	VisitRole(role *RoleNode) error
	VisitPermission(permission *PermissionNode) error
	VisitSubjectAssignment(subject *SubjectNode, role *RoleNode) error
	VisitPermissionAssignment(role *RoleNode, permission *PermissionNode) error
	VisitDocumentEnd() error
}

package rbac_document

import "github.com/scalm/rbac/pkg/json_io"

type DocumentWriter struct {
	upstreamWriter json_io.AbstractWriter
	document *Document
}

func NewDocumentWriter(upstreamWriter json_io.AbstractWriter) *DocumentWriter {
	return &DocumentWriter{upstreamWriter, nil}
}

func (w *DocumentWriter) VisitDocumentStart() error {
	w.document = &Document{}
	return nil
}

func (w *DocumentWriter) VisitSubject(subject *SubjectNode) error {
	w.document.Subjects = append(w.document.Subjects, subject)
	return nil
}

func (w *DocumentWriter) VisitRole(role *RoleNode) error {
	w.document.Roles = append(w.document.Roles, role)
	return nil
}

func (w *DocumentWriter) VisitPermission(permission *PermissionNode) error {
	w.document.Permissions = append(w.document.Permissions, permission)
	return nil
}

func (w *DocumentWriter) VisitSubjectAssignment(subject *SubjectNode, role *RoleNode) error {
	subject.Roles = append(subject.Roles, role.Id)
	return nil
}

func (w *DocumentWriter) VisitPermissionAssignment(role *RoleNode, permission *PermissionNode) error {
	role.Permissions = append(role.Permissions, permission.Id)
	return nil
}

func (w *DocumentWriter) VisitDocumentEnd() error {
	err := w.upstreamWriter.Write(w.document)
	if err != nil {
		return err
	}
	w.document = nil
	return nil
}


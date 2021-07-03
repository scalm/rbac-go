# Package `json_loader`

Interface `AbstractLoader`

Class `ByteALoader` loads an abstract JSON-bytes to an abstract object.

Class `StreamLoader` loads an abstract JSON-stream to an abstract object.

# Package `rbac_document`

Class `DocumentLoader` loads a typed object from a JSON-document in the special format.

Class `Document`, `SubjectNode`, `RoleNode`, `PermissionNode`. This objects presents document.

Class `PrintVisitor`. This visitor prints document nodes to a stdout.

# Package `rbac_simple`

Class `Factory`. The factory for "Simple Storage".

Class `Subject`

Class `Role`

Class `Permission`

Class `PrintVisitor`

# Package `rbac_document_simple`

`VisitorAdapter` adapts the document visitor to the simple visitor.

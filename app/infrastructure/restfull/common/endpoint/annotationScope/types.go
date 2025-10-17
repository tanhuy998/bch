package annotationScope

import "app/infrastructure/restfull/common/endpoint/annotation"

type (
	StructAnnotationConstraint interface {
		annotation.IEndpointAnnotation
		_struct()
	}

	MethodAnnotationConstraint interface {
		annotation.IEndpointAnnotation
		_method()
	}
)

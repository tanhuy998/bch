package mongoStorage

import (
	"reflect"
)

type (
	MongoStorageUnit[Entity_Mapping_Type any] struct {
		MongoDBQueryMonitorCollection
	}
)

func (this *MongoStorageUnit[Entity_Mapping_Type]) GetDBStorageUnitName() string {

	return this.collection.Name()
}

func (this *MongoStorageUnit[Entity_Mapping_Type]) GetEntityMappingType() reflect.Type {

	return reflect.TypeFor[Entity_Mapping_Type]()
}

func (this *MongoStorageUnit[Entity_Mapping_Type]) GetStorageUnit() *MongoDBQueryMonitorCollection {

	return &this.MongoDBQueryMonitorCollection
}

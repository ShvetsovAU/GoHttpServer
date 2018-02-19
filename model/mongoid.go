package model

import (
	"fmt"
	"github.com/shvetsovau/GoHttpServer/constants"
	"gopkg.in/mgo.v2/bson"
)

// Плохо, что модель, которая в нашем случае одновременно (еще и Dto) привязана к конкретной БД
// (тип поля id везде монговский)
// поэтому, когда в БЛ хочется что-то как-то сделать с Id, то автоматически БЛ привязываешь к монге
// Следующие функции - слабая и надеюсь временная попытка хоть как-то отвязать БЛ от конкретной БД

//генерирует новый Id
func GetId() bson.ObjectId {
	return bson.NewObjectId()
}

//преобразует 'что-то' в bson.ObjectId
func ToId(candidate interface{}) bson.ObjectId {
	result, err := TryToId(candidate)
	if err != nil {
		panic(err)
	}
	return result
}

//пробует преобразовать 'что-то' в bson.ObjectId
func TryToId(candidate interface{}) (bson.ObjectId, error) {
	if candidate == nil {
		return "", fmt.Errorf(constants.ErrorWrongId, candidate)
	}

	switch v := candidate.(type) {
	case string:
		if bson.IsObjectIdHex(v) {
			return bson.ObjectIdHex(v), nil
		}
		return "", fmt.Errorf(constants.ErrorWrongId, candidate)
	case bson.ObjectId:
		return v, nil
	default:
		return "", fmt.Errorf(constants.ErrorWrongId, candidate)
	}
}

func IdToString(id interface{}) string {
	switch v := id.(type) {
	case string:
		return v
	case bson.ObjectId:
		return v.Hex()
	default:
		panic(fmt.Errorf(constants.ErrorWrongId, id))
	}
}

func IdIsEqual(id1 interface{}, id2 interface{}) bool {
	return IdToString(id1) == IdToString(id2)
}
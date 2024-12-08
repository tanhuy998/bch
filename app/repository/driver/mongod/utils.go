package mongoRepository

import "go.mongodb.org/mongo-driver/mongo"

func ReturnResultOrNoDocuments[Model_T any](opRes *Model_T, err error) (*Model_T, error) {

	if err == mongo.ErrNoDocuments {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return opRes, nil
}

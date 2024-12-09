package mongoRepositorySorter

import repositoryAPI "app/repository/api"

type (
	MongoSorterGenerator struct {
		cur string
		m   map[string]int
	}
)

func (this *MongoSorterGenerator) Field(name string) repositoryAPI.ISortOperator {

	this.cur = name

	return this
}

func (this *MongoSorterGenerator) Ascending() {

	if this.cur == "" {

		return
	}

	this.m[this.cur] = 1
}

func (this *MongoSorterGenerator) Descending() {

	if this.cur == "" {

		return
	}

	this.m[this.cur] = -1
}

func (this *MongoSorterGenerator) Get() map[string]int {

	return this.m
}

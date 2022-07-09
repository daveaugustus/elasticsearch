package elasticsearch

// CreateIndices to create index 365 Indices based on dates
// Which means index-current date to index-date exactly a year back ie.
// index-2022-07-09 to index-2021-07-09 here the format is taken as YYYY-MM-DD
func (esc ESClient) CreateIndices(indexPrefix string) (interface{}, error) {
	return nil, nil
}

// FillTestData
func (esc ESClient) FillTestData(indexPrefix string) (interface{}, error) {
	return nil, nil
}

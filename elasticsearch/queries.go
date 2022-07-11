package elasticsearch

import "github.com/sirupsen/logrus"

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

// InsertDocToMapped inserts one document
func (esc ESClient) InsertDocToMapped(index, mapping string, doc interface{}) {
	indexResp, err := esc.client.Index().Index(index).Type(index).
		BodyJson(doc).Refresh("true").Do(esc.ctx)
	if err != nil {
		logrus.Errorf("cannot insert the document: %+v", err)
		return
	}
	_ = indexResp

}

// InsertBulkDocsToMapped inserts bulk documents
func (esc ESClient) InsertBulkDocsToMapped(index, mapping string, docs interface{}) {
	indexResp, err := esc.client.Index().Index(index).Type(index).
		BodyJson(docs).Refresh("true").Do(esc.ctx)
	if err != nil {
		logrus.Errorf("cannot insert the document: %+v", err)
		return
	}
	_ = indexResp

}

package elasticsearch

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ESClient struct {
	client *elastic.Client
	ctx    context.Context
}

func NewESClient() ESClient {
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		// TODO:keep URLs and credentials as env variables
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(5*time.Second),
	)
	if err != nil {
		fmt.Printf("cannot create client: %s.\n", err.Error())
	}
	return ESClient{client: client, ctx: context.Background()}
}

// IndexExists to check whether an index exists or not
func (esc ESClient) IndexExists(index string) bool {

	exist, err := esc.client.IndexExists(index).Do(esc.ctx)
	if err != nil {
		logrus.Errorf("ERROR: while checking index exists: %v", err)
		return false
	}
	return exist
}

// CreateIndex to create an index without passing the mapping string
func (esc ESClient) CreateIndex(index string) bool {
	_, err := esc.client.CreateIndex(index).Do(esc.ctx)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v", err)
		return false
	}
	return true
}

// CreateIndex to create an index without passing the mapping string
func (esc ESClient) CreateIndexWithMapping(index string, mapping string) bool {
	exists, err := esc.client.IndexExists(index).Do(esc.ctx)
	if err != nil {
		log.Fatalf("IndexExists ERROR: %v", err)
		return false
	}
	//if index does not exist, create a new one with the specified mapping
	if !exists {
		createIndex, err := esc.client.CreateIndex(index).BodyString(mapping).Do(esc.ctx)
		if err != nil {
			log.Fatalf("IndexExists ERROR: %v", err)
			return false
		}
		if !createIndex.Acknowledged {
			log.Println(createIndex)
		} else {
			log.Println("successfully created index")
		}
	} else {
		log.Println("Index already exist")
		return false
	}
	return true
}

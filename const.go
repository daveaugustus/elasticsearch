package main

// Mapping and index name. Elasticsearch index doctypes now deprecated
const (
	index    = "some_index"
	mappings = `
	{
		"settings": {
			"number_of_shards": 2,
			"number_of_replicas": 1
		},
		"mappings": {
			"properties": {
				"field str": {
					"type": "text"
				},
				"field int": {
					"type": "integer"
				},
				"field bool": {
					"type": "boolean"
				}
			}
		}
	}
	`
)

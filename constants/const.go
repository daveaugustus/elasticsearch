package constants

const (
	YYYYMMDD = "2006-01-02"
)
const (
	PostIndex = "post"
	UserIndex = "user"
	ItemIndex = "item"
)

const UserIndexMapping = `{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"records":{
			"properties":{
				"name":{
					"type":"keyword"
				},
				"age":{
					"type":"keyword"
				},
				"email":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"updated":{
					"type":"date"
				}
			}
		}
	}
}
`
const ItemIndexMapping = `
`

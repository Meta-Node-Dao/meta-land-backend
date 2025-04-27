package mock

import (
	"ceres/pkg/model/tag"
	"ceres/pkg/router"
	"encoding/json"
)

func GetTagsByTagType(c *router.Context) {
	var res tag.ListResponse
	var resString = "{\n    \"page\": 1,\n    \"size\": 20,\n    \"total\": 85,\n    \"list\": [\n        {\n            \"id\": 1,\n            \"name\": \"区块链基础协议\",\n            \"type\": 1\n        },\n        {\n            \"id\": 5,\n            \"name\": \"去中心化存储\",\n            \"type\": 2\n        },\n        {\n            \"id\": 8,\n            \"name\": \"DAO治理\",\n            \"type\": 3\n        }\n    ]\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

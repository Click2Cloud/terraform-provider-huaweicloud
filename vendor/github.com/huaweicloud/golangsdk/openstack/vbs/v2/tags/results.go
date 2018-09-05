package tags

import (
	"github.com/huaweicloud/golangsdk"
)

type Tags struct {
	//contains list of tags, i.e.key value pair
	Tag []ActionTags `json:"tags"`
}

type Tag struct {
	//tag Key
	Key string `json:"key"`
	//tag Value
	Value string `json:"value"`
}

type Resources struct {
	//List of resources
	Resource []Resource `json:"resources"`
	//Total number of resources
	TotalCount int `json:"total_count"`
}

type Resource struct {
	//Resource name
	ResourceName string `json:"resource_name"`
	//Resource ID
	ResourceID string `json:"resource_id"`
	//List of tags
	Tags []Tag `json:"tags"`
}

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

type DeleteResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type QueryResults struct {
	commonResult
}

type ActionResults struct {
	commonResult
}

func (r commonResult) Extract() (*Tags, error) {
	var response Tags
	err := r.ExtractInto(&response)
	return &response, err
}

func (r QueryResults) ExtractResources() (*Resources, error) {
	var response Resources
	err := r.ExtractInto(&response)
	return &response, err
}

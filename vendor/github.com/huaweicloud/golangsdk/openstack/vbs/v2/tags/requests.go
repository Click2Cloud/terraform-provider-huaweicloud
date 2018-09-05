package tags

import (
	"github.com/huaweicloud/golangsdk"
)

type CreateOptsBuilder interface {
	ToTagsCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Tag ActionTags `json:"tag" required:"true"`
}

// ActionTags is a structure of key value pair, used for create ,
// update and delete operations.
type ActionTags struct {
	//tag key
	Key string `json:"key" required:"true"`
	//tag value
	Value string `json:"value" required:"true"`
}

func (opts CreateOpts) ToTagsCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

//adds a tag for a backup policy.
func Create(client *golangsdk.ServiceClient, policyID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToTagsCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client, policyID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

//deletes a tag ,specified in the key for a backup policy.
func Delete(c *golangsdk.ServiceClient, policyID string, key string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, policyID, key), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// Get retrieves the tags of a specific backup policy.
func Get(client *golangsdk.ServiceClient, policyID string) (r GetResult) {
	_, r.Err = client.Get(createURL(client, policyID), &r.Body, nil)
	return
}

type QueryOptsBuilder interface {
	ToTagsQueryMap() (map[string]interface{}, error)
}

type QueryOpts struct {
	//Number of queried records. This parameter is not displayed if action is set to count.
	Limit string `json:"limit,omitempty"`
	//Query index,this parameter is not displayed if action is set to count.
	Offset string `json:"offset,omitempty"`
	//List of tags. Backup policies with these tags will be filtered.
	// This list can have a maximum of 10 tags.
	Tags []QueryTags `json:"tags,omitempty"`
	//Backup policies with any tags in this list will be filtered.
	AnyTags []QueryTags `json:"tags_any,omitempty"`
	//List of excluded tags. Backup policies without these tags will be filtered.
	NotTags []QueryTags `json:"not_tags,omitempty"`
	//List of excluded tags. Backups without any of these tags will be filtered.
	NotAnyTags []QueryTags `json:"not_tags_any,omitempty"`
	//Operator. Possible values are filter and count.
	Action string `json:"action" required:"true"`
}
type QueryTags struct {
	////tag key
	Key string `json:"key" required:"true"`
	//list of values
	Values []string `json:"values" required:"true"`
}

func (opts QueryOpts) ToTagsQueryMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

//retrives a backup policy details using tags.
func Query(client *golangsdk.ServiceClient, opts QueryOptsBuilder) (r QueryResults) {
	b, err := opts.ToTagsQueryMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(queryURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type BatchOptsBuilder interface {
	ToTagsBatchMap() (map[string]interface{}, error)
}

type BatchOpts struct {
	//List of tags to perform batch operation
	Tags []ActionTags `json:"tags,omitempty"`
	//Operator , Possible values are:create, update,delete
	Action ActionType `json:"action" required:"true"`
}

//ActionType specifies the type of batch operation action to be performed
type ActionType string

var (
	// ActionCreate is used to set action operator to create
	ActionCreate ActionType = "create"
	// ActionDelete is used to set action operator to delete
	ActionDelete ActionType = "delete"
	// ActionUpdate is used to set action operator to update
	ActionUpdate ActionType = "update"
)

func (opts BatchOpts) ToTagsBatchMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

//BatchAction is used to create ,update or delete the tags of a specified backup policy.
func BatchAction(client *golangsdk.ServiceClient, policyID string, opts BatchOptsBuilder) (r ActionResults) {
	b, err := opts.ToTagsBatchMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, policyID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

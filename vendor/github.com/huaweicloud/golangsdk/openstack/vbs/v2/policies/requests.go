package policies

import (
	"reflect"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains the options for create a Policy. This object is
// passed to Policy.Create(). For more information about these parameters,
// please refer to the Policy object, or the volume backup service API v2
// documentation
type CreateOpts struct {
	//Backup policy name.It cannot start with default.
	Name string `json:"backup_policy_name" required:"true"`
	//Details about the scheduling policy
	ScheduledPolicy CreateSchedule `json:"scheduled_policy" required:"true"`
	// Tags to be configured for the backup policy
	Tags []Tags `json:"tags,omitempty"`
}

//Details about the scheduling policy
type CreateSchedule struct {
	//Start time of the backup job.
	StartTime string `json:"start_time" required:"true"`
	//Backup interval (1 to 14 days)
	Frequency int `json:"frequency,omitempty"`
	//Number of retained backups, minimum 2.
	RententionNum int `json:"rentention_num,omitempty"`
	//Whether to retain the first backup in the current month, possible values Y or N
	RemainFirstBackup string `json:"remain_first_backup_of_curMonth" required:"true"`
	//Backup policy status, ON or OFF
	Status string `json:"status" required:"true"`
}

type Tags struct {
	//Tag key. A tag key consists of up to 36 characters
	Key string `json:"key" required:"true"`
	//Tag value. A tag value consists of 0 to 43 characters
	Value string `json:"value" required:"true"`
}

// ToPolicyCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create will create a new Policy based on the values in CreateOpts. To extract
// the Policy object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(commonURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains the options for Update a Policy.
type UpdateOpts struct {
	//Backup policy name.It cannot start with default.
	Name string `json:"backup_policy_name,omitempty"`
	//Details about the scheduling policy
	ScheduledPolicy UpdateSchedule `json:"scheduled_policy,omitempty"`
}

type UpdateSchedule struct {
	//Start time of the backup job.
	StartTime string `json:"start_time,omitempty"`
	//Backup interval (1 to 14 days)
	Frequency int `json:"frequency,omitempty"`
	//Number of retained backups, minimum 2.
	RententionNum int `json:"rentention_num,omitempty"`
	//Number of retained backups, minimum 2.
	RemainFirstBackup string `json:"remain_first_backup_of_curMonth,omitempty"`
	//Backup policy status, ON or OFF
	Status string `json:"status,omitempty"`
}

// ToPolicyUpdateMap assembles a request body based on the contents of a
// UpdateOpts.
func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

//Update will Update an existing backup Policy based on the values in CreateOpts.To extract
// the Policy object from the response, call the Extract method on the
// UpdateResult.
func Update(c *golangsdk.ServiceClient, policyID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Put(resourceURL(c, policyID), b, &r.Body, reqOpt)
	return
}

//Delete will delete the specified backup policy
func Delete(c *golangsdk.ServiceClient, policyID string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, policyID), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

//ListOpts allows filtering policies
type ListOpts struct {
	//Backup policy ID
	ID string `json:"backup_policy_id"`
	//Backup policy name
	Name string `json:"backup_policy_name"`
	//Backup policy status
	Status string `json:"status"`
}

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// Policies. It accepts a ListOpts struct, which allows you to
// filter the returned collection for greater efficiency.
func List(c *golangsdk.ServiceClient, opts ListOpts) ([]Policy, error) {

	pages, err := pagination.NewPager(c, commonURL(c), func(r pagination.PageResult) pagination.Page {
		return PolicyPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()

	allPolicies, err := ExtractPolicies(pages)
	if err != nil {
		return nil, err
	}

	return FilterPolicies(allPolicies, opts)
}

func FilterPolicies(policies []Policy, opts ListOpts) ([]Policy, error) {

	var refinedPolicies []Policy
	var matched bool

	m := map[string]FilterStruct{}

	if opts.ID != "" {
		m["ID"] = FilterStruct{Value: opts.ID}
	}
	if opts.Name != "" {
		m["Name"] = FilterStruct{Value: opts.Name}
	}

	if opts.Status != "" {
		m["Status"] = FilterStruct{Value: opts.Status, Driller: []string{"ScheduledPolicy"}}
	}

	if len(m) > 0 && len(policies) > 0 {
		for _, policies := range policies {
			matched = true

			for key, value := range m {
				if sVal := GetStructNestedField(&policies, key, value.Driller); !(sVal == value.Value) {
					matched = false
				}
			}
			if matched {
				refinedPolicies = append(refinedPolicies, policies)
			}
		}
	} else {
		refinedPolicies = policies
	}
	return refinedPolicies, nil
}

func GetStructNestedField(v *Policy, field string, structDriller []string) string {
	r := reflect.ValueOf(v)
	for _, drillField := range structDriller {
		f := reflect.Indirect(r).FieldByName(drillField).Interface()
		r = reflect.ValueOf(f)
	}
	f1 := reflect.Indirect(r).FieldByName(field)
	return string(f1.String())
}

type FilterStruct struct {
	Value   string
	Driller []string
}

// AssociateOptsBuilder allows extensions to add additional parameters to the
// Associate request.
type AssociateOptsBuilder interface {
	ToPolicyAssociateMap() (map[string]interface{}, error)
}

// AssociateOpts contains the options associate a resource to a Policy.
type AssociateOpts struct {
	//Backup policy name.It cannot start with default.
	PolicyID string `json:"backup_policy_id" required:"true"`
	//Details about the scheduling policy
	Resources []AssociateResources `json:"resources" required:"true"`
}

type AssociateResources struct {
	//Backup policy name.It cannot start with default.
	ResourceID string `json:"resource_id" required:"true"`
	//Details about the scheduling policy
	ResourceType string `json:"resource_type" required:"true"`
}

// ToPolicyAssociateMap assembles a request body based on the contents of a
// AssociateOpts.
func (opts AssociateOpts) ToPolicyAssociateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Associate will associate a resource tp a backup policy based on the values in AssociateOpts. To extract
// the associated resources from the response, call the ExtractResource method on the
// ResourceResult.
func Associate(client *golangsdk.ServiceClient, opts AssociateOpts) (r ResourceResult) {
	b, err := opts.ToPolicyAssociateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(associateURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// DisassociateOptsBuilder allows extensions to add additional parameters to the
// Disassociate request.
type DisassociateOptsBuilder interface {
	ToPolicyDisassociateMap() (map[string]interface{}, error)
}

// DisassociateOpts contains the options disassociate a resource from a Policy.
type DisassociateOpts struct {
	//Disassociate Resources
	Resources []DisassociateResources `json:"resources" required:"true"`
}

type DisassociateResources struct {
	//ResourceID
	ResourceID string `json:"resource_id" required:"true"`
}

// ToPolicyDisassociateMap assembles a request body based on the contents of a
// DisassociateOpts.
func (opts DisassociateOpts) ToPolicyDisassociateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Disassociate will disassociate a resource from a backup policy based on the values in DisassociateOpts. To extract
// the disassociated resources from the response, call the ExtractResource method on the
// ResourceResult.
func Disassociate(client *golangsdk.ServiceClient, policyID string, opts DisassociateOpts) (r ResourceResult) {
	b, err := opts.ToPolicyDisassociateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(disassociateURL(client, policyID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

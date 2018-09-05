package tags

import "github.com/huaweicloud/golangsdk"

func createURL(c *golangsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(c.ProjectID, "backuppolicy", policyID, "tags")
}

func deleteURL(c *golangsdk.ServiceClient, policyID string, key string) string {
	return c.ServiceURL(c.ProjectID, "backuppolicy", policyID, "tags", key)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, "backuppolicy", "tags")
}

func queryURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, "backuppolicy", "resource_instances", "action")
}

func actionURL(c *golangsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(c.ProjectID, "backuppolicy", policyID, "tags", "action")
}

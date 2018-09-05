package policies

import "github.com/huaweicloud/golangsdk"

func commonURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, "backuppolicy")
}

func resourceURL(c *golangsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(c.ProjectID, "backuppolicy", policyID)
}

func associateURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, "backuppolicyresources")
}

func disassociateURL(c *golangsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(c.ProjectID, "backuppolicyresources", policyID, "deleted_resources")
}

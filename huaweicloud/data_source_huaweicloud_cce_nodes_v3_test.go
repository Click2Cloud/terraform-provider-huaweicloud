package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCCENodesV3DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCCENode(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCCENodeV3DataSource_node,
			},
			resource.TestStep{
				Config: testAccCCENodeV3DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCCENodeV3DataSourceID("data.huaweicloud_cce_node_v3.nodes"),
					resource.TestCheckResourceAttr("data.huaweicloud_cce_node_v3.nodes", "name", "test-node2"),
					resource.TestCheckResourceAttr("data.huaweicloud_cce_node_v3.nodes", "flavor", "s1.medium"),
				),
			},
		},
	})
}

func testAccCheckCCENodeV3DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find nodes data source: %s ", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Node data source ID not set ")
		}

		return nil
	}
}

var testAccCCENodeV3DataSource_node = fmt.Sprintf(`
resource "huaweicloud_cce_node_v3" "node_1" {
cluster_id = "%s"
  name = "test-node2"
  flavor="s1.medium"
  az= "%s"
  sshkey="KeyPair-c2c"
  iptype="5_bgp"    
  root_volume = {
    size= 40,
    volumetype= "SATA"
  }
  data_volumes = [
    {
      size= 100,
      volumetype= "SATA"
    }]
 
}`, OS_CLUSTER_ID,OS_AVAILABILITY_ZONE)

var testAccCCENodeV3DataSource_basic = fmt.Sprintf(`
data "huaweicloud_cce_node_v3" "nodes" {
		cluster_id ="%s"
		name = "${huaweicloud_cce_node_v3.node_1.name}"
}
`, OS_CLUSTER_ID)


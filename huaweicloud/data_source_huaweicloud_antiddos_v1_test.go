package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAntiDdosV1DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckAntiddos(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAntiDdosV1DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAntiDdosV1DataSourceID("data.huaweicloud_antiddos_v1.antiddos"),
					resource.TestCheckResourceAttr("data.huaweicloud_antiddos_v1.antiddos", "network_type", "EIP"),
					resource.TestCheckResourceAttr("data.huaweicloud_antiddos_v1.antiddos", "status", "normal"),
				),
			},
		},
	})
}

func testAccCheckAntiDdosV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find defense status of EIP data source: %s ", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Defense status of EIP data source ID not set")
		}

		return nil
	}
}

var testAccAntiDdosV1DataSource_basic = fmt.Sprintf(`
data "huaweicloud_antiddos_v1" "antiddos" {  
  floating_ip_id = "%s"
}
`, OS_EIP_ID)

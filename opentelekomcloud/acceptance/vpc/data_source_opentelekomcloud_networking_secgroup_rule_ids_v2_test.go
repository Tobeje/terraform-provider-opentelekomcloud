package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/acceptance/common"
	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/acceptance/common/quotas"
)

func TestAccNetworkingSecGroupRuleIdsV2DataSource_basic(t *testing.T) {
	dataSourceName := "data.opentelekomcloud_networking_secgroup_rule_ids_v2.secgroup_ids"
	t.Parallel()
	th.AssertNoErr(t, quotas.SecurityGroup.Acquire())
	defer quotas.SecurityGroup.Release()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { common.TestAccPreCheck(t) },
		ProviderFactories: common.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingSecGroupIdsV2DataSourceSg,
			},
			{
				Config: testAccNetworkingSecGroupRuleIdsV2DataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "ids.#", "3"),
				),
			},
		},
	})
}

const testAccNetworkingSecGroupIdsV2DataSourceSg = `
resource "opentelekomcloud_networking_secgroup_v2" "secgroup_1" {
  name        = "secgroup_1_sg_ids"
  description = "My neutron security group"
}
`

var testAccNetworkingSecGroupRuleIdsV2DataSourceBasic = fmt.Sprintf(`
%s
data "opentelekomcloud_networking_secgroup_rule_ids_v2" "secgroup_ids" {
  security_group_id = opentelekomcloud_networking_secgroup_v2.secgroup_1.id
}
`, testAccNetworkingSecGroupIdsV2DataSourceSg)

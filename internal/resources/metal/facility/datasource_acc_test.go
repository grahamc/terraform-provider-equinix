package facility

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/equinix/terraform-provider-equinix/internal/tfacc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var (
	matchErrMissingFeature = regexp.MustCompile(`.*doesn't have feature.*`)
	matchErrNoCapacity     = regexp.MustCompile(`Not enough capacity.*`)
)

func TestAccDataSourceMetalFacility_basic(t *testing.T) {
	testFac := "dc13"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { tfacc.PreCheck(t) },
		Providers: tfacc.AccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMetalFacilityConfig_basic(testFac),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.equinix_metal_facility.test", "code", testFac),
				),
			},
			{
				Config: testAccDataSourceMetalFacilityConfig_capacityReasonable(testFac),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.equinix_metal_facility.test", "code", testFac),
				),
			},
			{
				Config:      testAccDataSourceMetalFacilityConfig_capacityUnreasonable(testFac),
				ExpectError: matchErrNoCapacity,
			},
			{
				Config:      testAccDataSourceMetalFacilityConfig_capacityUnreasonableMultiple(testFac),
				ExpectError: matchErrNoCapacity,
			},
		},
	})
}

func TestAccDataSourceMetalFacility_Features(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { tfacc.PreCheck(t) },
		Providers: tfacc.AccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceMetalFacilityConfig_missingFeatures(),
				ExpectError: matchErrMissingFeature,
			},
		},
	})
}

func testAccDataSourceMetalFacilityConfig_missingFeatures() string {
	return `
data "equinix_metal_facility" "test" {
    code = "da11"
    features_required = ["baremetal", "ibx", "foofeature"]
}
`
}

func testAccDataSourceMetalFacilityConfig_basic(facCode string) string {
	return fmt.Sprintf(`
data "equinix_metal_facility" "test" {
    code = "%s"
}
`, facCode)
}

func testAccDataSourceMetalFacilityConfig_capacityUnreasonable(facCode string) string {
	return fmt.Sprintf(`
data "equinix_metal_facility" "test" {
    code = "%s"
    capacity {
        plan = "c3.small.x86"
        quantity = 1000
    }
}
`, facCode)
}

func testAccDataSourceMetalFacilityConfig_capacityReasonable(facCode string) string {
	return fmt.Sprintf(`
data "equinix_metal_facility" "test" {
    code = "%s"
    capacity {
        plan = "c3.small.x86"
        quantity = 1
    }
    capacity {
        plan = "c3.medium.x86"
        quantity = 1
    }
}
`, facCode)
}

func testAccDataSourceMetalFacilityConfig_capacityUnreasonableMultiple(facCode string) string {
	return fmt.Sprintf(`
data "equinix_metal_facility" "test" {
    code = "%s"
    capacity {
        plan = "c3.small.x86"
        quantity = 1
    }
    capacity {
        plan = "c3.medium.x86"
        quantity = 1000
    }
}
`, facCode)
}
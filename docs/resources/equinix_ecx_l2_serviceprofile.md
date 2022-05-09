---
subcategory: "Fabric"
---

# equinix_ecx_l2_serviceprofile (Resource)

Resource `equinix_ecx_l2_serviceprofile` is used to manage layer 2 service profiles
in Equinix Fabric.

This resource relies on the Equinix Fabric API. The parameters
and attributes available map to the fields described at
<https://developer.equinix.com/catalog/sellerv3#operation/getProfileByIdOrNameUsingGET>.

## Example Usage

```hcl
resource "equinix_ecx_l2_serviceprofile" "private-profile" {
  name                               = "private-profile"
  description                        = "my private profile"
  connection_name_label              = "Connection"
  bandwidth_threshold_notifications  = ["John.Doe@example.com", "Marry.Doe@example.com"]
  profile_statuschange_notifications = ["John.Doe@example.com", "Marry.Doe@example.com"]
  vc_statuschange_notifications      = ["John.Doe@example.com", "Marry.Doe@example.com"]
  private                            = true
  private_user_emails                = ["John.Doe@example.com", "Marry.Doe@example.com"]
  features {
    allow_remote_connections = true
    test_profile = false
  }
  port {
    uuid       = "a867f685-422f-22f7-6de0-320a5c00abdd"
    metro_code = "NY"
  }
  port {
    uuid       = "a867f685-4231-2317-6de0-320a5c00abdd"
    metro_code = "NY"
  }
  speed_band {
    speed      = 1000
    speed_unit = "MB"
  }
  speed_band {
    speed      = 500
    speed_unit = "MB"
  }
  speed_band {
    speed      = 100
    speed_unit = "MB"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the service profile. An alpha-numeric 50 characters string which can
include only hyphens and underscores.
* `description` - (Optional) Description of the service profile.
* `bandwidth_alert_threshold` - (Optional) Specifies the port bandwidth threshold percentage. If
the bandwidth limit is met or exceeded, an alert is sent to the seller.
* `speed_customization_allowed` - (Optional) Boolean value that determines if customer is allowed
to enter a custom connection speed.
* `oversubscription_allowed` - (Optional) Boolean value that determines if, regardless of the
utilization, Equinix Fabric will continue to add connections to your links until we reach the
oversubscription limit. By selecting this service, you acknowledge that you will manage decisions
on when to increase capacity on these link.
* `api_integration` - (Optional) Boolean value that determines if API integration is enabled. It
allows you to complete connection provisioning in less than five minutes. Without API Integration,
additional manual steps will be required and the provisioning will likely take longer.
* `authkey_label` - (Optional) Name of the authentication key label to be used by the
Authentication Key service. It allows Service Providers with QinQ ports to accept groups of
connections or VLANs from Dot1q customers. This is similar to S-Tag/C-Tag capabilities.
* `connection_name_label` - (Optional) Custom name used for calling a connections
e.g. `circuit`. Defaults to `Connection`.
* `ctag_label` - (Optional) C-Tag/Inner-Tag label name for the connections.
* `servicekey_autogenerated` - (Optional) Boolean value that indicates whether multiple connections
can be created with the same authorization key to connect to this service profile after the first
connection has been approved by the seller.
* `equinix_managed_port_vlan` - (Optional) Applicable when `api_integration` is set to `true`. It
indicates whether the port and VLAN details are managed by Equinix.
* `integration_id` - (Optional) Specifies the API integration ID that was provided to the customer
during onboarding. You can validate your API integration ID using the validateIntegrationId API.
* `bandwidth_threshold_notifications` - (Optional) A list of email addresses that will receive
notifications about bandwidth thresholds.
* `profile_statuschange_notifications` - (Required) A list of email addresses that will receive
notifications about profile status changes.
* `vc_statuschange_notifications` - (Required) A list of email addresses that will receive
notifications about connections approvals and rejections.
* `oversubscription` - (Optional) You can set an alert for when a percentage of your profile has
been sold. Service providers like to use this functionality to alert them when they need to add
more ports or when they need to create a new service profile. Required with
`oversubscription_allowed`, defaults to `1x`.
* `private` - (Optional) Boolean value that indicates whether or not this is a private profile,
i.e. not public like AWS/Azure/Oracle/Google, etc. If private, it can only be available for
creating connections if correct permissions are granted.
* `private_user_emails` - (Optional) An array of users email ids who have permission to access this
service profile. Argument is required when profile is set as private.
* `redundancy_required` - (Optional) Boolean value that determines if your connections will require
redundancy. if yes, then users need to create a secondary redundant connection.
* `speed_from_api` - (Optional) Boolean valuta that determines if connection speed will be derived
from an API call. Argument has to be specified when `api_integration` is enabled.
* `tag_type` - (Optional) Specifies additional tagging information required by the seller profile
for Dot1Q to QinQ translation. See [Enhance Dot1q to QinQ translation support](https://docs.equinix.com/es/Content/Interconnection/Fabric/layer-2/Fabric-Create-Layer2-Service-Profile.htm#:~:text=Enhance%20Dot1q%20to%20QinQ%20translation%20support)
for additional information. Valid values are:
  * `CTAGED` - When seller side VLAN C-Tag has to be provided _(Default)_.
  * `NAMED` - When application named tag has to be provided.
  * `BOTH` - When both, application tag or seller side VLAN C-Tag can be provided.
* `secondary_vlan_from_primary` - (Optional) Indicates whether the VLAN ID of. the secondary
connection is the same as the primary connection.
* `features` - (Required) Block of profile features configuration. See [Features](#features) below
for more details.
* `port` - (Required) One or more definitions of ports residing in locations, from which your
customers will be able to access services using this service profile. See [Port](#port) below for
more details.
* `speed_band` - (Optional) One or more definitions of supported speed/bandwidth. Argument is
required when `speed_from_api` is set to `false`. See [Speed Band](#speed-band) below for more
details.

### Features

The `features` block has below fields:

* `allow_remote_connections` - (Required) Indicates whether or not connections to this profile
can be created from remote metro locations.
* `test_profile` - (Deprecated) Indicates whether or not this profile can be used for test
connections.

### Port

Each `port` block has below fields:

* `uuid` - (Required) Unique identifier of the port.
* `metro_code` - (Required) The metro code of location where the port resides.

### Speed Band

Each `speed_band` block has below fields:

* `speed` - (Required) Speed/bandwidth supported by this service profile.
* `speed_unit` - (Required) Unit of the speed/bandwidth supported by this service profile. One of
`MB`, `GB`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` - Unique identifier of the service profile.
* `state` - Service profile provisioning status.

## Import

This resource can be imported using an existing ID:

```sh
terraform import equinix_ecx_l2_serviceprofile.example {existing_id}
```
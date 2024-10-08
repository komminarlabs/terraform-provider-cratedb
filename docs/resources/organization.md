---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cratedb_organization Resource - terraform-provider-cratedb"
subcategory: ""
description: |-
  Creates and manages an organization.
---

# cratedb_organization (Resource)

Creates and manages an organization.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the organization.

### Read-Only

- `dc` (Attributes) The DublinCore of the organization. (see [below for nested schema](#nestedatt--dc))
- `email` (String) The notification email used in the organization.
- `id` (String) The id of the organization.
- `notifications_enabled` (Boolean) Whether notifications enabled for the organization.
- `plan_type` (Number) The support plan type used in the organization.
- `project_count` (Number) The project count in the organization.
- `role_fqn` (String) The role FQN.

<a id="nestedatt--dc"></a>
### Nested Schema for `dc`

Read-Only:

- `created` (String) The created time.
- `modified` (String) The modified time.

---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dockerhub_repositorygroup Resource - terraform-provider-dockerhub"
subcategory: ""
description: |-
  Manages an organization group / repository permission binding
---

# dockerhub_repositorygroup (Resource)

Manages an organization group / repository permission binding

## Example Usage

```terraform
resource "dockerhub_repositorygroup" "example" {
  repository = "organisation/project"
  group      = 123
  groupname  = "groupname"
  permission = "read"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group` (Number) The group to add.
- `groupname` (String) The group to add.
- `repository` (String) The repository path.

### Optional

- `permission` (String) The permission to assign the group. One of 'read', 'write' and 'admin'.

### Read-Only

- `id` (String) The provider specific id of the repository group.

## Import

Import is supported using the following syntax:

```shell
# import using the organization/registry/groupname
terraform import dockerhub_repositorygroup example organization/registry/groupname
```
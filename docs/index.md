---
layout: ""
page_title: "Provider: dockerhub"
description: |-
  Register hub.docker.io repositories.
---

# dockerhub Provider

Register hub.docker.io repositories.

## Example Usage

```terraform
terraform {
  required_providers {
    dockerhub = {
      source = "BarnabyShearer/dockerhub"
    }
  }
}

provider "dockerhub" {
  # Note this cannot be a Personal Access Token
  username = "USERNAME" # optionally use DOCKER_USERNAME env var
  password = "PASSWORD" # optionally use DOCKER_PASSWORD env var
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **password** (String) Password for authentication.
- **username** (String) Username for authentication.
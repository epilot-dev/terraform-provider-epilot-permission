terraform {
  required_providers {
    epilot-permission = {
      source  = "epilot-dev/epilot-permission"
      version = "0.0.3"
    }
  }
}

provider "epilot-permission" {
  epilot_auth = var.epilot_api_key
}

variable "epilot_api_key" {
  type = string
}

resource "epilot-permission_role" "new_admin" {
  role = {
    user_role = {
      name            = "File Admin_test 2.."
      id              = "15218142:file_admin_test"
      slug            = "file_admin_test"
      type            = "user_role"
      organization_id = "15218142"
      grants = [
        {
          action   = "entity:*"
          resource = "file*"
        }
      ]
    }
  }
}

terraform {
  required_providers {
    epilotpermissions = {
      source  = "epilot/epilotpermissions"
      version = "0.1.0"
    }
  }
}

provider "epilotpermissions" {
  authorization = var.epilot_api_key
}

variable "epilot_api_key" {
  type = string
}

resource "epilotpermissions_role" "new_admin" {
  role = {
    user_role = {
      name            = "File Admin_test"
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

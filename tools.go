// +build tools

package main

import (
	// analysistest does not work with Go Modules yet
	_ "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	_ "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	_ "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	_ "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

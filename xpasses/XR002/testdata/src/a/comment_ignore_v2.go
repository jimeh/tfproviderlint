package a

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fcommentignore_v2() {
	//lintignore:XR002
	_ = schema.Resource{
		Create: createFunc_v2,
	}
}

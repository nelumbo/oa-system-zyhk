package gin

import (
	"oa-backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GinArrayPreprocessing(c *gin.Context) (xForms models.XForms) {
	var pageSizeErr, pageNoErr error
	xForms.PageSize, pageSizeErr = strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	xForms.PageNo, pageNoErr = strconv.Atoi(c.DefaultQuery("pageNo", "0"))

	if pageSizeErr != nil || xForms.PageSize <= 0 {
		xForms.PageSize = 10
	}
	if pageNoErr != nil || xForms.PageNo <= 0 {
		xForms.PageNo = 1
	}
	return
}

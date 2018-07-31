package v1

import (
	"estate/middleware"
	"estate/models/v1"
	"strconv"

	"github.com/gin-gonic/gin"
)

var baseModel = new(v1.BaseModel)

// 本部中介-路由
func Base(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/base")
	router.Use(middleware.Auth())
	router.GET("/date_list", Base_DateList)                                   // 7.1 本部中介-日期列表
	router.GET("/sales_achievement", Base_SalesAchievement)                   // 7.2 本部中介-销售业绩
	router.GET("/sales_profit/list", Base_SalesProfitList)                    // 7.3.1 本部中介-中介费用统计-列表
	router.GET("/sales_profit/detail", Base_SalesProfitDetail)                // 7.3.2 本部中介-中介费用统计-详情
	router.GET("/sales_profit/setting_detail", Base_SalesProfitSettingDetail) // 7.3.3 本部中介-中介费用统计-设置详情
	router.GET("/sales_profit/setting_modify", Base_SalesProfitSettingModify) // 7.3.4 本部中介-中介费用统计-设置修改
}

// 本部中介-日期列表
func Base_DateList(c *gin.Context) {
	var data interface{}
	data, errMsg := baseModel.Base_DateList()
	if errMsg != "" {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  errMsg,
		})
		return
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
	return
}

// 本部中介-销售业绩
func Base_SalesAchievement(c *gin.Context) {
	addTime := string(c.Query("add_time"))
	perPage, _ := strconv.Atoi(c.Query("per_page"))
	lastId, _ := strconv.Atoi(c.Query("last_id"))
	groupId, _ := strconv.Atoi(c.Request.Header.Get("group_id"))
	if addTime == "" || groupId == 0 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "参数错误",
		})
		return
	}
	if groupId != 1 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "非本部中介不允许查看销售业绩",
		})
		return
	}

	// 销售业绩列表
	var data interface{}
	data, errMsg := baseModel.Base_SalesAchievement(addTime, perPage, lastId)
	if errMsg != "" {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  errMsg,
		})
		return
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
	return
}

// 本部中介-中介费用统计列表
func Base_SalesProfitList(c *gin.Context) {
	addTime := string(c.Query("add_time"))
	perPage, _ := strconv.Atoi(c.Query("per_page"))
	lastId, _ := strconv.Atoi(c.Query("last_id"))
	groupId, _ := strconv.Atoi(c.Request.Header.Get("group_id"))
	userType, _ := strconv.Atoi(c.Request.Header.Get("user_type"))
	if addTime == "" || groupId == 0 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "参数错误",
		})
		return
	}
	if groupId != 1 && userType != 1 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "非本部主管不允许查看中介费用统计",
		})
		return
	}

	// 中介费用统计列表
	var data interface{}
	data, errMsg := baseModel.Base_SalesProfitList(addTime, perPage, lastId)
	if errMsg != "" {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  errMsg,
		})
		return
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
	return
}

// 本部中介-中介费用统计详情
func Base_SalesProfitDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	groupId, _ := strconv.Atoi(c.Request.Header.Get("group_id"))
	userType, _ := strconv.Atoi(c.Request.Header.Get("user_type"))
	if groupId == 0 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "参数错误",
		})
		return
	}
	if groupId != 1 && userType != 1 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "非本部主管不允许查看中介费用统计详情",
		})
		return
	}

	// 中介费用统计详情
	data, errMsg := baseModel.Base_SalesProfitDetail(id)
	if errMsg != "" {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  errMsg,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
	return
}

// 本部中介-中介费用统计设置详情
func Base_SalesProfitSettingDetail(c *gin.Context) {
	estateId, _ := strconv.Atoi(c.Query("estate_id"))
	groupId, _ := strconv.Atoi(c.Request.Header.Get("group_id"))
	userType, _ := strconv.Atoi(c.Request.Header.Get("user_type"))
	if estateId == 0 || groupId == 0 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "参数错误",
		})
		return
	}
	if groupId != 1 && userType != 1 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "非本部主管不允许查看中介费用统计设置详情",
		})
		return
	}

	// 中介费用统计详情
	data, errMsg := baseModel.Base_SalesProfitSettingDetail(estateId)
	if errMsg != "" {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  errMsg,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
	return
}

// 本部中介-中介费用统计设置修改
func Base_SalesProfitSettingModify(c *gin.Context) {
	estateId, _ := strconv.Atoi(c.Query("estate_id"))
	agencyJson := c.Query("agency_json")
	groupId, _ := strconv.Atoi(c.Request.Header.Get("group_id"))
	userType, _ := strconv.Atoi(c.Request.Header.Get("user_type"))
	if estateId == 0 || agencyJson == "" || groupId == 0 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "参数错误",
		})
		return
	}
	if groupId != 1 && userType != 1 {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  "非本部主管不允许修改中介费用统计设置详情",
		})
		return
	}

	// 中介费用修改
	errMsg := baseModel.Base_SalesProfitSettingModify(estateId, agencyJson)
	if errMsg != "" {
		c.JSON(400, gin.H{
			"code": 1010,
			"msg":  errMsg,
		})
		return
	}
	c.JSON(201, gin.H{
		"code": 0,
		"msg":  "success",
	})
	return
}
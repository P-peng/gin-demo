package apis

import (
	orm "gin-demo/api/database"
	model "gin-demo/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 列表数据
func Users(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

// 条件查询
func GetUser(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.GetUser(user.Username)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

// 添加数据
func Insert(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")

	// 开启事务
	tx := orm.Eloquent.Begin()
	id, err := user.Insert(user)
	if err != nil {
		// 事务回滚
		tx.Rollback()
	} else {
		// 提交事务
		tx.Commit()
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    id,
	})
}

// 修改数据
func Update(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	value := c.Request.FormValue("id")

	int64, err := strconv.ParseInt(value, 10, 64)
	user.ID = int64
	result, err := user.Update(user)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}

//删除数据
func Delete(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Request.FormValue("id"), 10, 64)
	var rows int64
	// 开启事务
	tx := orm.Eloquent.Begin()
	rows, err = user.Destroy(id)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
		"data":    rows,
	})
}

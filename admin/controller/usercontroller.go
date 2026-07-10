package controller

import (
	"bytes"
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"github.com/demdxx/gocast"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	Services service.User
	BaseController
}

type RowUser struct {
	Id              int64           `json:"id" form:"id" gorm:"column:id"`
	Username        string          `json:"username" form:"username" gorm:"column:username"`
	Email           string          `json:"email" form:"email" gorm:"column:email"`
	DateJoined      time.Time       `json:"date_joined" form:"date_joined" gorm:"column:date_joined"`
	Phone           string          `json:"phone" form:"phone" gorm:"column:phone"`
	AreaCode        string          `json:"area_code" form:"area_code" gorm:"column:area_code"`
	AgentId         int64           `json:"agent_id" form:"agent_id" gorm:"column:agent_id"`
	AgentInviteId   int64           `json:"agent_invite_id" form:"agent_invite_id" gorm:"column:agent_invite_id"`
	ParentId        int64           `json:"parent_id" form:"parent_id" gorm:"column:parent_id"`
	InviteCode      string          `json:"invite_code" form:"invite_code" gorm:"column:invite_code"`
	InviteCount     int64           `json:"invite_count" form:"invite_count" gorm:"column:invite_count"`
	Status          int             `json:"status" form:"status" gorm:"column:status"`
	AgentAccount    string          `json:"agent_account" form:"agent_account" gorm:"column:agent_account"`
	ParentUsername  string          `json:"parent_username" form:"parent_username" gorm:"column:parent_username"`
	AgentInviteCode string          `json:"agent_invite_code" form:"agent_invite_code" gorm:"column:agent_invite_code"`
	Level           string          `json:"level" form:"level" gorm:"column:level"`
	IsTest          string          `json:"is_test" form:"is_test" gorm:"column:is_test"`
	IsAuth          *int            `json:"is_auth" form:"is_auth" `
	IsActive        *int            `json:"is_active" form:"is_active" `
	Exp             decimal.Decimal `json:"exp" form:"exp" gorm:"column:exp"`
}

func (this *UserController) Index(ctx *gin.Context) {

	db := global.SHOP_DB

	type Request struct {
		request.PageInfo
		Count          int64  `json:"count" form:"count"`
		Username       string `json:"username" form:"username"`
		ParentUsername string `json:"parent_username" form:"parent_username"`
		AgentAccount   string `json:"agent_account" form:"agent_account"`
		Email          string `json:"email" form:"email"`
		Phone          string `json:"phone" form:"phone"`
		Export         bool   `json:"export" form:"export"`
		IsTest         int    `json:"is_test" form:"is_test"`
	}

	req := Request{
		PageInfo: request.PageInfo{
			Limit: 20,
			Page:  1,
		},
	}
	err := ctx.ShouldBind(&req)
	if err != nil {
		this.ErrorHtml(ctx, err.Error())
		return
	}

	req.Offset = (req.Page - 1) * req.Limit
	query := db.
		Select(`
			a.id, a.username,a.exp, a.email, a.date_joined, a.is_auth,a.is_active,a.invite_code,a.invite_count,
			a.phone, a.area_code, 
			c.agent_id, c.agent_invite_id, c.parent_id, c.status, 
			d.acc as agent_account, 
			e.username as parent_username, 
			f.invite_code as agent_invite_code ,
            g.title as level
		`).
		Table("auth_user as a").
		//Joins("left join auth_user_extends as b on a.id = b.user_id").
		Joins("left join account_team as c on a.id = c.user_id").
		Joins("left join agent_info as d on c.agent_id = d.id").
		Joins("left join auth_user as e on c.parent_id = e.id").
		Joins("left join agent_invite as f on c.agent_invite_id = f.id").
		//Joins("left join account_user_authority as g on g.username = a.username").
		Joins("left join member_level as g on g.level =a.level")

	if req.Username != "" {
		query = query.Where("a.username like ?", "%"+req.Username+"%")
	}

	if req.ParentUsername != "" {
		query = query.Where("e.username like ?", "%"+req.ParentUsername+"%")
	}

	if req.Email != "" {
		query = query.Where("a.email like ?", "%"+req.Email+"%")
	}

	if req.Phone != "" {
		query = query.Where("b.phone like ?", "%"+req.Phone+"%")
	}
	if req.IsTest != 0 {
		query = query.Where("a.is_test = ?", req.IsTest)
	}

	if req.Export {
		f := excelize.NewFile()
		defer func() {
			f.Close()
		}()
		sheet1 := "Sheet1"
		_, err := f.NewSheet(sheet1)
		if err != nil {
			ctx.Error(err)
			return
		}

		row := 1
		col := 1
		_ = f.SetCellValue(sheet1, BuildCellPoint(row, col), "序号")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+1, col), "账号")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+2, col), "Email")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+3, col), "手机号")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+4, col), "状态")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+5, col), "注册时间")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+6, col), "代理商信息")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+7, col), "上级信息")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+8, col), "邀请码")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+9, col), "邀请人数")
		statusEnum := map[int]string{0: "注册", 1: "实名"}
		list := make([]RowUser, 0)
		query.Order("a.id desc").FindInBatches(&list, 1000, func(db *gorm.DB, batch int) error {
			for _, info := range list {
				col += 1
				_ = f.SetCellValue(sheet1, BuildCellPoint(row, col), info.Id)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+1, col), info.Username)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+2, col), info.Email)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+3, col), info.Phone)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+4, col), statusEnum[info.Status])
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+5, col), info.DateJoined.Format(time.DateTime))
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+6, col), info.AgentAccount)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+7, col), info.ParentUsername)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+8, col), info.InviteCode)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+9, col), info.InviteCount)
			}
			return nil
		})
		f.Path = fmt.Sprintf("用户管理-%s.xlsx", time.Now().Format("20060102150405"))
		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Path))
		ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		var buffer bytes.Buffer
		_ = f.Write(&buffer)
		http.ServeContent(ctx.Writer, ctx.Request, f.Path, time.Now(), bytes.NewReader(buffer.Bytes()))
		return
	}

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		this.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]RowUser, 0)
	if err = query.
		Order("a.id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&list).Error; err != nil {
		this.ErrorHtml(ctx, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "user_index.html", gin.H{
		"status": "200",
		"list":   list,
		"req":    req,
	})
}
func (this *UserController) Changepwd(ctx *gin.Context) {
	method := ctx.Request.Method

	if method == "GET" {
		var adminUser model.User

		if err := ctx.ShouldBind(&adminUser); err == nil {

			fmt.Printf("login-request:%+v\n", adminUser)
			global.SHOP_DB.First(&adminUser)
			var language []model.Language
			global.SHOP_DB.Model(model.Language{}).Find(&language)
			var memberlevel []model.MemberLevel
			global.SHOP_DB.Model(model.MemberLevel{}).Where("is_display=1").Find(&memberlevel)

			ctx.HTML(http.StatusOK, "user_changepwd.html", gin.H{

				"admininfo": adminUser,

				"IsUpdate":    true,
				"memberlevel": memberlevel,

				"language": language,
			})
		} else {
			this.Error(ctx, "查询错误", err.Error())
		}
	} else {
		var adminUser model.User
		if err := ctx.ShouldBindJSON(&adminUser); err == nil {

			fmt.Printf("login-request:%+v\n", adminUser)
			//global.SHOP_DB.Where("username = ?", adminUser.Username).First(&adminUser)
			//id := adminUser.Id
			if adminUser.Password == "" {
				err = this.Services.Changepwd(adminUser)
				this.Success(ctx, "修改成功", adminUser)
				return
			}

			if adminUser.Password == "" {
				this.Error(ctx, "修改失败,密码为空")
				return
			}
			err := this.Services.Changepwd(adminUser)
			if err != nil {

				this.Error(ctx, "失败", err.Error())
				return
			}
			//service.RunPublisher(ctx, "shop_message", "修改"+fmt.Sprintf("%v", id)+"成功")
			this.Success(ctx, "修改成功", adminUser)
		} else {
			fmt.Println(err.Error())
			this.Error(ctx, "修改失败", err.Error())
		}
	}

}
func (this *UserController) Changetradepwd(ctx *gin.Context) {

	var adminUser model.User
	if err := ctx.ShouldBind(&adminUser); err == nil {

		//global.SHOP_DB.Where("username = ?", adminUser.Username).First(&adminUser)
		//id := adminUser.Id
		adminUser.TradePassword = strconv.Itoa(rand.Intn(900000) + 100000)
		err := this.Services.Changetradepwd(adminUser)
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			this.Error(ctx, "失败", err.Error())
			return
		}
		//service.RunPublisher(ctx, "shop_message", "修改"+fmt.Sprintf("%v", id)+"成功")
		this.Success(ctx, fmt.Sprintf("修改成功,密码为%s", adminUser.TradePassword), adminUser)
	} else {
		global.SHOP_LOG.Error(err.Error())
		this.Error(ctx, "修改失败", err.Error())
	}

}
func (this *UserController) SetAgent(ctx *gin.Context) {
	db := global.SHOP_DB
	type Request struct {
		UserId          int64  `json:"user_id" form:"user_id" gorm:"column:user_id"`
		AgentAccount    string `json:"agent_account" form:"agent_account"`
		AgentInviteCode string `json:"agent_invite_code" form:"agent_invite_code"`
	}
	req := &Request{}
	err := ctx.ShouldBind(req)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	agentInfo := model.AgentInfo{}
	agentInvite := model.AgentInvite{}
	if req.AgentAccount != "" {

		err = db.Where("acc = ?", req.AgentAccount).First(&agentInfo).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			this.Error(ctx, err.Error())
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			this.Error(ctx, "代理商未找到")
			return
		}

		err = db.Where("agent_id = ?", agentInfo.Id).First(&agentInvite).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			this.Error(ctx, err.Error())
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			this.Error(ctx, "代理商邀请码未找到")
			return
		}

	}

	if req.AgentInviteCode != "" {
		err = db.Raw("select * from agent_invite where invite_code = ?", req.AgentInviteCode).First(&agentInvite).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			this.Error(ctx, err.Error())
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			this.Error(ctx, "代理商邀请码未找到")
			return
		}

		err = db.Raw("select * from agent_info where id = ?", agentInvite.AgentId).First(&agentInfo).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			this.Error(ctx, err.Error())
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			this.Error(ctx, "代理商未找到")
			return
		}

	}

	if ctx.Request.Method == "GET" {
		this.Success(ctx, "查询成功", gin.H{
			"agent_info":   agentInfo,
			"agent_invite": agentInvite,
		})
		return
	}

	updateUserChilds := func(db *gorm.DB, userId int64, agentId int64, agentInviteId int64) error {
		userIds := []int64{userId}
		for {
			err = db.Raw("select user_id from account_team where parent_id in (?)", userIds).Find(&userIds).Error
			if err != nil {
				return err
			}

			if len(userIds) == 0 {
				break
			}

			if len(userIds) > 0 {
				err = db.Exec("update account_team set agent_id = ?, agent_invite_id = ? where user_id in (?)", agentId, agentInviteId, userIds).Error
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	updateAgent := func(db *gorm.DB, agentId int64, agentInviteId int64) error {
		inviteCount := 0
		err = db.Raw("select count(1) from account_team where agent_id = ? and parent_id = 0", agentId).Find(&inviteCount).Error
		if err != nil {
			return err
		}

		teamCount := 0
		err = db.Raw("select count(1) from account_team where agent_id = ?", agentId).Find(&teamCount).Error
		if err != nil {
			return err
		}

		agentInviteCodeInviteCount := 0
		err = db.Raw("select count(1) from account_team where agent_invite_id = ? and parent_id = 0", agentInviteId).Find(&agentInviteCodeInviteCount).Error
		if err != nil {
			return err
		}

		agentInviteCodeTeamCount := 0
		err = db.Raw("select count(1) from account_team where agent_invite_id = ?", agentInviteId).Find(&agentInviteCodeTeamCount).Error
		if err != nil {
			return err
		}

		err = db.Exec("update agent_info set invite_count = ?, team_count = ? where id = ?", inviteCount, teamCount, agentId).Error
		if err != nil {
			return err
		}

		db.Exec("update agent_invite set invite_count = ?, team_count = ? where id = ?", agentInviteCodeInviteCount, agentInviteCodeTeamCount, agentInviteId)
		if err != nil {
			return err
		}

		return nil
	}
	if ctx.Request.Method == "POST" {
		err = db.Transaction(func(tx *gorm.DB) error {
			accountTeam := &model.AccountTeam{}
			err = tx.Where("user_id = ?", req.UserId).First(accountTeam).Error
			if err != nil {
				return err
			}

			originAgentId := accountTeam.AgentId
			originAgentInviteCodeId := accountTeam.AgentInviteId

			accountTeam.AgentId = gocast.ToInt64(agentInfo.Id)
			accountTeam.AgentInviteId = agentInvite.Id

			err = tx.Save(accountTeam).Error
			if err != nil {
				return err
			}

			// 修改用户所有下级所属团队信息
			if err = updateUserChilds(tx, accountTeam.UserId, accountTeam.AgentId, accountTeam.AgentInviteId); err != nil {
				return err
			}

			// 如果有原始的代理商
			if originAgentId != 0 {
				if err = updateAgent(tx, originAgentId, originAgentInviteCodeId); err != nil {
					return err
				}
			}

			if err = updateAgent(tx, accountTeam.AgentId, accountTeam.AgentInviteId); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			this.Error(ctx, err.Error())
			return
		}

		this.Success(ctx, "操作成功")
		return
	}
}

package apps

import (
	"winstore/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// SelectPage 分页查询软件列表
func (a *App) SelectPage(pageNum int, pageSize int, classify string) string {
	applist := new(model.AppList)
	page, err := applist.SelectAllPage(pageNum, pageSize, classify)
	runtime.LogDebugf(a.ctx, "%v", page)
	if err != nil {
		baseModel := model.BaseModel{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
		return baseModel.New()
	}
	baseModel := model.BaseModel{
		Code:    0,
		Message: "success",
		Data:    page,
	}
	return baseModel.New()
}

// SelectByID 根据ID查询软件详情
func (a *App) SelectByID(id string) string {
	data := struct {
		AppList *model.AppList  `json:"appList"`
		Img     *[]model.AppImg `json:"img"`
	}{}

	applist := new(model.AppList)
	err := applist.SelectByID(id)
	if err != nil {
		baseModel := model.BaseModel{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
		return baseModel.New()
	}
	appImg := new(model.AppImg)
	appImg.AppID = id
	img, err := appImg.GetImgByAppID()
	if err != nil {
		baseModel := model.BaseModel{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
		return baseModel.New()
	}
	data.AppList = applist
	data.Img = img
	baseModel := model.BaseModel{
		Code:    0,
		Message: "success",
		Data:    data,
	}
	return baseModel.New()
}

package apps

import (
	"winstore/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// SelectPage 分页查询软件列表
func (a *App) SelectPage(pageNum int, pageSize int, classify string) model.BaseModel {
	applist := new(model.AppList)
	page, err := applist.SelectAllPage(pageNum, pageSize, classify)
	data := struct {
		Count int64            `json:"count"`
		List  *[]model.AppList `json:"list"`
	}{}
	count, err := applist.Count(classify)
	if err != nil {
		return model.BaseModel{}
	}
	data.Count = count
	data.List = page
	runtime.LogDebugf(a.ctx, "%v", page)
	if err != nil {
		return model.BaseModel{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return model.BaseModel{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

// SelectByID 根据ID查询软件详情
func (a *App) SelectByID(id string) model.BaseModel {
	data := struct {
		AppList *model.AppList  `json:"appList"`
		Img     *[]model.AppImg `json:"img"`
	}{}

	applist := new(model.AppList)
	err := applist.SelectByID(id)
	if err != nil {
		return model.BaseModel{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}
	// 查询软件图片
	appImg := new(model.AppImg)
	appImg.AppID = id
	img, err := appImg.GetImgByAppID()
	if err != nil {
		return model.BaseModel{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}
	data.AppList = applist
	data.Img = img
	return model.BaseModel{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

// SelectByName 根据名称查询软件
func (a *App) SelectByName(name string) model.BaseModel {
	applist := new(model.AppList)
	applist.AppName = name
	appLists, err := applist.SelectByAppName()
	if err != nil {
		return model.BaseModel{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return model.BaseModel{
		Code:    0,
		Message: "success",
		Data:    appLists,
	}
}

// SelectClassify 查询分类
func (a *App) SelectClassify() model.BaseModel {
	classify := new(model.Classify)
	classifyList, err := classify.SelectAll()
	if err != nil {
		return model.BaseModel{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return model.BaseModel{
		Code:    0,
		Message: "success",
		Data:    classifyList,
	}
}

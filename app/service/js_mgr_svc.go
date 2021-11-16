package service

import (
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/exception"
	"strconv"
	"strings"
	"sync"

	"gorm.io/gorm"
)

var (
	jsmServiceInstance JsmService
	jsmOnce            sync.Once
)

type jsmServiceImpl struct {
	db         *gorm.DB
	repo       repositories.JsmRepo
}

func GetJsmService() JsmService {
	jsmOnce.Do(func() {
		jsmServiceInstance = &jsmServiceImpl{
			db:         database.GetDriver(),
			repo:       repositories.GetJsmRepo(),
		}
	})
	return jsmServiceInstance
}

type JsmService interface {
	Create(openID string, param *vo.JsManageReq) exception.Exception
	Get(id uint) (*vo.JsManageResp, exception.Exception)
	ListByCategoryID(page *vo.PageInfo, pid uint) (*vo.DataPagination, exception.Exception)
	Update(openID string, id uint, param *vo.JsManageUpdateReq) exception.Exception
	Delete(id uint) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func (jsi *jsmServiceImpl) Create(openID string, param *vo.JsManageReq) exception.Exception {
	jsmMgr := param.ToModel(openID)
	return jsi.repo.Create(jsi.db, jsmMgr)
}

func (jsi *jsmServiceImpl) Get(id uint) (*vo.JsManageResp, exception.Exception) {
	jsm, ex := jsi.repo.Get(jsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewJsManageResponse(jsm), nil
}

func (jsi *jsmServiceImpl) ListByCategoryID(pageInfo *vo.PageInfo, pid uint) (*vo.DataPagination, exception.Exception) {
	count, jsms, ex := jsi.repo.ListByCategoryID(jsi.db, pageInfo, pid)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.JsManageResp, 0, len(jsms))
	for i := range jsms {
		resp = append(resp, *vo.NewJsManageResponse(&jsms[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (jsi *jsmServiceImpl) Update(openID string, id uint, param *vo.JsManageUpdateReq) exception.Exception {
	return jsi.repo.Update(jsi.db, id, param.ToMap(openID))
}

func (jsi *jsmServiceImpl) Delete(id uint) exception.Exception {
	return jsi.repo.Delete(jsi.db, id)
}

func (jsi *jsmServiceImpl) MultiDelete(ids string) exception.Exception {
	idslice := strings.Split(ids, ",")
	if len(idslice) == 0 {
		return exception.New(response.ExceptionInvalidRequestParameters, "无效参数")
	}
	jid := make([]uint, 0, len(idslice))
	for i := range idslice {
		id, err := strconv.ParseUint(idslice[i], 10, 0)
		if err != nil {
			return exception.Wrap(response.ExceptionParseStringToUintError, err)
		}
		jid = append(jid, uint(id))
	}
	return jsi.repo.MultiDelete(jsi.db, jid)
}

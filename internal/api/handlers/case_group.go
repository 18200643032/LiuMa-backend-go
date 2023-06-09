package handlers

import "github.com/gin-gonic/gin"

type caseGroupStruct struct {
}
type CaseGroupInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Query(c *gin.Context)
}

func (c2 *caseGroupStruct) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c2 *caseGroupStruct) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c2 *caseGroupStruct) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c2 *caseGroupStruct) Query(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

var _ CaseGroupInterface = (*caseGroupStruct)(nil)

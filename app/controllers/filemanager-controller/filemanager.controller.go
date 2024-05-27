package filemanagercontroller

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	filemanagermodel "github.com/gptverse/init/app/models/filemanager-model"
	filemanagerservice "github.com/gptverse/init/app/service/filemanager-service"
	"github.com/gptverse/init/app/utils"
	"github.com/gptverse/init/framework/http/responser"
	"github.com/gptverse/init/framework/translator"
)

type FileManagerController struct {
	File    filemanagermodel.FileModel
	Service filemanagerservice.FileManagerService
}

// @Tags			Files
// @Summary		Upload file
// @Description	Upload file
// @ID				upload-file
// @Accept			json
// @Security		BearerAuth
// @Produce		json
// @Success		200				{object}	filemanagermodel.FileModel
// @Param			file			formData	file	true	"File to upload"
// @Param			Accept-Language	header		string	false	"Language preference"
//
// @Router			/api/v1/file-manager/upload [post]
func (fController *FileManagerController) Upload(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	newFileName := uuid.New().String() + filepath.Ext(fController.File.File.Filename)
	destPath := filepath.Join(fController.Service.Fw.Configs.STORAGE, newFileName)
	currentUserInfo := utils.GetCurrentUserInformations(ginCtx)

	err := ginCtx.SaveUploadedFile(fController.File.File, destPath)
	if err != nil {

		response.SetMessage(translator.GetMessage(ginCtx, "file_connot_upload")).BadWithAbort()
		return
	}

	fileUrl := fmt.Sprintf("%s/%s/%s", fController.Service.Fw.Configs.BASE_URL, fController.Service.Fw.Configs.STORAGE, newFileName)

	// payload := filemanagermodel.FileManagerUploadeFileResponse{
	// 	Url: fileUrl,
	// }
	file := filemanagermodel.FileDatabaseModel{
		File:     fController.File,
		UserId:   currentUserInfo.Id,
		Url:      fileUrl,
		CreateAt: time.Now(),
	}
	fController.Service.Save(file)

	response.Payload(file).SetMessage(translator.GetMessage(ginCtx, "file_uploaded_successfuly")).Success()
}

// @Tags			Files
// @Summary		List All Records
// @Description	Get user details by id
// @ID				fetch-all-files
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	usermodel.UserWithoutPasswordModel
// @Param			page			query		string	true	"page number"	int
// @Param			limit			query		string	true	"limit number"	int
// @Param			Accept-Language	header		string	false	"Language preference"
//
// @Router			/api/v1/file-manager/fetch-all [GET]
func (fController *FileManagerController) FetchAll(ginCtx *gin.Context) {}

// @Tags			Files
// @Summary		Delete user
// @Description	Deletes a user by given user id
// @ID				delete-file
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	translator.TranslationSwaggerResponse
// @Param			request			body		usermodel.UserDeleteModel	true	"query params"
// @Param			Accept-Language	header		string						false	"Language preference"
//
// @Router			/api/v1/file-manager/delete [delete]
func (fController *FileManagerController) Delete(ginCtx *gin.Context) {}

package filemanagerservice

import (
	"context"
	"fmt"

	filemanagermodel "github.com/gptverse/init/app/models/filemanager-model"
	"github.com/gptverse/init/framework"
)

// check auth.route
type FileManagerService struct {
	Collection string // user_auth
	Fw         *framework.Framework
}

func (Srv *FileManagerService) Save(file filemanagermodel.FileDatabaseModel) bool {
	ctx := context.Background()
	fmt.Println("---file---", file)
	coll := Srv.Fw.Database.Instance.Collection(Srv.Collection)
	_, err := coll.InsertOne(ctx, file)
	return err == nil
}

func (Srv *FileManagerService) Delete()   {}
func (Srv *FileManagerService) FetchAll() {}

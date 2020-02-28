package export

import (
	logfile "github.com/chunpat/go-gin-example/pkg/file"
	"github.com/chunpat/go-gin-example/pkg/setting"
	"os"
)

func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath // export/
}

func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath() // runtime/
}

//获取excel系统保存路径
func  GetPwdFullPath(filename string) (string, error){
	//mkdir
	dir, err := os.Getwd()
	if err != nil {
		return "os.Getwd err: %v", err
	}
	systemFullPath := dir + "/" + GetExcelFullPath()

	perm := logfile.CheckPermission(systemFullPath)
	if perm == true {
		//err
		return "",err
	}

	err = logfile.IsNotExistMkDir(systemFullPath)
	if err != nil {
		return "",err
	}

	return GetExcelFullPath() + filename , nil
}
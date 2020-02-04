package export

import "github.com/FromChinaBoy/go-gin-example/pkg/setting"

func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath // export/
}

func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath() // runtime/
}

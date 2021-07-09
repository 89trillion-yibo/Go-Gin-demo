package pflag

import (
	"strings"

	"github.com/spf13/pflag"
)

var cliFileName = pflag.StringP("filename","a","just for test","Input errFileName")

func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

func Identify() string {
	// 设置标准化参数名称的函数
	pflag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	// 为 filename 参数设置 NoOptDefVal
	pflag.Lookup("filename").NoOptDefVal = "nil"

	// 把 badflag 参数标记为即将废弃的，请用户使用 des-detail 参数
	pflag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")
	// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 des-detail 的 shorthand 参数
	pflag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")

	// 在帮助文档中隐藏参数 gender
	pflag.CommandLine.MarkHidden("badflag")

	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()

	return *cliFileName
}
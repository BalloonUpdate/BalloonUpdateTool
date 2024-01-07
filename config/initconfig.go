package config

import (
	"BalloonUpdateTool/checkTools"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitAliyunConfig(ConfigDir string) {
	// 创建配置文件
	viper.SetConfigName("Aliyun")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigDir)
	// 添加默认配置项目
	viper.SetDefault("AliyunAccessKey", "AccessKeyAsDefault")
	viper.SetDefault("AliyunAccessSecret", "AccessSecretAsDefault")
	viper.SetDefault("AliyunRegionId", "RegionIdAsDefault")
	viper.SetDefault("AliyunOSSBucketName", "BucketNameAsDefault")
	viper.SetDefault("AliyunOSSBucketPath", "/BucketPathAsDefault")
	viper.SetDefault("AliyunOSSBucketUrl", "OSS://BucketUrlAsDefault")
	// 写入配置文件
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("配置文件写入失败，请检查权限。")
		panic(err)
	}
}
func InitTencentCloudConfig(ConfigDir string) {
	// 创建配置文件
	viper.SetConfigName("TencentCloud")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigDir)
	// 添加默认配置项目
	viper.SetDefault("TencentCloudSecretId", "SecretIdAsDefault")
	viper.SetDefault("TencentCloudSecretKey", "SecretKeyAsDefault")
	viper.SetDefault("TencentCloudRegion", "RegionAsDefault")
	viper.SetDefault("TencentCloudBucketName", "BucketNameAsDefault")
	viper.SetDefault("TencentCloudBucketPath", "/BucketPathAsDefault")
	viper.SetDefault("TencentCloudBucketUrl", "COS://BucketUrlAsDefault")
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("配置文件写入失败，请检查权限。")
		panic(err)
	}
}
func InitWorkspaceConfig(ConfigDir string) string {
	// 创建配置文件
	viper.SetConfigName("Workspace")
	viper.SetConfigType("yaml")
	// 探测操作系统，设置默认路径
	var WorkspacePath = checkTools.SetConfigDir()
	err, _ := os.Stat(ConfigDir)
	if err != nil {
		fmt.Println("配置文件夹不存在，正在创建。")
		err := os.MkdirAll(ConfigDir, 0755)
		if err != nil {
			fmt.Println("配置文件夹创建失败，请检查权限。")
			panic(err)
		}
	}
	viper.AddConfigPath(ConfigDir)
	viper.New()
	// 添加默认配置项目
	viper.SetDefault("WorkspacePath", WorkspacePath)
	viper.SetDefault("Public", WorkspacePath+"public")
	viper.SetDefault("MPCWorkspace", WorkspacePath+"Workspace")
	viper.SetDefault("History", WorkspacePath+"History")
	// 写入配置文件
	_ = viper.WriteConfig()
	return ""
}

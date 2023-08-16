// Package cmd
/*
this code use apache license 2.0
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "管理配置文件，用于保存路径和其他信息。",
	Long: `用法：
	BalloonUpdateTool config subcommand [flags]
	例如：
		BalloonUpdateTool config aliyun set --access-key-id=xxx --access-key-secret=xxx
		BalloonUpdateTool config WorkspacePath get`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("具体用法请查阅help命令。")
		fmt.Println("\t" + "子命令列表：")
		for _, SubCommand := range cmd.Commands() {
			fmt.Println("\t" + SubCommand.Use + "\t" + SubCommand.Short)
		}
	},
}
var configSetCmdForAliyun = &cobra.Command{
	Use:   "aliyun",
	Short: "设置阿里云的登陆凭证，OSS的配置，以及其他配置。",
	Long: `用法：
	BalloonUpdateTool config Aliyun set --flag1=value1 --flag2=value2
	例如：
		BalloonUpdateTool config aliyun set --BucketName=xxx --Endpoint=xxx`,
}
var configSetCmdForTencentCloud = &cobra.Command{
	Use:   "TencentCloud",
	Short: "设置腾讯云的登陆凭证，COS的配置，以及其他配置。",
	Long: `用法：
	BalloonUpdateTool config TencentCloud set --flag1=value1 --flag2=value2
	例如：
		BalloonUpdateTool config TencentCloud set --BucketName=xxx --Endpoint=xxx`,
}
var configSetCmdForWorkspace = &cobra.Command{
	Use:   "Workspace",
	Short: "设置工作目录选项。",
	Long: `用法：
	BalloonUpdateTool config Workspace set --flag1=value1 --flag2=value2
	例如：
		BalloonUpdateTool config Workspace set --WorkspaceLocation=C:\Users\example\Workspace`,
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configSetCmdForAliyun)
	configCmd.AddCommand(configSetCmdForTencentCloud)
	configCmd.AddCommand(configSetCmdForWorkspace)
}

// Package cmd is the entry point for the command-line application. It defines the root command and adds all the child commands to it. It also sets flags appropriately.
/*
this code use apache license 2.0
*/
package cmd

import (
	"BalloonUpdateTool/checkTools"
	"BalloonUpdateTool/config"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	workspaceLocation        string
	createConfigForWorkspace = &cobra.Command{
		Use:   "create",
		Short: "创建工作目录。",
		Long: `用法：
	BalloonUpdateTool config Workspace create --workspaceLocation=D:\Data\Workspace
	例如：
		BalloonUpdateTool config Workspace create --WorkspaceLocation=C:\Users\example\Workspace`,
		Run: func(cmd *cobra.Command, args []string) {
			if workspaceLocation == "" {
				fmt.Println("在默认工作目录创建配置文件...")
				err := config.InitWorkspaceConfig(checkTools.SetConfigDir())
				if err != "" {
					panic(err)
				}
			} else {
				fmt.Println("在指定位置创建配置文件：", workspaceLocation, "workspace.yml")
				err := config.InitWorkspaceConfig(workspaceLocation)
				if err != "" {
					panic(err)
				}
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "配置文件创建成功！\n")
			}
		},
	}
)

func init() {
	configCmd.AddCommand(configSetCmdForWorkspace)
	configSetCmdForWorkspace.AddCommand(createConfigForWorkspace)
	createConfigForWorkspace.Flags().StringVarP(&workspaceLocation, "workspaceLocation", "w", workspaceLocation, "工作目录的位置。")
}

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wwfyde/demo-cli/cmd/kayn/author"
	"github.com/wwfyde/demo-cli/cmd/kayn/uuid"
	"github.com/wwfyde/demo-cli/cmd/kayn/version"
	"os"
)

var (
	Verbose bool
	Source  string
)

var rootCmd = &cobra.Command{
	Use:   "kayn",
	Short: "kayn is a simple app",
	// TODO(wwfyde) add detail function introduction here
	Long:             "list version",
	Version:          version.Version,
	TraverseChildren: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Persistent flag, available for all command as well as every command under that command
	// 持久命令标志对该命令和该命令的所有子命令可用
	//
	// Global flag
	// 对于全局命令, 为root命令设置持久flag即可.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	//rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	//rootCmd.PersistentFlags().StringP("author", "a", "wwfyde", "Author name for copyright attribution")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	_ = rootCmd.MarkFlagRequired("verbose")

	// Local flags
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")

	_ = viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	_ = viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	_ = viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "wwfyde <wwfyde@gmail.com>")
	viper.SetDefault("license", "apache")

	// TODO add additional command here
	rootCmd.AddCommand(version.NewCmdVersion())
	rootCmd.AddCommand(uuid.NewCmdUuid())
	rootCmd.AddCommand(author.NewCmdAuthor())
}

func main() {
	Execute()
}

func initConfig() {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

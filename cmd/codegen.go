/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/linjunhao1997/sweet/codegen"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"time"
)

const (
	gormColumn = "gorm_column"
)

// codegenCmd represents the codegen command
var codegenCmd = &cobra.Command{
	Use:   "codegen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		defer func() {
			s.Stop()
			fmt.Println("generate gorm column file successful.")
		}()
		pathsVal, _ := cmd.Flags().GetString(gormColumn)
		if len(pathsVal) > 0 {
			paths := strings.Split(pathsVal, " ")
			for _, path := range paths {
				split := strings.Split(path, "/")
				path = split[len(split)-1]
				err := codegen.GenerateGormColumn(strings.TrimSpace(path))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	codegenCmd.Flags().StringP(gormColumn, "", "", "go文件路径,多个文件使用空格")
	rootCmd.AddCommand(codegenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codegenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codegenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

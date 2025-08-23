/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-tool/config"
	"cli-tool/store"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		store_task := store.FileTaskStore{Path: config.TaskFile}
		title, _ := cmd.Flags().GetString("title")
		done, _ := cmd.Flags().GetBool("done")

		list, err := store_task.LoadTask()

		if err != nil {
			return err
		}

		id := len(list) + 1

		list = append(list, store.Task{ID: id, Title: title, Done: done})

		err = store_task.SaveTask(list)

		if err != nil {
			return err
		}

		fmt.Println("Tarea agregada")

		return nil
	},
}

func init() {
	addCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		title, _ := cmd.Flags().GetString("title")
		if len(title) < 3 {
			return fmt.Errorf("el titulo debe ser mayor")
		}

		return nil
	}
	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().String("title", "", "Titulo para mi tarea")
	addCmd.MarkFlagRequired("title")
	addCmd.Flags().Bool("done", false, "Marca si la tarea esta completa")

}

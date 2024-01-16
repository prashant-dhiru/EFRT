/*
Copyright © 2024 Prashant Dhirendra prashant.dhiru@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/prashant-dhiru/efrt/internal/jira"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "log effort to jira",
	Long:  `log opens the interactive shell to log effort to jira based on your config`,
	Run:   runLogCmd,
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	logCmd.Flags().BoolP("comment", "c", false, "add comment to your jira log")
}

func runLogCmd(cmd *cobra.Command, args []string) {
	//fetch the list of the
	taskResp := jira.GetAllActiveTask()

	var task_promt_txt []string
	for _, issue := range taskResp.Issues {
		task_promt_txt = append(task_promt_txt, string(issue.Key)+": "+string(issue.Fields.Summary))
	}

	task_promt := promptui.Select{
		Label: "Select Issue",
		Items: task_promt_txt,
	}

	task_selected_index, _, err := task_promt.Run()
	if err != nil {
		fmt.Println("input error while selecting JIRA issue.")
		os.Exit(1)
	}

	validate_efforts := func(input string) error {
		matched, _ := regexp.MatchString(`^[0-9]{1,}[dmh]{1}$`, input)
		if !matched {
			return errors.New("invalid effort format")
		}
		return nil
	}
	effort_promt := promptui.Prompt{
		Label:       "Time spent (m/h/d)",
		HideEntered: false,
		Validate:    validate_efforts,
		Pointer:     promptui.PipeCursor,
	}
	efforts_string, err := effort_promt.Run()
	if err != nil {
		fmt.Println("input error while getting efforts for the JIRA issue.")
		os.Exit(1)
	}

	isCommentSet, err := cmd.Flags().GetBool("comment")
	if err != nil {
		fmt.Println("input error while getting comments for the JIRA issue.")
		os.Exit(1)
	}
	var comment string = ""
	if isCommentSet {
		fmt.Printf("Comment : ")
		reader := bufio.NewReader(os.Stdin)
		comment, _ = reader.ReadString('\n')
	}
	// fmt.Println(taskResp.Issues[task_selected_index].Key, efforts_string, comment)
	jira.LogEfforts(taskResp.Issues[task_selected_index].Key, efforts_string, comment)
}

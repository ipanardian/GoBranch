/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2018 ipanardian <https://github.com/ipanardian>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package core

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey"
)

//Answers struct for answer
type Answers struct {
	branchType     string
	baseBranchName string
	branchName     string
}

//getBranchTypeQuestion Question for type of branch
func getBranchTypeQuestion(branchType BranchType) []*survey.Question {
	return []*survey.Question{
		{
			Name: "branchType",
			Prompt: &survey.Select{
				Message: "Please choose type of branch:",
				Options: []string{
					branchType.Feature,
					branchType.Enhance,
					branchType.Bugfix,
					branchType.Hotfix,
					branchType.Test,
					branchType.Release,
					branchType.Custom,
				},
				Default: "Feature",
			},
			Transform: TransformInput,
		},
	}
}

//getCustomBaseBranchQuestion Question for base branch
func getCustomBaseBranchQuestion() []*survey.Question {
	return []*survey.Question{
		{
			Name:      "baseBranchName",
			Prompt:    &survey.Input{Message: "What's your BASE branch name? e.g master"},
			Validate:  survey.Required,
			Transform: TransformInput,
		},
	}
}

//getBranchNameQuestion Question for branch name
func getBranchNameQuestion() []*survey.Question {
	return []*survey.Question{
		{
			Name:      "branchName",
			Prompt:    &survey.Input{Message: "What's your branch name?"},
			Validate:  survey.Required,
			Transform: TransformInput,
		},
	}
}

//createQuestions Create the questions
func createQuestions() Answers {
	answers := Answers{}

	branchTypeQuest := getBranchTypeQuestion(branchType)
	err := survey.Ask(branchTypeQuest, &answers.branchType)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if answers.branchType == branchType.Custom {
		baseBranchQuest := getCustomBaseBranchQuestion()
		err2 := survey.Ask(baseBranchQuest, &answers.baseBranchName)
		if err2 != nil {
			fmt.Println(err2.Error())
			os.Exit(1)
		}
		branchTypeMap[branchType.Custom] = answers.baseBranchName
	}

	branchNameQuest := getBranchNameQuestion()
	err3 := survey.Ask(branchNameQuest, &answers.branchName)
	if err3 != nil {
		fmt.Println(err3.Error())
		os.Exit(1)
	}

	return answers
}

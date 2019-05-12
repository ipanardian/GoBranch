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
	"strings"

	"github.com/urfave/cli"
)

//BranchType List predefined of BranchType
type BranchType struct {
	Feature    string
	Enhance    string
	Bugfix     string
	Hotfix     string
	Hotfeature string
	Test       string
	Custom     string
}

//BranchTypeMap Collection of base branch
type BranchTypeMap map[string]string

//TypeConvention Type convention
type TypeConvention struct {
	ForwardSlash string
	Hyphen       string
	Underscore   string
}

//NamingConvention Naming convention
type NamingConvention struct {
	ToSnake string
	ToKebab string
}

//typeConvention values
var typeConvention = TypeConvention{
	"/",
	"-",
	"_",
}

//Default TC
var typeConventionSelected = typeConvention.ForwardSlash

//namingConvention values
var namingConvention = NamingConvention{
	"snake",
	"kebab",
}

//Default NC
var namingConventionSelected = namingConvention.ToSnake

//branchType Init branchType
var branchType = BranchType{
	"feature",
	"enhance",
	"bugfix",
	"hotfix",
	"hotfeature",
	"test",
	"custom",
}

//branchTypeMap Init branchTypeMap
var branchTypeMap = BranchTypeMap{
	branchType.Feature:    "development",
	branchType.Enhance:    "development",
	branchType.Bugfix:     "development",
	branchType.Hotfix:     "hotfix",
	branchType.Hotfeature: "hotfix",
	branchType.Test:       "development",
	branchType.Custom:     "",
}

//createBranchName Transform branch name
func createBranchName(answer Answers) string {
	if len(answer.branchName) == 0 {
		fmt.Println("error: Branch name can not be empty.")
		os.Exit(1)
	}

	isCustom := false
	if answer.branchType == branchType.Custom {
		isCustom = true
	}

	newBranchName := ToValidBranchName(answer.branchName, namingConventionSelected, isCustom)
	if answer.branchType != branchType.Custom {
		newBranchName = strings.Join([]string{answer.branchType, newBranchName}, typeConventionSelected)
	}
	return newBranchName
}

//createBranch Start create git branch
func createBranch(answer Answers) string {
	newBranchName := createBranchName(answer)

	checkoutBaseBranch(branchTypeMap[answer.branchType])
	pullOriginBranch(branchTypeMap[answer.branchType])
	checkoutNewBranch(newBranchName)

	return newBranchName
}

//flagActions set actions for flag
func flagActions(c *cli.Context) {
	tcFlag := c.String("tc")
	ncFlag := c.String("nc")
	if len(tcFlag) > 0 {
		if tcFlag == typeConvention.ForwardSlash {
			typeConventionSelected = typeConvention.ForwardSlash
		} else if tcFlag == typeConvention.Hyphen {
			typeConventionSelected = typeConvention.Hyphen
		} else if tcFlag == typeConvention.Underscore {
			typeConventionSelected = typeConvention.Underscore
		} else {
			fmt.Println("Invalid value of type convention. Available \"/\", \"-\", & \"_\" ")
			os.Exit(1)
		}
	}

	if len(ncFlag) > 0 {
		if ncFlag == namingConvention.ToSnake {
			namingConventionSelected = namingConvention.ToSnake
		} else if ncFlag == namingConvention.ToKebab {
			namingConventionSelected = namingConvention.ToKebab
		} else {
			fmt.Println("Invalid value of naming convention. Available \"snake\" & \"kebab\" ")
			os.Exit(1)
		}
	}
}

//execute create question and branch
func execute(c *cli.Context) {
	flagActions(c)
	answers := createQuestions()
	newBranchName := createBranch(answers)

	fmt.Printf("Here you go, %s is ready!\n", newBranchName)
}

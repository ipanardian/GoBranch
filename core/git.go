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
	"os/exec"
)

const git = "git"

//checkoutBaseBranch Check out to base branch
func checkoutBaseBranch(baseBranchName string) {
	args := []string{"checkout", baseBranchName}
	out, err := executeGit(args)
	if err != nil {
		fmt.Println("error: Base branch " + baseBranchName + " did not exist.")
		os.Exit(1)
	}
	printOutput(out)
}

//pullOriginBranch Pull origin base branch
func pullOriginBranch(baseBranchName string) {
	args := []string{"pull", "origin", baseBranchName}
	out, err := executeGit(args)
	if err != nil {
		fmt.Println("warning: git pull origin " + baseBranchName + " was failed.")
	} else {
		printOutput(out)
	}
}

//checkoutNewBranch Checkout to new branch
func checkoutNewBranch(newBranchName string) {
	args := []string{"checkout", "-b", newBranchName}
	out, err := executeGit(args)
	if err != nil {
		fmt.Println("error: Could not create branch " + newBranchName + ". Please try again.")
		os.Exit(1)
	}
	printOutput(out)
}

//CheckGitPath Check git command is available
func CheckGitPath() string {
	binary, lookErr := exec.LookPath(git)
	if lookErr != nil {
		fmt.Println("git command is not available in your system.")
		os.Exit(1)
	}

	return binary
}

//executeGit Execute git command
func executeGit(args []string) ([]byte, error) {
	CheckGitPath()

	cmd := exec.Command(git, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}
	return out, nil
}

//printOutput print output from git
func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("%s\n", string(outs))
	}
}

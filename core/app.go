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
	"github.com/urfave/cli"
)

//InitializeApp Init App
func InitializeApp(app *cli.App) {
	app.Name = "GoBranch"
	app.Usage = "A command line app to help you quickly creating git branch without hassle step."
	app.UsageText = "Just type gobranch"
	app.Version = "0.1.0.0"
	app.Author = "Ipan Ardian <https://github.com/ipanardian>"
}

//SetFlags Set the flags
func SetFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "tc",
			Value: "",
			Usage: "Set type convention ('/', '-', '_'). e.g feature/{branch}, feature-{branch} or feature_{branch}",
		},
		cli.StringFlag{
			Name:  "nc",
			Value: "",
			Usage: "Set naming convention ('snake', 'kebab'). e.g branch_name or branch-name",
		},
	}
}

//SetActions Set the actions
func SetActions(app *cli.App) {
	app.Action = func(c *cli.Context) error {
		execute(c)
		return nil
	}
}

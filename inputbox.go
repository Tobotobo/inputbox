// inputbox.go
// Copyright (c) 2021 Tobotobo
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// +build windows

package inputbox

import (
	ps "github.com/Tobotobo/powershell"
)

func Show(prompt_title_default ...string) string {
	var prompt, title, defaultAnswer string

	if l := len(prompt_title_default); l >= 1 {
		prompt = prompt_title_default[0]
		if l >= 2 {
			title = prompt_title_default[1]
			if l >= 3 {
				defaultAnswer = prompt_title_default[2]
			}
		}
	}

	if title == "" {
		title = " "
	}

	out, err := ps.Execute(`
		[void][Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic')
		$title = '` + title + `'
		$msg = '` + prompt + `'
		$default = '` + defaultAnswer + `'
		$answer = [Microsoft.VisualBasic.Interaction]::InputBox($msg, $title, $default)
		Write-Output $answer
		`)
	if err != nil {
		panic(err)
	}
	return out
}

```
[
    { "keys": ["ctrl+y"], "command": "run_macro_file", "args": {"file": "res://Packages/Default/Delete Line.sublime-macro"} },
	{ "keys": ["alt+/"], "command": "auto_complete" },
            { "keys": ["alt+m"], "command": "move_to", "args": {"to": "eol", "extend": false} },
            { "keys": ["alt+n"], "command": "move_to", "args": {"to": "brackets"} },
            { "keys": ["alt+j"], "command": "move", "args": {"by": "characters", "forward": false} },
            { "keys": ["alt+l"], "command": "move", "args": {"by": "characters", "forward": true} },
            { "keys": ["alt+i"], "command": "move", "args": {"by": "lines", "forward": false} },
            { "keys": ["alt+k"], "command": "move", "args": {"by": "lines", "forward": true} },
	{ "keys": ["alt+/"], "command": "replace_completion_with_auto_complete", "context":
		[
			{ "key": "last_command", "operator": "equal", "operand": "insert_best_completion" },
			{ "key": "auto_complete_visible", "operator": "equal", "operand": false },
			{ "key": "setting.tab_completion", "operator": "equal", "operand": true }
		]
	},
    // HTML, XML close tag
    { "keys": ["/"], "command": "close_tag", "args": { "insert_slash": true }, "context":
        [
            { "key": "selector", "operator": "equal", "operand": "(text.html, text.xml) - string - comment", "match_all": true },
            { "key": "preceding_text", "operator": "regex_match", "operand": ".*<$", "match_all": true },
            { "key": "setting.auto_close_tags" }
        ]
    }
]
```

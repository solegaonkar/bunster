package parser_test

import "github.com/yassinebenaid/bunny/ast"

var groupingTests = []testCase{
	{`{ cmd; }`, ast.Script{
		Statements: []ast.Statement{
			ast.Group{
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`{ cmd; cmd; }`, ast.Script{
		Statements: []ast.Statement{
			ast.Group{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	{`{cmd;cmd;}`, ast.Script{
		Statements: []ast.Statement{
			ast.Group{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},

	{`{
		cmd
	 	cmd
	}`, ast.Script{
		Statements: []ast.Statement{
			ast.Group{
				ast.Command{Name: ast.Word("cmd")},
				ast.Command{Name: ast.Word("cmd")},
			},
		},
	}},
	{`{cmd&cmd&}`, ast.Script{
		Statements: []ast.Statement{
			ast.Group{
				ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
				ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
			},
		},
	}},
	{`{cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3;}`, ast.Script{
		Statements: []ast.Statement{
			ast.Group{
				ast.BinaryConstruction{
					Left: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd1")}},
						{Command: ast.Command{Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
				ast.BinaryConstruction{
					Left: ast.Pipeline{
						{Command: ast.Command{Name: ast.Word("cmd1")}},
						{Command: ast.Command{Name: ast.Word("cmd2")}},
					},
					Operator: "&&",
					Right:    ast.Command{Name: ast.Word("cmd3")},
				},
			},
		},
	}},

	{`{cmd; cmd;} | {cmd; cmd;}&& {cmd; cmd;}`, ast.Script{
		Statements: []ast.Statement{
			ast.BinaryConstruction{
				Left: ast.Pipeline{
					{Command: ast.Group{
						ast.Command{Name: ast.Word("cmd")},
						ast.Command{Name: ast.Word("cmd")},
					}},
					{Command: ast.Group{
						ast.Command{Name: ast.Word("cmd")},
						ast.Command{Name: ast.Word("cmd")},
					}},
				},
				Operator: "&&",
				Right: ast.Group{
					ast.Command{Name: ast.Word("cmd")},
					ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`{cmd};}`, ast.Script{
		Statements: []ast.Statement{
			ast.Group{
				ast.Command{Name: ast.Word("cmd}")},
			},
		},
	}},

	// {`( cmd )`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.SubShell{
	// 			Body:ast.Statement{
	// 			ast.Command{Name: ast.Word("cmd")},
	// 		},
	// 	},
	// }},
	// {`( cmd; cmd )`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.SubShell{
	// 			Body:ast.Statement{
	// 			ast.Command{Name: ast.Word("cmd")},
	// 			ast.Command{Name: ast.Word("cmd")},
	// 		},
	// 	},
	// }},
	// {`( cmd; cmd; )`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.SubShell{
	// 			Body:ast.Statement{
	// 			ast.Command{Name: ast.Word("cmd")},
	// 			ast.Command{Name: ast.Word("cmd")},
	// 		},
	// 	},
	// }},

	// {`(cmd;cmd)`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.SubShell{
	// 			Body:ast.Statement{
	// 			ast.Command{Name: ast.Word("cmd")},
	// 			ast.Command{Name: ast.Word("cmd")},
	// 		},
	// 	},
	// }},

	// {`(
	// 	cmd
	//  	cmd
	// )`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.SubShell{
	// 			Body:ast.Statement{
	// 			ast.Command{Name: ast.Word("cmd")},
	// 			ast.Command{Name: ast.Word("cmd")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`(cmd&cmd&)`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.SubShell{
	// 			Body:ast.Statement{
	// 			ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
	// 			ast.BackgroundConstruction{Statement: ast.Command{Name: ast.Word("cmd")}},
	// 			},
	// 		},
	// 	},
	// }},
	// {`(cmd1 | cmd2 && cmd3; cmd1 | cmd2 && cmd3)`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.SubShell{
	// 			Body:ast.Statement{
	// 			ast.BinaryConstruction{
	// 				Left: ast.Pipeline{
	// 					{Command: ast.Command{Name: ast.Word("cmd1")}},
	// 					{Command: ast.Command{Name: ast.Word("cmd2")}},
	// 				},
	// 				Operator: "&&",
	// 				Right:    ast.Command{Name: ast.Word("cmd3")},
	// 			},
	// 			ast.BinaryConstruction{
	// 				Left: ast.Pipeline{
	// 					{Command: ast.Command{Name: ast.Word("cmd1")}},
	// 					{Command: ast.Command{Name: ast.Word("cmd2")}},
	// 				},
	// 				Operator: "&&",
	// 				Right:    ast.Command{Name: ast.Word("cmd3")},
	// 			},
	// 		},
	// 	},
	// }},

	// {`(cmd; cmd) | (cmd; cmd)&& (cmd; cmd)`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.BinaryConstruction{
	// 			Left: ast.Pipeline{
	// 				{Command: ast.SubShell{
	// 					Body:ast.Statement{
	// 					ast.Command{Name: ast.Word("cmd")},
	// 					ast.Command{Name: ast.Word("cmd")},
	// 				}},
	// 				{Command: ast.SubShell{
	// 					Body:ast.Statement{
	// 					ast.Command{Name: ast.Word("cmd")},
	// 					ast.Command{Name: ast.Word("cmd")},
	// 				}},
	// 			},
	// 			Operator: "&&",
	// 			Right: ast.SubShell{
	// 				Body:ast.Statement{
	// 				ast.Command{Name: ast.Word("cmd")},
	// 				ast.Command{Name: ast.Word("cmd")},
	// 			},
	// 		},
	// 	},
	// }},
	// {`(cmd) >output.txt <input.txt 2>error.txt >&3 \
	// 	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
	// 	Statements: []ast.Statement{
	// 		ast.SubShell{
	// 			Body:ast.Statement{
	// 			ast.Command{Name: ast.Word("cmd")},
	// 			ast.Command{Name: ast.Word("cmd")},
	// 		},
	// 	},
	// }},
}

var groupingErrorHandlingCases = []errorHandlingTestCase{
	{`{`, "syntax error: expeceted a command list after `{`."},
	{`{}`, "syntax error: expeceted a command list after `{`."},
	{`{cmd`, "syntax error: unexpected end of file, expeceted `}`."},
	{`{cmd}`, "syntax error: unexpected end of file, expeceted `}`."},
	{`{cmd |;}`, "syntax error: `;` has a special meaning here and cannot be used as a command name."},
	{`{cmd | |}`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},

	{`(`, "syntax error: expeceted a command list after `(`."},
	{`()`, "syntax error: expeceted a command list after `(`."},
	{`(cmd`, "syntax error: unexpected end of file, expeceted `)`."},
	{`(cmd |)`, "syntax error: `)` has a special meaning here and cannot be used as a command name."},
	{`(cmd | |)`, "syntax error: `|` has a special meaning here and cannot be used as a command name."},
}

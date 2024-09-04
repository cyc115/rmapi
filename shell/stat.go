package shell

import (
	"encoding/json"
	"errors"

	"github.com/abiosoft/ishell"
)

func statCmd(ctx *ShellCtxt) *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "stat",
		Help:      "fetch entry metadata",
		Completer: createEntryCompleter(ctx),
		Func: func(c *ishell.Context) {
			if len(c.Args) == 0 {
				c.Err(errors.New("missing source file"))
				return
			}

			srcName := c.Args[0]

			node, err := ctx.api.Filetree().NodeByPath(srcName, ctx.node)

			if err != nil {
				c.Err(errors.New("file doesn't exist"))
				return
			}
			// // Create a custom struct to include FileTags
			// type CustomDocument struct {
			// 	*model.Document
			// 	FileTags []string `json:"tags,omitempty"`
			// }

			// customDoc := CustomDocument{
			// 	Document: node.Document,
			// 	FileTags: node.Document.FileTags,
			// }

			jsn, err := json.MarshalIndent(node.Document, "", "  ")

			if err != nil {
				c.Err(errors.New("can't serialize to json"))
				return
			}

			c.Println(string(jsn))
		},
	}
}

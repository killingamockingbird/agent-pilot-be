package tool

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
	"os/exec"
	"runtime"
)

type ShellTool struct{}

func (t *ShellTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "shell",
		Desc: "Execute shell command",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"cmd": {
				Type:     schema.String,
				Required: true,
			},
		}),
	}, nil
}

func (t *ShellTool) InvokableRun(
	ctx context.Context,
	args string,
	opts ...tool.Option,
) (string, error) {

	var input struct {
		Cmd string `json:"cmd"`
	}

	if err := json.Unmarshal([]byte(args), &input); err != nil {
		return "", err
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", input.Cmd)
	} else {
		cmd = exec.Command("bash", "-c", input.Cmd)
	}

	out, err := cmd.CombinedOutput()

	return string(out), err
}

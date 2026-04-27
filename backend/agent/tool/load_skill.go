package tool

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"

	"github.com/agent-pilot/agent-pilot-be/agent/tool/skill"
)

type LoadSkillTool struct {
	Reg *skill.Registry
}

func (t *LoadSkillTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "load_skill",
		Desc: "Load a skill by name and inject it into the conversation. Use when task requires external capability like lark-im.",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"name": {
				Type:     schema.String,
				Desc:     "skill name",
				Required: true,
			},
		}),
	}, nil
}

func (t *LoadSkillTool) InvokableRun(
	ctx context.Context,
	args string,
	opts ...tool.Option,
) (string, error) {

	var input struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal([]byte(args), &input); err != nil {
		return "", err
	}

	get := t.Reg.Get(input.Name)
	if get == nil {
		return "get not found: " + input.Name, nil
	}

	// 加载 references
	refs := get.LoadReferences()

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[SKILL LOADED: %s]\n\n", get.Name))
	sb.WriteString(get.Content)

	if len(refs) > 0 {
		sb.WriteString("\n\n=== REFERENCES ===\n")
		for name, body := range refs {
			sb.WriteString(fmt.Sprintf("\n--- %s ---\n%s\n", name, body))
		}
	}

	return sb.String(), nil
}

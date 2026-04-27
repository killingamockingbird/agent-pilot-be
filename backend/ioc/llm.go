package ioc

import (
	"context"
	"time"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"

	"github.com/agent-pilot/agent-pilot-be/pkg/llm"
)

// NewOpenAIModelClient 创建一个使用eino openai.model的client
func NewOpenAIModelClient(ctx context.Context, model, baseUrl, apikey string) *llm.ChatClient {
	om, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  apikey,
		Model:   model,
		BaseURL: baseUrl,
		Timeout: 360 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	return llm.NewChatClient(om)
}

func NewReactAgent(ctx context.Context, c *llm.ChatClient, tools []tool.BaseTool) *llm.ReactAgent {
	core, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: c,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: tools,
		},
	})

	if err != nil {
		panic(err)
	}
	return &llm.ReactAgent{
		Core: core,
	}
}

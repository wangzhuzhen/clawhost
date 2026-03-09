package model

import (
	"encoding/json"
	"time"
)

// OpenClawConfig represents the full openclaw.json configuration format
// This structure matches the openclaw configuration file format
type OpenClawConfig struct {
	Meta     *MetaConfig     `json:"meta,omitempty"`
	Models   *ModelsConfig   `json:"models,omitempty"`
	Agents   *AgentsConfig   `json:"agents,omitempty"`
	Channels ChannelsConfig  `json:"channels,omitempty"`
	Gateway  *GatewayConfig  `json:"gateway,omitempty"`
	Auth     *AuthConfig     `json:"auth,omitempty"`
	Plugins  *PluginsConfig  `json:"plugins,omitempty"`
	Messages *MessagesConfig `json:"messages,omitempty"`
	Commands *CommandsConfig `json:"commands,omitempty"`
	Hooks    *HooksConfig    `json:"hooks,omitempty"`
	Wizard   *WizardConfig   `json:"wizard,omitempty"`
	Identity *IdentityConfig `json:"identity,omitempty"`
	Logging  *LoggingConfig  `json:"logging,omitempty"`
	Routing  *RoutingConfig  `json:"routing,omitempty"`
	Tools    *ToolsConfig    `json:"tools,omitempty"`
	Session  *SessionConfig  `json:"session,omitempty"`
	Cron     *CronConfig     `json:"cron,omitempty"`
	Skills   *SkillsConfig   `json:"skills,omitempty"`
}

// MetaConfig represents metadata about the configuration
type MetaConfig struct {
	LastTouchedVersion string `json:"lastTouchedVersion,omitempty"`
	LastTouchedAt      string `json:"lastTouchedAt,omitempty"`
}

// WizardConfig represents wizard about the configuration
type WizardConfig struct {
	LastRunAt      time.Time `json:"lastRunAt,omitempty"`
	LastRunVersion string    `json:"LastRunVersion,omitempty"`
	LastRunCommand string    `json:"LastRunCommand,omitempty"`
	LastRunMode    string    `json:"LastRunMode,omitempty"`
}

// ModelsConfig represents the models configuration
//
//	自定义模型提供商 配置示例：
//	models: {
//	  mode: "merge",
//	  providers: {
//	    "custom-proxy": {
//	      baseUrl: "http://localhost:4000/v1",
//	      apiKey: "LITELLM_KEY",
//	      api: "openai-responses",
//	      authHeader: true,
//	      headers: { "X-Proxy-Region": "us-west" },
//	      models: [
//	        {
//	          id: "llama-3.1-8b",
//	          name: "Llama 3.1 8B",
//	          api: "openai-responses",
//	          reasoning: false,
//	          input: ["text"],
//	          cost: { input: 0, output: 0, cacheRead: 0, cacheWrite: 0 },
//	          contextWindow: 128000,
//	          maxTokens: 32000,
//	        },
//	      ],
//	    },
//	  },
//	},
type ModelsConfig struct {
	Mode      string                     `json:"mode,omitempty"` // "merge" or "replace"
	Providers map[string]*ProviderConfig `json:"providers,omitempty"`
}

// ProviderConfig represents a single provider configuration
type ProviderConfig struct {
	BaseURL    string                `json:"baseUrl,omitempty"`
	APIKey     string                `json:"apiKey,omitempty"`
	Auth       string                `json:"auth,omitempty"`       // api-key, bearer
	AuthHeader bool                  `json:"authHeader,omitempty"` // whether to send API key in Authorization header
	API        string                `json:"api,omitempty"`        // anthropic-messages, openai-completions
	Models     []ProviderModelConfig `json:"models,omitempty"`
	Headers    map[string]string     `json:"headers,omitempty"`
}

// ProviderModelConfig represents a model configuration within a provider
type ProviderModelConfig struct {
	ID            string   `json:"id"`
	Name          string   `json:"name,omitempty"`
	API           string   `json:"api,omitempty"` // eg.  api: "openai-responses",
	Reasoning     bool     `json:"reasoning,omitempty"`
	Input         []string `json:"input,omitempty"`
	Cost          *Cost    `json:"cost,omitempty"`
	ContextWindow int      `json:"contextWindow,omitempty"`
	MaxTokens     int      `json:"maxTokens,omitempty"`
}

type Cost struct {
	Input      int64 `json:"input,omitempty"`
	Output     int64 `json:"output,omitempty"`
	CacheRead  int64 `json:"cacheRead,omitempty"`
	CacheWrite int64 `json:"cacheWrite,omitempty"`
}

// AgentsConfig represents the agents configuration
//
//	智能体运行时 配置实例：
//	agents: {
//	  defaults: {
//	    workspace: "~/.openclaw/workspace",
//	    userTimezone: "America/Chicago",
//	    model: {
//	      primary: "anthropic/claude-sonnet-4-5",
//	      fallbacks: ["anthropic/claude-opus-4-5", "openai/gpt-5.2"],
//	    },
//	    imageModel: {
//	      primary: "openrouter/anthropic/claude-sonnet-4-5",
//	    },
//	    models: {
//	      "anthropic/claude-opus-4-5": { alias: "opus" },
//	      "anthropic/claude-sonnet-4-5": { alias: "sonnet" },
//	      "openai/gpt-5.2": { alias: "gpt" },
//	    },
//	    thinkingDefault: "low",
//	    verboseDefault: "off",
//	    elevatedDefault: "on",
//	    blockStreamingDefault: "off",
//	    blockStreamingBreak: "text_end",
//	    blockStreamingChunk: {
//	      minChars: 800,
//	      maxChars: 1200,
//	      breakPreference: "paragraph",
//	    },
//	    blockStreamingCoalesce: {
//	      idleMs: 1000,
//	    },
//	    humanDelay: {
//	      mode: "natural",
//	    },
//	    timeoutSeconds: 600,
//	    mediaMaxMb: 5,
//	    typingIntervalSeconds: 5,
//	    maxConcurrent: 3,
//	    heartbeat: {
//	      every: "30m",
//	      model: "anthropic/claude-sonnet-4-5",
//	      target: "last",
//	      to: "+15555550123",
//	      prompt: "HEARTBEAT",
//	      ackMaxChars: 300,
//	    },
//	    memorySearch: {
//	      provider: "gemini",
//	      model: "gemini-embedding-001",
//	      remote: {
//	        apiKey: "${GEMINI_API_KEY}",
//	      },
//	      extraPaths: ["../team-docs", "/srv/shared-notes"],
//	    },
//	    sandbox: {
//	      mode: "non-main",
//	      perSession: true,
//	      workspaceRoot: "~/.openclaw/sandboxes",
//	      docker: {
//	        image: "openclaw-sandbox:bookworm-slim",
//	        workdir: "/workspace",
//	        readOnlyRoot: true,
//	        tmpfs: ["/tmp", "/var/tmp", "/run"],
//	        network: "none",
//	        user: "1000:1000",
//	      },
//	      browser: {
//	        enabled: false,
//	      },
//	    },
//	  },
//	},
type AgentsConfig struct {
	Defaults *AgentDefaultsConfig `json:"defaults,omitempty"`
}

// AgentDefaultsConfig represents agent default settings
type AgentDefaultsConfig struct {
	Model                  *AgentModelConfig           `json:"model,omitempty"`
	Models                 map[string]*AgentModelAlias `json:"models,omitempty"`
	Workspace              string                      `json:"workspace,omitempty"`
	Compaction             *CompactionConfig           `json:"compaction,omitempty"`
	MaxConcurrent          int                         `json:"maxConcurrent,omitempty"`
	Subagents              *SubagentsConfig            `json:"subagents,omitempty"`
	UserTimezone           string                      `json:"userTimezone,omitempty"`
	ImageModel             *AgentModelConfig           `json:"imageModel,omitempty"`
	ThinkingDefault        string                      `json:"thinkingDefault,omitempty"`
	VerboseDefault         string                      `json:"verboseDefault,omitempty"`
	ElevatedDefault        string                      `json:"elevatedDefault,omitempty"`
	BlockStreamingDefault  string                      `json:"blockStreamingDefault,omitempty"`
	BlockStreamingBreak    string                      `json:"blockStreamingBreak,omitempty"`
	BlockStreamingChunk    BlockStreamingChunk         `json:"blockStreamingChunk,omitempty"`
	BlockStreamingCoalesce BlockStreamingCoalesce      `json:"blockStreamingCoalesce,omitempty"`
	HumanDelay             HumanDelay                  `json:"humanDelay,omitempty"`
	TimeoutSeconds         *int                        `json:"timeoutSeconds,omitempty"`
	MediaMaxMb             *int                        `json:"mediaMaxMb,omitempty"`
	TypingIntervalSeconds  *int                        `json:"typingIntervalSeconds,omitempty"`
	Heartbeat              *Heartbeat                  `json:"heartbeat,omitempty"`
	MemorySearch           *MemorySearch               `json:"memorySearch,omitempty"`
	Sandbox                *Sandbox                    `json:"sandbox,omitempty"`
}

type BlockStreamingChunk struct {
	MinChars        int    `json:"minChars,omitempty"`
	MaxChars        int    `json:"maxChars,omitempty"`
	BreakPreference string `json:"breakPreference,omitempty"`
}

type BlockStreamingCoalesce struct {
	IdleMs *int `json:"idleMs,omitempty"`
}

type HumanDelay struct {
	Mode string `json:"mode,omitempty"`
}

type Heartbeat struct {
	Every       string `json:"every,omitempty"`
	Model       string `json:"model,omitempty"`
	Target      string `json:"target,omitempty"`
	To          string `json:"to,omitempty"`
	Prompt      string `json:"prompt,omitempty"`
	AckMaxChars int    `json:"ackMaxChars,omitempty"`
}

type MemorySearch struct {
	Provider   string             `json:"provider,omitempty"`
	Model      string             `json:"model,omitempty"`
	Remote     MemorySearchRemote `json:"remote,omitempty"`
	ExtraPaths []string           `json:"extraPaths,omitempty"`
}

type MemorySearchRemote struct {
	ApiKey string `json:"apiKey,omitempty"`
}

type Sandbox struct {
	Mode          string        `json:"mode,omitempty"`
	PerSession    bool          `json:"perSession,omitempty"`
	WorkspaceRoot string        `json:"workspaceRoot,omitempty"`
	Docker        DockerConfig  `json:"docker,omitempty"`
	Browser       BrowserConfig `json:"browser,omitempty"`
}
type DockerConfig struct {
	Image        string   `json:"image,omitempty"`
	Workdir      string   `json:"workdir,omitempty"`
	ReadOnlyRoot bool     `json:"readOnlyRoot,omitempty"`
	Tmpfs        []string `json:"tmpfs"`
	Network      string   `json:"network,omitempty"`
	User         string   `json:"user,omitempty"` // eg, "1000:1000",
}

type BrowserConfig struct {
	Enabled bool `json:"enabled,omitempty"`
}

// AgentModelConfig represents the primary model configuration
type AgentModelConfig struct {
	Primary   string   `json:"primary,omitempty"`   // e.g., "anthropic/claude-sonnet-4-20250514"
	Fallbacks []string `json:"fallbacks,omitempty"` // e.g.,  ["anthropic/claude-opus-4-5", "openai/gpt-5.2"]
}

// AgentModelAlias represents a model alias configuration
type AgentModelAlias struct {
	Alias string `json:"alias,omitempty"`
}

// CompactionConfig represents compaction settings
type CompactionConfig struct {
	Mode string `json:"mode,omitempty"` // "safeguard", etc.
}

// SubagentsConfig represents subagent settings
type SubagentsConfig struct {
	MaxConcurrent int `json:"maxConcurrent,omitempty"`
}

// ChannelsConfig is a map of channel name to channel configuration
// Using interface{} to support various channel config formats (accounts, botToken, etc.)
// 渠道 配置示例：
//
//	channels: {
//	  whatsapp: {
//	    dmPolicy: "pairing",
//	    allowFrom: ["+15555550123"],
//	    groupPolicy: "allowlist",
//	    groupAllowFrom: ["+15555550123"],
//	    groups: { "*": { requireMention: true } },
//	  },
//
//	  telegram: {
//	    enabled: true,
//	    botToken: "YOUR_TELEGRAM_BOT_TOKEN",
//	    allowFrom: ["123456789"],
//	    groupPolicy: "allowlist",
//	    groupAllowFrom: ["123456789"],
//	    groups: { "*": { requireMention: true } },
//	  },
//
//	  discord: {
//	    enabled: true,
//	    token: "YOUR_DISCORD_BOT_TOKEN",
//	    dm: { enabled: true, allowFrom: ["steipete"] },
//	    guilds: {
//	      "123456789012345678": {
//	        slug: "friends-of-openclaw",
//	        requireMention: false,
//	        channels: {
//	          general: { allow: true },
//	          help: { allow: true, requireMention: true },
//	        },
//	      },
//	    },
//	  },
//
//	  slack: {
//	    enabled: true,
//	    botToken: "xoxb-REPLACE_ME",
//	    appToken: "xapp-REPLACE_ME",
//	    channels: {
//	      "#general": { allow: true, requireMention: true },
//	    },
//	    dm: { enabled: true, allowFrom: ["U123"] },
//	    slashCommand: {
//	      enabled: true,
//	      name: "openclaw",
//	      sessionPrefix: "slack:slash",
//	      ephemeral: true,
//	    },
//	  },
//	},
type ChannelsConfig map[string]interface{}

// ChannelConfig represents a single channel configuration
type ChannelConfig struct {
	Enabled        bool                   `json:"enabled,omitempty"`
	BotToken       string                 `json:"botToken,omitempty"`      // Telegram, Discord
	AppToken       string                 `json:"appToken,omitempty"`      // Slack
	AppID          string                 `json:"appId,omitempty"`         // Feishu, Teams
	AppSecret      string                 `json:"appSecret,omitempty"`     // Feishu
	AppPassword    string                 `json:"appPassword,omitempty"`   // Teams
	ChannelSecret  string                 `json:"channelSecret,omitempty"` // LINE
	DMPolicy       string                 `json:"dmPolicy,omitempty"`      // pairing, allowlist, open, disabled
	GroupPolicy    string                 `json:"groupPolicy,omitempty"`   // open, allowlist, disabled
	TextChunkLimit int                    `json:"textChunkLimit,omitempty"`
	MediaMaxMb     int                    `json:"mediaMaxMb,omitempty"`
	Extra          map[string]interface{} `json:"-"` // Additional fields
}

// MarshalJSON implements custom JSON marshaling for ChannelConfig
func (c *ChannelConfig) MarshalJSON() ([]byte, error) {
	type Alias ChannelConfig
	data, err := json.Marshal((*Alias)(c))
	if err != nil {
		return nil, err
	}

	if len(c.Extra) == 0 {
		return data, nil
	}

	// Merge extra fields
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	for k, v := range c.Extra {
		if _, exists := m[k]; !exists {
			m[k] = v
		}
	}
	return json.Marshal(m)
}

// UnmarshalJSON implements custom JSON unmarshaling for ChannelConfig
func (c *ChannelConfig) UnmarshalJSON(data []byte) error {
	type Alias ChannelConfig
	if err := json.Unmarshal(data, (*Alias)(c)); err != nil {
		return err
	}

	// Capture unknown fields in Extra
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	knownFields := map[string]bool{
		"enabled": true, "botToken": true, "appToken": true, "appId": true,
		"appSecret": true, "appPassword": true, "channelSecret": true,
		"dmPolicy": true, "groupPolicy": true, "textChunkLimit": true, "mediaMaxMb": true,
	}

	c.Extra = make(map[string]interface{})
	for k, v := range m {
		if !knownFields[k] {
			c.Extra[k] = v
		}
	}
	return nil
}

// GatewayConfig represents gateway configuration
type GatewayConfig struct {
	Port           int                  `json:"port,omitempty"`
	Mode           string               `json:"mode,omitempty"` // local, remote
	Bind           string               `json:"bind,omitempty"` // loopback, lan
	Auth           *GatewayAuthConfig   `json:"auth,omitempty"`
	Tailscale      *TailscaleConfig     `json:"tailscale,omitempty"`
	TrustedProxies []string             `json:"trustedProxies,omitempty"`
	ControlUI      *ControlUIConfig     `json:"controlUi,omitempty"`
	Nodes          *Nodes               `json:"nodes,omitempty"`
	Remote         *RemoteGatewayConfig `json:"remote,omitempty"`
	Reload         *GatewayReload       `json:"reload,omitempty"`
}

type RemoteGatewayConfig struct {
	URL   string `json:"url,omitempty"`
	Token string `json:"token,omitempty"`
}

type GatewayReload struct {
	Mode       string `json:"mode,omitempty"`
	DebounceMs int    `json:"debounceMs,omitempty"`
}

type Nodes struct {
	DenyCommands []string `json:"denyCommands,omitempty"`
}

// GatewayAuthConfig represents gateway authentication configuration
type GatewayAuthConfig struct {
	Mode           string `json:"mode,omitempty"` // password, token
	Password       string `json:"password,omitempty"`
	Token          string `json:"token,omitempty"`
	AllowTailscale *bool  `json:"allowTailscale,omitempty"`
}

// TailscaleConfig represents Tailscale configuration
type TailscaleConfig struct {
	Mode        string `json:"mode,omitempty"` // off, on
	ResetOnExit bool   `json:"resetOnExit,omitempty"`
	Path        string `json:"path,omitempty"`
}

// ControlUIConfig represents control UI configuration
type ControlUIConfig struct {
	Enabled                                  *bool    `json:"enabled,omitempty"`
	AllowedOrigins                           []string `json:"allowedOrigins,omitempty"`
	DangerouslyDisableDeviceAuth             bool     `json:"dangerouslyDisableDeviceAuth,omitempty"`
	DangerouslyAllowHostHeaderOriginFallback *bool    `json:"dangerouslyAllowHostHeaderOriginFallback,omitempty"`
	AllowInsecureAuth                        *bool    `json:"allowInsecureAuth,omitempty"`
}

// AuthConfig represents authentication profiles configuration
// 认证配置文件元数据（密钥存储在 auth-profiles.json 中）
type AuthConfig struct {
	// Profiles 配置实例:
	// profiles: {
	//   "anthropic:me@example.com": { provider: "anthropic", mode: "oauth", email: "me@example.com" },
	//   "anthropic:work": { provider: "anthropic", mode: "api_key" },
	//   "openai:default": { provider: "openai", mode: "api_key" },
	//   "openai-codex:default": { provider: "openai-codex", mode: "oauth" },
	// }
	Profiles map[string]*AuthProfile `json:"profiles,omitempty"`
	// Orders 配置示例:
	// order: {
	//   anthropic: ["anthropic:me@example.com", "anthropic:work"],
	//   openai: ["openai:default"],
	//   "openai-codex": ["openai-codex:default"],
	// },
	Orders map[string][]string `json:"orders,omitempty"`
}

// AuthProfile represents an authentication profile
type AuthProfile struct {
	Provider string `json:"provider,omitempty"`
	Mode     string `json:"mode,omitempty"`
}

// PluginsConfig represents plugins configuration
type PluginsConfig struct {
	Entries map[string]interface{} `json:"entries,omitempty"`
}

// MessagesConfig represents messages configuration
// 消息格式配置实例
//
//	messages: {
//	  messagePrefix: "[openclaw]",
//	  responsePrefix: ">",
//	  ackReaction: "👀",
//	  ackReactionScope: "group-mentions",
//	},
type MessagesConfig struct {
	MessagePrefix    string `json:"messagePrefix,omitempty"`
	ResponsePrefix   string `json:"responsePrefix,omitempty"`
	AckReaction      string `json:"ackReaction,omitempty"`
	AckReactionScope string `json:"ackReactionScope,omitempty"`
}

// CommandsConfig represents commands configuration
type CommandsConfig struct {
	Native       string `json:"native,omitempty"`
	NativeSkills string `json:"nativeSkills,omitempty"`
}

// HooksConfig represents hooks configuration
type HooksConfig struct {
	Internal      *InternalHooksConfig `json:"internal,omitempty"`
	Enabled       bool                 `json:"enabled,omitempty"`
	Path          string               `json:"path,omitempty"`
	Token         string               `json:"token,omitempty"`
	Presets       []string             `json:"presets,omitempty"`
	TransformsDir string               `json:"transformsDir,omitempty"`
	Mappings      []Hook               `json:"mappings,omitempty"`
	Gmail         Gmail                `json:"gmail,omitempty"`
}

type Hook struct {
	ID              string        `json:"id,omitempty"`
	Match           HookMatch     `json:"match,omitempty"`
	Action          string        `json:"action,omitempty"`
	WakeMode        string        `json:"wakeMode,omitempty"`
	Name            string        `json:"name,omitempty"`
	SessionKey      string        `json:"sessionKey,omitempty"`
	MessageTemplate string        `json:"messageTemplate,omitempty"`
	TextTemplate    string        `json:"textTemplate,omitempty"`
	Deliver         bool          `json:"deliver,omitempty"`
	Channel         string        `json:"channel,omitempty"`
	To              string        `json:"to,omitempty"`
	Thinking        string        `json:"thinking"`
	TimeoutSeconds  int           `json:"timeoutSeconds"`
	Transform       HookTransform `json:"transform,omitempty"`
}

type HookMatch struct {
	Path string `json:"path,omitempty"`
}

type HookTransform struct {
	Module string `json:"module,omitempty"`
	Export string `json:"export,omitempty"`
}

type Gmail struct {
	Account           string          `json:"acount,omitempty"`
	Label             string          `json:"label,omitempty"`
	Topic             string          `json:"topic,omitempty"`
	Subscription      string          `json:"subscription,omitempty"`
	PushToken         string          `json:"pushToken,omitempty"`
	HookUrl           string          `json:"hookUrl,omitempty"`
	IncludeBody       bool            `json:"includeBody,omitempty"`
	MaxBytes          int64           `json:"maxBytes,omitempty"`
	RenewEveryMinutes int             `json:"renewEveryMinutes,omitempty"`
	Serve             GmailServe      `json:"serve,omitempty"`
	Tailscale         TailscaleConfig `json:"tailscale,omitempty"`
}

type GmailServe struct {
	Bind string `json:"bind,omitempty"`
	Port int    `json:"port,omitempty"`
	Path string `json:"path,omitempty"`
}

// SkillsConfig represents skills configuration
type SkillsConfig struct {
	AllowBundled []string              `json:"allowBundled,omitempty"`
	Load         SkillLoad             `json:"load,omitempty"`
	Install      SkillInstall          `json:"install,omitempty"`
	Entries      map[string]SkillEntry `json:"entries,omitempty"`
}

type SkillLoad struct {
	ExtraDirs []string `json:"extraDirs,omitempty"`
}

type SkillInstall struct {
	PreferBrew  bool   `json:"preferBrew,omitempty"`
	NodeManager string `json:"nodeManager,omitempty"`
}

type SkillEntry struct {
	Enabled bool              `json:"enabled,omitempty"`
	ApiKey  string            `json:"apiKey,omitempty"`
	Env     map[string]string `json:"env,omitempty"`
}

// InternalHooksConfig represents internal hooks configuration
type InternalHooksConfig struct {
	Enabled bool                        `json:"enabled,omitempty"`
	Entries map[string]*HookEntryConfig `json:"entries,omitempty"`
}

// HookEntryConfig represents a hook entry configuration
type HookEntryConfig struct {
	Enabled bool `json:"enabled,omitempty"`
}

// IdentityConfig represents identity about the configuration
//
//	身份配置示例
//	identity: {
//	  name: "Samantha",
//	  theme: "helpful sloth",
//	  emoji: "🦥",
//	},
type IdentityConfig struct {
	Name  string `json:"name,omitempty"`
	Theme string `json:"theme,omitempty"`
	Emoji string `json:"emoji,omitempty"`
}

// LoggingConfig represents logging about the configuration
// 日志配置示例:
//
//	logging: {
//	   level: "info",
//	   file: "/tmp/openclaw/openclaw.log",
//	   consoleLevel: "info",
//	   consoleStyle: "pretty",
//	   redactSensitive: "tools",
//	 },
type LoggingConfig struct {
	Level           string `json:"level,omitempty"`
	File            string `json:"file,omitempty"`
	ConsoleLevel    string `json:"consoleLevel,omitempty"`
	ConsoleStyle    string `json:"consoleStyle,omitempty"`
	RedactSensitive string `json:"redactSensitive,omitempty"`
}

// RoutingConfig represents routing about the configuration
// 路由 + 队列 配置示例：
//
//	routing: {
//	  groupChat: {
//	    mentionPatterns: ["@openclaw", "openclaw"],
//	    historyLimit: 50,
//	  },
//	  queue: {
//	    mode: "collect",
//	    debounceMs: 1000,
//	    cap: 20,
//	    drop: "summarize",
//	    byChannel: {
//	      whatsapp: "collect",
//	      telegram: "collect",
//	      discord: "collect",
//	      slack: "collect",
//	      signal: "collect",
//	      imessage: "collect",
//	      webchat: "collect",
//	    },
//	  },
//	},
type RoutingConfig struct {
	GroupChat GroupChat `json:"groupChat,omitempty"`
	Queue     *Queue    `json:"queue,omitempty"`
}

type GroupChat struct {
	MentionPatterns []string `json:"mentionPatterns,omitempty"`
	HistoryLimit    int      `json:"historyLimit,omitempty"`
}

type Queue struct {
	Model      string    `json:"model,omitempty"`
	DebounceMs *int      `json:"debounceMs,omitempty"`
	Cap        *int      `json:"cap,omitempty"`
	Drop       string    `json:"drop,omitempty"`
	ByChannel  ByChannel `json:"byChannel,omitempty"`
}

type ByChannel struct {
	WhatsApp string `json:"whatsapp,omitempty"`
	Telegram string `json:"telegram,omitempty"`
	Discord  string `json:"discord,omitempty"`
	Slack    string `json:"slack,omitempty"`
	Signal   string `json:"signal,omitempty"`
	Imessage string `json:"imessage,omitempty"`
	WebChat  string `json:"webchat,omitempty"`
}

// ToolsConfig represents tools about the configuration
//
//	工具 配置示例
//	tools: {
//	  media: {
//	    audio: {
//	      enabled: true,
//	      maxBytes: 20971520,
//	      models: [
//	        { provider: "openai", model: "gpt-4o-mini-transcribe" },
//	        // 可选的 CLI 回退（Whisper 二进制）：
//	        // { type: "cli", command: "whisper", args: ["--model", "base", "{{MediaPath}}"] }
//	      ],
//	      timeoutSeconds: 120,
//	    },
//	    video: {
//	      enabled: true,
//	      maxBytes: 52428800,
//	      models: [{ provider: "google", model: "gemini-3-flash-preview" }],
//	    },
//	  },
//	},
type ToolsConfig struct {
	Media Media `json:"media,omitempty"`
}

type Media struct {
	Audio MediaInfo `json:"audio,omitempty"`
}

type MediaInfo struct {
	Enabled        bool         `json:"enabled,omitempty"`
	MaxBytes       *int         `json:"maxBytes,omitempty"`
	Models         []MediaModel `json:"models,omitempty"`
	TimeoutSeconds *int         `json:"timeoutSeconds,omitempty"`
}

type MediaModel struct {
	Provider string `json:"provider,omitempty"`
	Model    string `json:"model,omitempty"`
}

// SessionConfig represents session about the configuration
// 会话行为配置示例:
//
//	session: {
//	  scope: "per-sender",
//	  reset: {
//	    mode: "daily",
//	    atHour: 4,
//	    idleMinutes: 60,
//	  },
//	  resetByChannel: {
//	    discord: { mode: "idle", idleMinutes: 10080 },
//	  },
//	  resetTriggers: ["/new", "/reset"],
//	  store: "~/.openclaw/agents/default/sessions/sessions.json",
//	  typingIntervalSeconds: 5,
//	  sendPolicy: {
//	    default: "allow",
//	    rules: [{ action: "deny", match: { channel: "discord", chatType: "group" } }],
//	  },
//	},
type SessionConfig struct {
	Scope                 string                  `json:"scope,omitempty"`
	Reset                 SessionReset            `json:"reset,omitempty"`
	ResetByChannel        map[string]SessionReset `json:"resetByChannel,omitempty"`
	ResetTriggers         []string                `json:"resetTriggers,omitempty"`
	Store                 string                  `json:"store,omitempty"`
	TypingIntervalSeconds *int                    `json:"typingIntervalSeconds,omitempty"`
	SendPolicy            SendPolicy              `json:"sendPolicy,omitempty"`
}

type SessionReset struct {
	Mode        string `json:"mode,omitempty"`
	AtHour      *int   `json:"atHour,omitempty"`
	IdleMinutes *int   `json:"idleMinutes,omitempty"`
}

type SendPolicy struct {
	Default string           `json:"default,omitempty"`
	Rules   []SendPolicyRule `json:"rules,omitempty"`
}

type SendPolicyRule struct {
	Action string              `json:"action,omitempty"`
	Match  SendPolicyRuleMatch `json:"match,omitempty"`
}

type SendPolicyRuleMatch struct {
	Channel  string `json:"channel,omitempty"`
	ChatType string `json:"chatType,omitempty"`
}

// CronConfig represents cron about the configuration
//
//	Cron 作业 配置示例:
//	cron: {
//	  enabled: true,
//	  store: "~/.openclaw/cron/cron.json",
//	  maxConcurrentRuns: 2,
//	},
type CronConfig struct {
	Enabled           bool   `json:"enabled,omitempty"`
	Store             string `json:"store,omitempty"`
	MaxConcurrentRuns int    `json:"maxConcurrentRuns,omitempty"`
}

// GetBotOpenClawConfig returns the OpenClaw configuration from bot.Config
func (b *Bot) GetOpenClawConfig() (*OpenClawConfig, error) {
	if b.Config == nil {
		return &OpenClawConfig{}, nil
	}
	var config OpenClawConfig
	if err := json.Unmarshal(b.Config, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// SetOpenClawConfig sets the OpenClaw configuration to bot.Config
func (b *Bot) SetOpenClawConfig(config *OpenClawConfig) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	b.Config = data
	return nil
}

// MergeOpenClawConfig merges partial config into existing config
func (b *Bot) MergeOpenClawConfig(partial *OpenClawConfig) error {
	existing, err := b.GetOpenClawConfig()
	if err != nil {
		existing = &OpenClawConfig{}
	}

	// Merge models
	if partial.Models != nil {
		if existing.Models == nil {
			existing.Models = &ModelsConfig{}
		}
		if partial.Models.Mode != "" {
			existing.Models.Mode = partial.Models.Mode
		}
		if partial.Models.Providers != nil {
			if existing.Models.Providers == nil {
				existing.Models.Providers = make(map[string]*ProviderConfig)
			}
			for name, provider := range partial.Models.Providers {
				existing.Models.Providers[name] = provider
			}
		}
	}

	// Merge agents
	if partial.Agents != nil {
		if existing.Agents == nil {
			existing.Agents = &AgentsConfig{}
		}
		if partial.Agents.Defaults != nil {
			if existing.Agents.Defaults == nil {
				existing.Agents.Defaults = &AgentDefaultsConfig{}
			}
			if partial.Agents.Defaults.Model != nil {
				existing.Agents.Defaults.Model = partial.Agents.Defaults.Model
			}
			if partial.Agents.Defaults.Models != nil {
				existing.Agents.Defaults.Models = partial.Agents.Defaults.Models
			}
			if partial.Agents.Defaults.MaxConcurrent > 0 {
				existing.Agents.Defaults.MaxConcurrent = partial.Agents.Defaults.MaxConcurrent
			}
		}
	}

	// Merge channels
	if partial.Channels != nil {
		if existing.Channels == nil {
			existing.Channels = make(ChannelsConfig)
		}
		for name, channel := range partial.Channels {
			existing.Channels[name] = channel
		}
	}

	// Merge gateway
	if partial.Gateway != nil {
		existing.Gateway = partial.Gateway
	}

	return b.SetOpenClawConfig(existing)
}

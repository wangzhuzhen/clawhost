package k8s

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"

	"github.com/fastclaw-ai/fastclaw/model"
)

// buildOpenClawConfig builds the openclaw.json configuration content
// setDefaultModel: if true, also sets agents.defaults.model.primary (for first-time setup)
func buildOpenClawConfigJson(config *BotConfig, setDefaultModel bool) string {

	// Build gateway section with password or token auth
	gatewayPort := getGatewayPort()

	enableGatewayControlUI := true
	allowInsecureAuth := true
	dangerouslyAllowHostHeaderOriginFallback := true
	openClawConfig := &model.OpenClawConfig{
		//Models:   nil,
		//Agents:   nil,
		//Channels: nil,
		Gateway: &model.GatewayConfig{
			Port: gatewayPort,
			Mode: "local",
			Bind: "lan",
			Auth: &model.GatewayAuthConfig{
				Mode:  "token",
				Token: config.AccessToken,
			},
			Tailscale: &model.TailscaleConfig{
				Mode:        "off",
				ResetOnExit: false,
			},
			TrustedProxies: getTrustedProxiesForJson(),
			ControlUI: &model.ControlUIConfig{
				Enabled:                                  &enableGatewayControlUI,
				DangerouslyDisableDeviceAuth:             true,
				DangerouslyAllowHostHeaderOriginFallback: &dangerouslyAllowHostHeaderOriginFallback,
				AllowInsecureAuth:                        &allowInsecureAuth,
				AllowedOrigins:                           viper.GetStringSlice("openclaw.allowed_origins"),
			},
			//Nodes: nil,
		},
		//Auth:     nil,
		//Plugins:  nil,
		//Messages: nil,
		//Commands: nil,
		//Hooks:    nil,
		//Skills:   nil,
	}

	// Build channels section if present
	if len(config.Channels) > 0 {
		openClawConfig.Channels = config.Channels
	}

	// Build Models section
	modelProviders := buildProvidersForJSON(config)
	openClawConfig.Models = &model.ModelsConfig{
		Mode:      "merge",
		Providers: modelProviders,
	}

	// Determine default model
	defaultModel := getDefaultModelFromConfigForJson(config)

	// Check if we have AgentDefaults set explicitly
	if config.AgentDefaults != nil && config.AgentDefaults.PrimaryModel != "" {
		defaultModel = config.AgentDefaults.PrimaryModel
	}

	if setDefaultModel && defaultModel != "" {
		openClawConfig.Agents = &model.AgentsConfig{
			Defaults: &model.AgentDefaultsConfig{
				Model: &model.AgentModelConfig{
					Primary: defaultModel,
				},
			},
		}

		if config.AgentDefaults != nil && config.AgentDefaults.FallbackModel != "" {
			agentModels := make(map[string]*model.AgentModelAlias)
			agentModels["fallback"] = &model.AgentModelAlias{
				Alias: config.AgentDefaults.FallbackModel,
			}
			openClawConfig.Agents.Defaults.Models = agentModels
		}
	}

	prettyJSON, err := json.MarshalIndent(*openClawConfig, "", "  ")
	if err != nil {
		fmt.Printf("MarshalIndent for openClawConfig failed: %v\n", err)
		return ""
	}

	fmt.Printf("New Bot Openclaw Config prettyJSON: \n%v\n", string(prettyJSON))
	return string(prettyJSON)
}

// getTrustedProxiesForJson returns the trusted proxy list from config
// Defaults to private network ranges if not configured
func getTrustedProxiesForJson() []string {
	proxies := viper.GetStringSlice("openclaw.trusted_proxies")
	if len(proxies) == 0 {
		// Default: trust all private network ranges (RFC 1918)
		// This allows proxies from 10.x.x.x, 172.16-31.x.x, 192.168.x.x
		proxies = []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "127.0.0.0/8"}
	}

	return proxies
}

// getDefaultModelFromConfigForJson returns the default model ID from config
func getDefaultModelFromConfigForJson(config *BotConfig) string {
	// Check AgentDefaults first
	if config.AgentDefaults != nil && config.AgentDefaults.PrimaryModel != "" {
		return config.AgentDefaults.PrimaryModel
	}

	// If we have providers, use the first provider's first model
	if len(config.Providers) > 0 {
		p := config.Providers[0]
		if len(p.Models) > 0 {
			return fmt.Sprintf("%s/%s", p.Name, p.Models[0].ID)
		}
	}

	// Fallback to legacy fields
	if config.Model != "" {
		providerName := getProviderName(config.Provider, config.BaseURL)
		return fmt.Sprintf("%s/%s", providerName, config.Model)
	}

	// Default
	return "tokenpony/deepseek-v3-0324"
}

// buildProvidersForJSON builds the providers JSON object from BotConfig
func buildProvidersForJSON(config *BotConfig) map[string]*model.ProviderConfig {
	providers := make(map[string]*model.ProviderConfig)

	// If we have multiple providers configured, use them
	if len(config.Providers) > 0 {
		for _, p := range config.Providers {

			models := make([]model.ProviderModelConfig, 0)

			if len(p.Models) > 0 {
				for _, m := range p.Models {
					newModel := model.ProviderModelConfig{
						ID:            m.ID,
						Name:          m.Name,
						Reasoning:     m.Reasoning,
						Input:         nil,
						Cost:          nil,
						ContextWindow: m.ContextWindow,
						MaxTokens:     m.MaxTokens,
					}
					if len(m.Input) > 0 {
						newModel.Input = m.Input
					} else {
						newModel.Input = []string{"text"}
					}
					if m.Name == "" {
						newModel.Name = m.ID
					}
					if m.ContextWindow == 0 {
						newModel.ContextWindow = 200000
					}
					if m.MaxTokens == 0 {
						newModel.MaxTokens = 8192
					}

					models = append(models, newModel)
				}
			}
			providers[p.Name] = &model.ProviderConfig{
				BaseURL:    p.BaseURL,
				APIKey:     p.APIKey,
				Auth:       getAuthOrDefault(p.Auth),
				AuthHeader: p.AuthHeader,
				API:        getAPIOrDefault(p.API, p.Name),
				Models:     models,
			}
		}
		return providers
	}

	// Fallback to legacy single provider
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://api.tokenpony.cn/v1"
	}
	defModel := getDefaultModel(config.Model)
	providerName := getProviderName(config.Provider, baseURL)

	auth := getAuthOrDefault(config.Auth)
	api := getAPIOrDefault(config.API, providerName)

	providers[providerName] = &model.ProviderConfig{
		BaseURL:    baseURL,
		APIKey:     config.APIKey,
		Auth:       auth,
		AuthHeader: false,
		API:        api,
		Models: []model.ProviderModelConfig{
			{
				ID:            defModel,
				Name:          defModel,
				Reasoning:     false,
				Input:         []string{"text"},
				Cost:          nil,
				ContextWindow: 200000,
				MaxTokens:     8192,
			},
		},
	}

	return providers
}

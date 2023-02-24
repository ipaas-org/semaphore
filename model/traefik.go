package model

type (
	MiddlewareRuleType string
)

const (
	MiddlewareRuleTypeBasicAuth MiddlewareRuleType = "basicAuth"
)

type (
	MiddlewareRuleer interface {
		Parse() string
	}

	MiddlewareRuleBasicAuth struct {
		HeaderField string `json:"headerField"`
	}
)

type (
	TraefikConfig struct {
		Routers     []Router
		Services    []Service
		Middlewares []Middleware
	}

	Router struct {
		ID          string
		Rule        string
		Middlewares []string
		Service     string
	}

	Service struct {
		ID         string
		ServersUrl []string
	}

	Middleware struct {
		ID       string
		RuleType MiddlewareRuleType
		Rule     MiddlewareRuleer
	}
)

func (m MiddlewareRuleBasicAuth) Parse() string {
	return "ciao"
}

func init() {

	rule := &MiddlewareRuleBasicAuth{
		HeaderField: "ciao",
	}

	m := Middleware{
		"ciao",
		MiddlewareRuleTypeBasicAuth,
		rule,
	}

	r := m.Rule
}

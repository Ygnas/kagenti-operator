package injector

// Label constants used by the precedence evaluator.
const (
	// Per-sidecar workload labels — set value to "false" to disable injection
	LabelEnvoyProxyInject   = "kagenti.io/envoy-proxy-inject"
	LabelSpiffeHelperInject = "kagenti.io/spiffe-helper-inject"

	// LabelClientRegistrationInject — legacy sidecar opt-in: set to "true" to inject;
	// default is operator-managed Keycloak credentials (no sidecar).
	LabelClientRegistrationInject = "kagenti.io/client-registration-inject"
)

// AuthBridge deployment modes. Selected per workload via AgentRuntime
// CR `Spec.AuthBridgeMode`, falling back to the namespace
// `authbridge-runtime-config` ConfigMap's `mode` field, then to
// ModeProxySidecar as the cluster-wide default.
const (
	ModeEnvoySidecar = "envoy-sidecar" // iptables + Envoy + ext_proc
	ModeProxySidecar = "proxy-sidecar" // default: HTTP_PROXY env + authbridge proxy
	ModeWaypoint     = "waypoint"      // standalone deployment (not injected)

	// Container name for proxy-sidecar mode
	AuthBridgeProxyContainerName = "authbridge-proxy"

	// Identity type constants
	IdentityTypeSpiffe         = "spiffe"
	ClientAuthTypeFederatedJWT = "federated-jwt"
)

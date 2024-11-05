package actions

// Factory defines an interface for creating new Action instances with specific configurations.
// This interface is designed to support dynamic and flexible action instantiation, allowing
// actions to be configured at runtime using a provided configuration map.
type Factory interface {
	// New creates a new instance of an Action based on the provided configuration.
	// The method takes a configuration map that specifies the parameters required
	// for the action's initialization, enabling customized behavior.
	//
	// Parameters:
	//   - cfg: a map containing configuration options as key-value pairs, where
	//     keys are strings representing configuration names, and values are of any type.
	//
	// Returns:
	//   - Action: an instance of an Action initialized with the provided configuration.
	//   - error: nil if the action is created successfully, or an error describing
	//     the issue if creation fails.
	New(cfg map[string]any) (Action, error)
}

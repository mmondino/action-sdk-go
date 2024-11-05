package events

// Event represents an object that can hold dynamic fields as key-value pairs.
// This interface is designed for use in event-processing pipelines where each
// processing node may need to add or retrieve specific fields in the event.
// The flexibility provided by this interface supports a variety of use cases
// by enabling nodes to enrich or modify the event data.
type Event interface {
	// AddField adds a new field to the event, or updates the value of an existing field.
	// The field is identified by the specified key, which is associated with the given value.
	//
	// Parameters:
	//   - key: a string representing the unique identifier for the field.
	//   - value: any type of data associated with the key.
	AddField(key string, value any)

	// GetField retrieves the value associated with a specific key in the event.
	// If the key exists, the value is returned along with a boolean 'true';
	// if the key is not found, it returns 'nil' and 'false'.
	//
	// Parameters:
	//   - key: a string representing the unique identifier for the field.
	//
	// Returns:
	//   - any: the value associated with the specified key, or nil if not found.
	//   - bool: true if the key exists, false otherwise.
	GetField(key string) (any, bool)
}

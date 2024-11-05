package actions

import "github.com/mmondino/action-sdk-go/events"

// Action defines a processing unit in a pipeline that performs a specific operation on an event.
// Each Action implementation should encapsulate its own logic to process the event data
// as needed, enabling a modular and extensible approach to event handling.
type Action interface {
	// Execute applies the specific action logic to the provided event.
	// This method receives an Event, modifies or reads its data as required by
	// the action logic, and returns an error if any issue occurs during execution.
	//
	// Parameters:
	//   - event: an Event instance on which the action will operate.
	//
	// Returns:
	//   - error: nil if the action completes successfully, or an error describing
	//     the issue if the action fails.
	Execute(event events.Event) error
}

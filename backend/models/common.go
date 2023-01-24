package models

import "reflect"

// Helper function to have a reusable way to overwrite a model's properties with another model's properties. null fields from the sending model are ignored.
func OverwriteModel[GenericModel any](receivingModel GenericModel, sendingModel GenericModel) {
	receiverValue := reflect.ValueOf(receivingModel).Elem()
	senderValue := reflect.ValueOf(sendingModel).Elem()

	// loop through the fields of the playlist
	for i := 0; i < receiverValue.NumField(); i++ {
		modelField := receiverValue.Field(i)
		newModelField := senderValue.Field(i)

		if !newModelField.IsZero() {
			// set the playlist field to the newPlaylist field
			modelField.Set(newModelField)
		}
	}
}

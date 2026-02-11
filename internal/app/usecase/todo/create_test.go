package todo

import (
	"testing"
)

func TestCreate_Handle_SuccessCases_TableDriven(t *testing.T) {
	//tests := []struct {
	//	name      string
	//	title     string
	//	completed bool
	//	want      CreateOutput
	//}{
	//	{
	//		name:      "creates todo when completed true",
	//		title:     "example title",
	//		completed: true,
	//		want:      CreateOutput{ID: "123", Title: "example title", Completed: true},
	//	},
	//	{
	//		name:      "creates todo when completed false",
	//		title:     "another title",
	//		completed: false,
	//		want:      CreateOutput{ID: "456", Title: "another title", Completed: false},
	//	},
	//}
	//
	//for _, tc := range tests {
	//	t.Run(tc.name, func(t *testing.T) {
	//		m := new(mockRepo)
	//		m.On("Create", tc.title, tc.completed).Return(tc.want.ID, nil).Once()
	//
	//		uc := NewCreate(m)
	//		_, err := uc.Handle(CreateInput{Title: tc.title, Completed: tc.completed})
	//
	//		assert.NoError(t, err)
	//	}
	//}
}

package sortutil

import "testing"

func TestOrder_String(t *testing.T) {
	tests := []struct {
		name string
		o    Order
		want string
	}{
        {
            name: "Asc",
            o: Ascending,
            want: "Ascending",
        },
        {
            name: "Desc",
            o: Descending,
            want: "Descending",
        },
        {
            name: "AscFold",
            o: AscendingCaseInsensitive,
            want: "AscendingCaseInsensitive",
        },
        {
            name: "DescFold",
            o: DescendingCaseInsensitive,
            want: "DescendingCaseInsensitive",
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("Order.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

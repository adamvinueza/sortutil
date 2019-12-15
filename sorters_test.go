package sortutil

import (
	"reflect"
	"sort"
	"testing"
)

func Test_stringAscending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stringAscending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("stringAscending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringDescending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stringDescending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("stringDescending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringInsensitiveAscending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stringInsensitiveAscending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("stringInsensitiveAscending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringInsensitiveDescending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stringInsensitiveDescending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("stringInsensitiveDescending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boolAscending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := boolAscending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("boolAscending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boolDescending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := boolDescending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("boolDescending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intAscending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := intAscending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("intAscending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intDescending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := intDescending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("intDescending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uintAscending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := uintAscending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("uintAscending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uintDescending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := uintDescending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("uintDescending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floatAscending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := floatAscending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("floatAscending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floatDescending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := floatDescending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("floatDescending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeAscending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := timeAscending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("timeAscending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeDescending_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := timeDescending{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("timeDescending.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverser_Len(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := reverser{
				sorter: tt.fields.sorter,
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("reverser.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverser_Less(t *testing.T) {
	type fields struct {
		sorter *sorter
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := reverser{
				sorter: tt.fields.sorter,
			}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("reverser.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		slice  interface{}
		getter getter
		order  Order
	}
	tests := []struct {
		name string
		args args
		want *sorter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new(tt.args.slice, tt.args.getter, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("new() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		slice  interface{}
		getter getter
		order  Order
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Sort(tt.args.slice, tt.args.getter, tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Sort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAsc(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Asc(tt.args.slice); (err != nil) != tt.wantErr {
				t.Errorf("Asc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDesc(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Desc(tt.args.slice); (err != nil) != tt.wantErr {
				t.Errorf("Desc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCiAsc(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CiAsc(tt.args.slice); (err != nil) != tt.wantErr {
				t.Errorf("CiAsc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCiDesc(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CiDesc(tt.args.slice); (err != nil) != tt.wantErr {
				t.Errorf("CiDesc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type A struct {
	B string
	C int
}

var aa = []A{
	A{B: "B", C: 2},
	A{B: "A", C: 1},
	A{B: "C", C: 3},
}

var bb = []A{
	A{B: "B", C: 2},
	A{B: "A", C: 1},
	A{B: "C", C: 3},
}

func TestByField(t *testing.T) {
	type args struct {
		slice interface{}
		name  string
		order Order
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		expectedOrder []interface{}
	}{
		{
			name: "Sorting ok",
			args: args{
				slice: aa,
				name:  "B",
				order: ascending,
			},
			wantErr: false,
			expectedOrder: []interface{}{
				bb[1],
				bb[0],
				bb[2],
			},
		},
		{
			name: "Sorting by nonexistent field not ok",
			args: args{
				slice: aa,
				name:  "dog",
				order: ascending,
			},
			wantErr: true,
		},
		{
			name: "Sorting by field case insensitive not ok",
			args: args{
				slice: aa,
				name:  "b",
				order: ascending,
			},
			wantErr: true,
		},
		{
			name: "Sorting by nonexistent field (space in name) not ok",
			args: args{
				slice: aa,
				name:  " B",
				order: ascending,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ByField(tt.args.slice, tt.args.name, tt.args.order)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("ByField() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestByCiField(t *testing.T) {
	type args struct {
		slice interface{}
		name  string
		order Order
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		expectedOrder []interface{}
	}{
		{
			name: "Sorting by field ok",
			args: args{
				slice: aa,
				name:  "B",
				order: ascending,
			},
			wantErr: false,
			expectedOrder: []interface{}{
				bb[1],
				bb[0],
				bb[2],
			},
		},
		{
			name: "Sorting by case insensitive field ok",
			args: args{
				slice: aa,
				name:  "b",
				order: ascending,
			},
			wantErr: false,
			expectedOrder: []interface{}{
				bb[1],
				bb[0],
				bb[2],
			},
		},
		{
			name: "Sorting by nonexistent field not ok",
			args: args{
				slice: aa,
				name:  "dog",
				order: ascending,
			},
			wantErr: true,
		},
		{
			name: "Sorting by nonexistent field (space in name) not ok",
			args: args{
				slice: aa,
				name:  " b",
				order: ascending,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ByCiField(tt.args.slice, tt.args.name, tt.args.order)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("ByCiField() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				for i, e := range tt.args.slice.([]A) {
					if tt.expectedOrder[i] != e {
						t.Errorf("Expected %v, found %v", tt.expectedOrder[i], e)
					}
				}
			}
		})
	}
}

func TestByIndex(t *testing.T) {
	type args struct {
		slice interface{}
		index int
		order Order
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ByIndex(tt.args.slice, tt.args.index, tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("ByIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.args.slice)
		})
	}
}

func TestReverseInterface(t *testing.T) {
	type args struct {
		s sort.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInterface(tt.args.s)
		})
	}
}

func TestSortReverseInterface(t *testing.T) {
	type args struct {
		s sort.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortReverseInterface(tt.args.s)
		})
	}
}

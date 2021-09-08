package main

import (
	"testing"
)

func Test_task_calculate(t1 *testing.T) {
	type fields struct {
		id, x, y, result, check int
		op                      string
		err                     error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "addition", fields: fields{
			id:    0,
			x:     2,
			y:     10,
			check: 12,
			op:    "+"},
		},
		{name: "subtraction", fields: fields{
			id:    1,
			x:     10,
			y:     12,
			check: -2,
			op:    "-"},
		},
		{name: "multiplication", fields: fields{
			id:    2,
			x:     100,
			y:     150,
			check: 15000,
			op:    "*"},
		},
		{name: "division", fields: fields{
			id:    3,
			x:     100,
			y:     2,
			check: 50,
			op:    "/"},
		},
		{name: "unknown operation", fields: fields{
			id: 4,
			x:  7,
			y:  8,
			op: "UnknownOp"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &task{
				id:     tt.fields.id,
				x:      tt.fields.x,
				y:      tt.fields.y,
				result: tt.fields.result,
				op:     tt.fields.op,
				err:    tt.fields.err,
			}
			t.calculate()
			if t.result != tt.fields.check {
				t1.Errorf("Error in operation %s. Result is %d. Expected %d\n", t.op, t.result, tt.fields.check)
			}
			if t.op == "UnknownOp" && t.err == nil {
				t1.Errorf("No error on UnknownOp operation")
			}
		})
	}
}

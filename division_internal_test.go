package division

/*
# ============================
# Author: Nitika Mishra
# ============================
*/

import ( 
	"testing"
)

func TestDivision(t *testing.T) {
    tests := []struct{
        n      int
        d      int
        q      int
	    r      int
        err    error
    } {
        { n: -6, d: -3, q: 2, r: 0, err: nil },
        { n: 55, d: 17, q: 3, r: 4, err: nil},
        { n: 55, d: 0, q: 0, r: 0, err: getDivisionByZeroError()},
    }

    for _, test := range tests {
        Q, R, er := divide(test.n, test.d)

        //Interface values are comparable. operands operator won't work 
        if er != nil || test.err != nil {
            var op_err, t_err string
            if er != nil {
                op_err = er.Error()
            }
            if test.err != nil {
                t_err = test.err.Error()
            }
            assertEqual(t, op_err, t_err, "error")
        } 

        assertEqual(t, test.q, Q, "quotient")
	    assertEqual(t, test.r, R, "reminder")
    }
}

func assertEqual(t *testing.T, a interface{}, b interface{}, dtype string) {
  	if a != b {
		t.Fatalf("expected %s: %d != actual %s: %d", dtype, a, dtype, b)
	}
}
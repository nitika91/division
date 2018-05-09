package division

/*
# ============================
# test inputdata validity
# only integers are valid
# ============================
*/

func testInput(val interface{}) (a int, err error){
	switch val.(type) {
		case bool:
			err = getInputBoolError()
			break
		case int:
			break
		case int32:
			err = getInputInt32Error()
			break
		case int64:
			err = getInputInt64Error()
			break
		case string:
			err = getInputStringError()
			break
	 	case float32:
			err = getInputFloat32Error()
			break
		case float64:
        	err = getInputFloat64Error()
	}
	if err == nil {
		 if num, ok := val.(int); ok == false {
		 	err = getInterfaceCastError()
		 } else {
			a = num
		 }
	}
return
}

/*
# ============================
# Func to divide two numbers and
# get the quotient and remender
# ============================
*/

func divide(num, denom interface{}) (q int, rem int, err error) {
	var n, d int
	// check input validity then process division
	if n, err = testInput(num); err == nil {
		if d, err = testInput(denom); err == nil {
			sign := 1

			//get abstruct value and result signature
			n, d, sign = getUnsignedNumbers(n, d, sign)

			//handle exception and throw error
			if d == 0 {
				err = getDivisionByZeroError()
				return
			}

			if n < d {
				rem = n
				return
			}

			tmp_ans := 1
			tmp_d := d

			// use binary search approach to get the nearest quotient
			for tmp_d <= n {
				tmp_d = tmp_d << 1
				tmp_ans = tmp_ans << 1
			}
			tmp_d = tmp_d >> 1
			tmp_ans = tmp_ans >> 1

			rem = n - tmp_d

			// recursive call to get the additonal quotient
			if rem > d {
				q, rem, err =  divide(rem, d)
			}

			//handle negative number division
			q = (tmp_ans + q) * sign
			rem = rem * sign
		}
	}
return
}

/*
# ============================
# get the abstruct value
# of number
# ============================
*/
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

/*
# ============================
# check the if any numerator
# or denomenator is negative
# ============================
*/

func getUnsignedNumbers (n, d, sign int) (int, int, int) {
	if n < 0  {
	   n = abs(n)
	   sign = -sign
	}
	if d < 0 {
		d = abs(d)
		sign = -sign
	}
	return n, d, sign
}

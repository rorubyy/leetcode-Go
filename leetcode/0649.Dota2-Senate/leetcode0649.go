package leetcode

func predictPartyVictory(senate string) string {
	qr, qd := make([]int, 0), make([]int, 0)
	for idx, val := range senate {
		if val == 'R' {
			qr = append(qr, idx)
		} else {
			qd = append(qd, idx)
		}
	}
	i := 0
	for len(qr) != 0 && len(qd) != 0 {
		r_idx, d_idx := qr[0], qd[0]
		if r_idx > d_idx {
			qd = append(qd, i+len(senate))
		} else {
			qr = append(qr, i+len(senate))
		}
		qd = qd[1:]
		qr = qr[1:]
		i++
	}
	if len(qr)>0{
		return "Radiant"
	}else{
		return "Dire"
	}
}

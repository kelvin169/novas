package novas

import (
	"errors"
	"time"
)

// Compute the time of rise of a body above the dip angle.
// Returns an error if the body doesn't rise above dip within 24 hours from the time given.
func (p *Body) Rise(t Time, geo *Place, dip float64, precision time.Duration, refr RefractType) (Time, BodyTopoData, error) {

	alt1 := p.Topo(t, geo, refr).Alt
	t1 := t
	t2 := t
	found := false
	for i := 0; i < 48; i++ {
		t2.Time = t1.Add(30 * time.Minute)
		alt2 := p.Topo(t2, geo, refr).Alt
		if alt1 < dip && alt2 > dip {
			found = true
			break
		}
		t1 = t2
		alt1 = alt2
	}
	if !found {
		return Time{}, BodyTopoData{}, errors.New("No rise above dip in the next 24 hours")
	}

	var topo BodyTopoData
	tt := t1
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		topo = p.Topo(tt, geo, refr)
		if topo.Alt > dip {
			t2 = tt
		} else {
			t1 = tt
		}
	}
	return tt, topo, nil
}

// Compute the time of set of a body below the dip angle.
// Returns an error if the body doesn't set below dip within 24 hours from the time given.
func (p *Body) Set(t Time, geo *Place, dip float64, precision time.Duration, refr RefractType) (Time, BodyTopoData, error) {

	alt1 := p.Topo(t, geo, refr).Alt
	t1 := t
	t2 := t
	found := false
	for i := 0; i < 48; i++ {
		t2.Time = t1.Add(30 * time.Minute)
		alt2 := p.Topo(t2, geo, refr).Alt
		if alt1 > dip && alt2 < dip {
			found = true
			break
		}
		t1 = t2
		alt1 = alt2
	}
	if !found {
		return Time{}, BodyTopoData{}, errors.New("No set below dip in the next 24 hours")
	}

	var topo BodyTopoData
	tt := t1
	for t2.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t2.Sub(t1.Time) / 2)
		topo = p.Topo(tt, geo, refr)
		if topo.Alt < dip {
			t2 = tt
		} else {
			t1 = tt
		}
	}
	return tt, topo, nil
}

// Compute the time of highest position in the sky of a body.
// Returns an error if the body doesn't goes up then down within 24 hours from the time given.
func (p *Body) High(t Time, geo *Place, precision time.Duration, refr RefractType) (Time, BodyTopoData, error) {

	alt1 := p.Topo(t, geo, refr).Alt
	alt2 := alt1
	alt3 := alt1
	t1 := t
	t2 := t
	t3 := t
	found := false
	for i := 0; i < 48; i++ {
		t3.Time = t2.Add(30 * time.Minute)
		alt3 = p.Topo(t3, geo, refr).Alt
		if i > 1 && alt2 > alt3 && alt2 > alt1 {
			found = true
			break
		}
		t1, t2 = t2, t3
		alt1, alt2 = alt2, alt3
	}

	if !found {
		return Time{}, BodyTopoData{}, errors.New("No high point in the next 24 hours")
	}

	var topo BodyTopoData
	tt := t1
	for t3.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t3.Sub(t1.Time) / 2)
		topo = p.Topo(tt, geo, refr)
		if alt1 < alt3 {
			t1 = tt
			alt1 = topo.Alt
		} else {
			t3 = tt
			alt3 = topo.Alt
		}
	}
	return tt, topo, nil
}

// Compute the time of lowest position in the sky of a body.
// Returns an error if the body doesn't goes down then up within 24 hours from the time given.
func (p *Body) Low(t Time, geo *Place, precision time.Duration, refr RefractType) (Time, BodyTopoData, error) {

	alt1 := p.Topo(t, geo, refr).Alt
	alt2 := alt1
	alt3 := alt1
	t1 := t
	t2 := t
	t3 := t
	found := false
	for i := 0; i < 48; i++ {
		t3.Time = t2.Add(30 * time.Minute)
		alt3 = p.Topo(t3, geo, refr).Alt
		if i > 1 && alt2 < alt3 && alt2 < alt1 {
			found = true
			break
		}
		t1, t2 = t2, t3
		alt1, alt2 = alt2, alt3
	}

	if !found {
		return Time{}, BodyTopoData{}, errors.New("No low point in the next 24 hours")
	}

	var topo BodyTopoData
	tt := t1
	for t3.Sub(t1.Time) > precision {
		tt.Time = t1.Add(t3.Sub(t1.Time) / 2)
		topo = p.Topo(tt, geo, refr)
		if alt1 > alt3 {
			t1 = tt
			alt1 = topo.Alt
		} else {
			t3 = tt
			alt3 = topo.Alt
		}
	}
	return tt, topo, nil
}

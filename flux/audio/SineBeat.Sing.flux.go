// Generated by Flux, not meant for human consumption.  Editing may make it unreadable by Flux.

package audio

func (x *SineBeat) Sing() (x2 Audio) {
	var v *SineBeat
	var v2 *SineBeat
	var v3 float64
	var v4 float64
	var v5 Audio
	var v6 float64
	var v7 *SineBeat
	var v8 *SineBeat
	var v9 *NormalOsc
	var v10 float64
	var v11 *SineBeat
	var v12 *SineBeat
	var v13 *SineOsc
	var v14 Audio
	var v15 Audio
	v11 = x
	v12 = x
	v7 = x
	v8 = x
	v = x
	v2 = x
	x22 := &v11.Sine
	v13 = x22
	x3 := &v2.sineFreq
	v3 = *x3
	x4 := v13.SineFreq(v3)
	v14 = x4
	x5 := &v.Env
	v9 = x5
	x6 := &v8.beatFreq
	v10 = *x6
	x7 := &v7.beatWidth
	v4 = *x7
	x23 := v9.Osc(v10, v4)
	v5 = x23
	x8 := Mul(v14, v5)
	v15 = x8
	x9 := &v12.amp
	v6 = *x9
	x32 := v15.MulX(v6)
	x2 = x32
	return
}

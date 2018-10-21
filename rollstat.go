package rollstat

// FloatStat rolling statistics calculator for floating point numbers
type FloatStat struct {
	n                  float64
	base               float64
	varianceSum        float64
	varianceSumSquared float64
}

// IntStat rolling statistics calculator for integers
type IntStat struct {
	n                  int64
	base               int64
	varianceSum        int64
	varianceSumSquared int64
}

// Add add a variable to the rolling calculator
func (me *IntStat) Add(x int64) {
	if me.n == 0 {
		me.base = x
	}

	me.n = me.n + 1

	me.varianceSum += x - me.base
	me.varianceSumSquared += (x - me.base) * (x - me.base)
}

// Mean get the rolling calculator's mean
func (me *IntStat) Mean() int64 {
	return me.base + (me.varianceSum / me.n)
}

// Stddev get the rolling calculator's standard deviation
func (me *IntStat) Stddev() int64 {
	return (me.varianceSumSquared - ((me.varianceSum * me.varianceSum) / me.n)) / (me.n - 1)
}

// Add add a variable to the rolling calculator
func (me *FloatStat) Add(x float64) {
	if me.n == 0.0 {
		me.base = x
	}

	me.n = me.n + 1.0

	me.varianceSum += x - me.base
	me.varianceSumSquared += (x - me.base) * (x - me.base)
}

// Mean get the rolling calculator's mean
func (me *FloatStat) Mean() float64 {
	return me.base + (me.varianceSum / me.n)
}

// Stddev get the rolling calculator's standard deviation
func (me *FloatStat) Stddev() float64 {
	return (me.varianceSumSquared - ((me.varianceSum * me.varianceSum) / me.n)) / (me.n - 1.0)
}

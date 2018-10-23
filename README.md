# rollstat

A simple rolling average / variance calculator

Example:

    stat := &rollstat.IntStat{} // can use rollstat.FloatStat{} for floating point numbers

    stat.Add(2)
    stat.Add(3)
    stat.Add(-5)

    avg := stat.Mean()
    var := stat.Var()

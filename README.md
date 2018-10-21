# rollstat
Simple rolling average / standard deviation calculator

Example:

    stat := &rollstat.IntStat{} // can use rollstat.FloatStat{} for floating point numbers

    stat.Add(2)
    stat.Add(3)
    stat.Add(-5)

    avg := stat.Mean()
    std := stat.Stddev()
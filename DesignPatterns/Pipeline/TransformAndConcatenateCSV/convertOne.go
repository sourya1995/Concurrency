package main

func convertOne(in <-chan []string) chan []string {
	out := make(chan []string)
    go func(in <- chan []string) {
        defer close(out)
        for v := range in {
            v[0], v[1], v[2] = v[1], v[2], v[0]
			out <- v
        }
    }(in)
    return out
}
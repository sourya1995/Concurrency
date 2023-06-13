package main

func convertTwo(in <- chan []string) chan []string {
	out := make(chan []string)
    go func(in <- chan []string) {
        defer close(out)
        for c := range in {
            for i := 0; i < len(c); i++ {
                c[i] = "9" + c[i] + "0"
            }
			out <- c
        }
    }(in)
    return out
}
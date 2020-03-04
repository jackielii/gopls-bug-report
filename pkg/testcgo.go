package pkg

// int one(){ return 1; }
import "C"

func One() int {
	return int(C.one())
}

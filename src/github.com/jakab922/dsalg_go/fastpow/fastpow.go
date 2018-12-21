// Note that both functions can overflow.
package dsalg

const (
	Zero = int64(0)
	One  = int64(1)
)

func Fastpow(a, b int64) int64 {
	if a == Zero {
		return Zero
	}
	ret := One
	for b > Zero {
		if b&One == One {
			ret *= a
		}
		a *= a
		b >>= 1
	}
	return ret
}

func FastpowMod(a, b, m int64) int64 {
	a %= m
	if a == Zero {
		return Zero
	}
	ret := One
	for b > Zero {
		if b&One == One {
			ret = (ret * a) % m
		}
		a = (a * a) % m
		b >>= 1
	}
	return ret
}

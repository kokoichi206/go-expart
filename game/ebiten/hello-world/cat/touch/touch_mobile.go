//go:build android || ios

package touch

func isTouchMain() bool {
	return true
}

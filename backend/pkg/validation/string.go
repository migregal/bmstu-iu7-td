package validation

func IsHex(s string) bool {
    for _, b := range []byte(s) {
        if !(b >= '0' && b <= '9' || b >= 'a' && b <= 'f' || b >= 'A' && b <= 'F') {
            return false
        }
    }
    return true
}

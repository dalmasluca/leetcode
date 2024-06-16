func reverseVowels(s string) string {
    var w []byte
    var pos []int
    b := []byte(s)

    for i := 0; i < len(s); i++ {
        switch s[i] {
            case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
                w = append(w, s[i])
                pos = append(pos, i)
            break
        }
    }

    l := len(w)-1
    for i := 0; i < len(pos); i++ {
        b[pos[i]] = w[l-i]
    }
    return string(b)
}

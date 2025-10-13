func removeAnagrams(words []string) []string {
    last := sortString(words[0])
    i := 1
    for i < len(words) {
        curr := sortString(words[i])
        if curr == last {
            words = append(words[:i], words[i+1:]...)
            continue
        }
        last = curr
        i++
    }
    return words
}

func sortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
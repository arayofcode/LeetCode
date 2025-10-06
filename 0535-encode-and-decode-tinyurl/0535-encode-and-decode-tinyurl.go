import (
    "math/big"
)

type Codec struct {
    urls map[string]string
}


func Constructor() Codec {
    return Codec{
        urls: make(map[string]string),
    }
}

func toBase62(s string) string {
    var i big.Int
    i.SetBytes([]byte(s))
    return i.Text(62)
}

func (this *Codec) encode(longUrl string) string {
    encoded := toBase62(longUrl)
    this.urls[encoded] = longUrl
    return encoded
}

func (this *Codec) decode(shortUrl string) string {
    return this.urls[shortUrl]
}

package main

// based on following article.
// http://paulbourke.net/dataformats/asciiart/
//
// from black to white:
//
// $ @ B % 8 & W M # * o a h k b d p q w m Z O 0 Q L C J U Y X z c v u n x r j f t / \ | ( ) 1 { } [ ] ? - _ + ~ < > i ! l I
// ; : , " ^ ` ' .
//

var charmap = map[int]string{}

func init() {
	p := 0
	s := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
	for i := 0; i < 256; i++ {
		charmap[i] = s[p : p+1]
		if i%4 == 0 {
			p++
		}
	}
}

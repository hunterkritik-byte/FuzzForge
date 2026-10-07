package mutator

import (
	"math/rand"
	"strings"
)

var snippets = []string{
	"let x = 0; x++;",
	"let a = [1,2,3]; a.length;",
	"({a: 1, b: 2}).a;",
	"for (let i = 0; i < 4; i++) {}",
	"function f(x) { return x + 1; } f(1);",
	"try { throw 1; } catch (e) {}",
	"new Proxy({}, {});",
	"new ArrayBuffer(8);",
	"let s = 'fuzz'; s.repeat(2);",
}

func Mutate(seed string, r *rand.Rand) string {
	parts := strings.Fields(seed)
	switch r.Intn(4) {
	case 0:
		return seed + "
" + snippets[r.Intn(len(snippets))]
	case 1:
		return snippets[r.Intn(len(snippets))] + "
" + seed
	case 2:
		if len(parts) > 1 {
			i := r.Intn(len(parts))
			parts[i] = snippets[r.Intn(len(snippets))]
			return strings.Join(parts, " ")
		}
		return seed + "
" + snippets[r.Intn(len(snippets))]
	default:
		return seed + "
for(let i=0;i<" + string(rune('0'+r.Intn(8))) + ";i++){}"
	}
}

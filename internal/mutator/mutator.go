package mutator

import ("math/rand"; "strings")

var snippets=[]string{
	"let x=0; x++;","let a=[1,2,3]; a.length;",
	"({a:1,b:2}).a;","for(let i=0;i<4;i++){}",
	"function f(x){return x+1} f(1);","try{throw 1}catch(e){}",
	"new Proxy({},{});","new ArrayBuffer(8);",
	"let s='fuzz'; s.repeat(2);",
}

func Mutate(seed string,r *rand.Rand) string {
	parts:=strings.Fields(seed)
	switch r.Intn(4) {
	case 0:return seed+"\n"+snippets[r.Intn(len(snippets))]
	case 1:return snippets[r.Intn(len(snippets))]+"\n"+seed
	case 2:
		if len(parts)>1 { parts[r.Intn(len(parts))]=snippets[r.Intn(len(snippets))]; return strings.Join(parts," ") }
		return seed+"\n"+snippets[r.Intn(len(snippets))]
	default:return seed+"\nfor(let i=0;i<8;i++){}"
	}
}

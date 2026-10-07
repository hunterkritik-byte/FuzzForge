package generator

import "math/rand"

var atoms=[]string{"0","1","-1","NaN","Infinity","[]","{}","'fuzz'","new ArrayBuffer(16)"}

func Program(r *rand.Rand) string {
	n:=3+r.Intn(8); out:=""
	for i:=0;i<n;i++ { switch r.Intn(7) {
	case 0: out+="let a"+itoa(i)+"=["+atoms[r.Intn(len(atoms))]+"];\n"
	case 1: out+="function f"+itoa(i)+"(x){return x+"+atoms[r.Intn(len(atoms))]+";} f"+itoa(i)+"(0);\n"
	case 2: out+="let o"+itoa(i)+"={x:"+atoms[r.Intn(len(atoms))]+"}; o"+itoa(i)+".x;\n"
	case 3: out+="for(let i"+itoa(i)+"=0;i"+itoa(i)+"<3;i"+itoa(i)+"++){}\n"
	case 4: out+="try{throw "+atoms[r.Intn(len(atoms))]+";}catch(e){}\n"
	case 5: out+="new Proxy({},{});\n"
	default: out+="let s"+itoa(i)+"='fuzz'.repeat("+itoa(1+r.Intn(4))+");\n"
	}}
	return out
}
func itoa(n int) string { if n==0{return "0"}; b:=[]byte{}; for n>0{b=append([]byte{byte('0'+n%10)},b...);n/=10};return string(b) }

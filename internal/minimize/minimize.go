package minimize
import("context";"os";"os/exec";"time")
func Fails(engine,program string,t time.Duration)bool{p,e:=os.CreateTemp("","fuzzforge-min-*.js");if e!=nil{return false};name:=p.Name();defer os.Remove(name);if _,e=p.WriteString(program);e!=nil{return false};p.Close();ctx,cancel:=context.WithTimeout(context.Background(),t);defer cancel();e=exec.CommandContext(ctx,engine,name).Run();return e!=nil&&ctx.Err()==nil}
func Reduce(engine,program string,t time.Duration)string{lines:=split(program);for changed:=true;changed;{changed=false;for i:=0;i<len(lines);i++{c:=joinExcept(lines,i);if c!=""&&Fails(engine,c,t){lines=split(c);changed=true;break}}};return join(lines)}
func split(s string)[]string{var a []string;start:=0;for i,c:=range s{if c=='\n'{a=append(a,s[start:i]);start=i+1}};if start<len(s){a=append(a,s[start:])};return a}
func joinExcept(a []string,i int)string{o:="";for j,v:=range a{if j!=i{o+=v+"\n"}};return o}
func join(a []string)string{o:="";for _,v:=range a{o+=v+"\n"};return o}

package runner

import ("context";"fmt";"math/rand";"os";"os/exec";"path/filepath";"time";"github.com/hunterkritik-byte/FuzzForge/internal/generator";"github.com/hunterkritik-byte/FuzzForge/internal/mutator")

type Config struct{Engine,CorpusDir,ArtifactDir string;Runs int;Timeout time.Duration}
func execute(engine,path string,t time.Duration)([]byte,error,bool){ctx,cancel:=context.WithTimeout(context.Background(),t);defer cancel();out,err:=exec.CommandContext(ctx,engine,path).CombinedOutput();return out,err,ctx.Err()==context.DeadlineExceeded}
func Run(cfg Config) error{
	if cfg.Timeout<=0{cfg.Timeout=2*time.Second}; seeds,err:=os.ReadDir(cfg.CorpusDir);if err!=nil{return fmt.Errorf("read corpus: %w",err)};if len(seeds)==0{return fmt.Errorf("corpus is empty")};if err=os.MkdirAll(cfg.ArtifactDir,0755);err!=nil{return err}
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<cfg.Runs;i++{entry:=seeds[r.Intn(len(seeds))];data,err:=os.ReadFile(filepath.Join(cfg.CorpusDir,entry.Name()));if err!=nil{return err};program:=mutator.Mutate(string(data),r);if r.Intn(3)==0{program+=generator.Program(r)}
		path:=filepath.Join(cfg.ArtifactDir,fmt.Sprintf("case-%06d.js",i));if err=os.WriteFile(path,[]byte(program),0644);err!=nil{return err};out,runErr,timedOut:=execute(cfg.Engine,path,cfg.Timeout)
		if timedOut{dst:=filepath.Join(cfg.ArtifactDir,fmt.Sprintf("timeout-%06d.js",i));_ = os.WriteFile(dst,[]byte(program),0644);fmt.Printf("[timeout] %s\n",dst);continue}
		if runErr!=nil{dst:=filepath.Join(cfg.ArtifactDir,fmt.Sprintf("crash-%06d.js",i));_ = os.WriteFile(dst,[]byte(program),0644);_ = os.WriteFile(dst+".stderr",out,0644);fmt.Printf("[crash] %s\n",dst);continue}
		if i%10==0{dst:=filepath.Join(cfg.CorpusDir,fmt.Sprintf("generated-%06d.js",i));if _,e:=os.Stat(dst);os.IsNotExist(e){_ = os.WriteFile(dst,[]byte(program),0644);seeds,_=os.ReadDir(cfg.CorpusDir);fmt.Printf("[corpus+] %s\n",dst)}}
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"sync/atomic"
	pb "tiflash-auto-scaling/supervisor_proto"
)

var assignTenantID atomic.Value
var pid atomic.Int32
var mu sync.Mutex
var ch = make(chan *pb.AssignRequest)

func AssignTiFlash(in *pb.AssignRequest) (*pb.Result, error) {
	log.Printf("received assign request by: %v", in.GetTenantID())
	if assignTenantID.Load() == nil {
		mu.Lock()
		defer mu.Unlock()
		if assignTenantID.Load() == nil {
			assignTenantID.Store(in.GetTenantID())
			ch <- in
			return &pb.Result{HasErr: false, ErrInfo: ""}, nil
		}
	} else if assignTenantID.Load().(string) == in.GetTenantID() {
		return &pb.Result{HasErr: false, ErrInfo: ""}, nil
	}
	return &pb.Result{HasErr: true, ErrInfo: "TiFlash has been occupied by a tenant"}, nil
}

func UnassignTiFlash(in *pb.UnassignRequest) (*pb.Result, error) {
	log.Printf("received unassign request by: %v", in.GetAssertTenantID())
	if in.AssertTenantID == assignTenantID.Load() {
		mu.Lock()
		defer mu.Unlock()
		if in.AssertTenantID == assignTenantID.Load() && pid.Load() != 0 {
			assignTenantID.Store("")
			cmd := exec.Command("kill", "-9", fmt.Sprintf("%v", pid.Load()))
			err := cmd.Run()
			pid.Store(0)
			if err != nil {
				return &pb.Result{HasErr: true, ErrInfo: err.Error()}, err
			}
			return &pb.Result{HasErr: false, ErrInfo: ""}, nil
		}
	}
	return &pb.Result{HasErr: true, ErrInfo: "TiFlash is not assigned to this tenant"}, nil

}

func TiFlashMaintainer() {
	for true {
		in := <-ch
		configFile := fmt.Sprintf("tiflash-%s.toml", in.GetTenantID())
		f, err := os.Create(configFile)
		if err != nil {
			log.Fatalf("create config file failed: %v", err)
		}
		defer f.Close()
		_, err = f.WriteString(in.GetTenantConfig())

		if err != nil {
			log.Fatalf("write config file failed: %v", err)
		}

		for assignTenantID.Load().(string) == in.GetTenantID() {
			cmd := exec.Command("./tiflash", "server", "--config-file", configFile)
			err = cmd.Start()
			pid.Store(int32(cmd.Process.Pid))
			if err != nil {
				log.Printf("start tiflash failed: %v", err)
			}
			err = cmd.Wait()
			log.Printf("tiflash exited: %v", err)
		}
	}
}
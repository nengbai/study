package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"

	client "github.com/heketi/heketi/client/api/go-client"
	// api"github.com/heketi/heketi/pkg/glusterfs/api"
)

type Options struct {
	Url  string
	User string
	Key  string
}
type Cmd struct {
	Path         string
	Args         []string
	Env          []string
	Dir          string
	Stdin        io.Reader
	Stdout       io.Writer
	Stderr       io.Writer
	ExtraFiles   []*os.File
	SysProcAttr  *syscall.SysProcAttr
	Process      *os.Process
	ProcessState *os.ProcessState
}

func Command(name string, arg ...string) *Cmd {

	var outInfo bytes.Buffer
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("select * from user")
	cmd.Stdout = &outInfo

	cmd.Run()

	// fmt.Println(out.String())
	return nil
}

func main() {
	options := Options{Url: "http://10.100.104.21:8080", User: "admin", Key: "admin@123"}
	// Create a client
	clientObj := client.NewClient(options.Url, options.User, options.Key)

	// Get a cluster id
	clusterListRes, _ := clientObj.ClusterList()
	if len((*clusterListRes).Clusters) == 0 {
		fmt.Println("don't have a cluster")
		return
	}
	// clusterList := (*clusterListRes).Clusters
	// clusterInfo, _ := clientObj.ClusterInfo(clusterList[0])

	//Create node
	// nodeReq := &api.NodeAddRequest{}
	// n := 0
	// nodeReq.ClusterId = clusterInfo.Id
	// nodeReq.Hostnames.Manage = []string{"cent" + fmt.Sprintf("%v", 1)}
	// nodeReq.Hostnames.Storage = []string{"storage" + fmt.Sprintf("%v", 1)}
	// nodeReq.Zone = n + 1

	// // Add node
	// node, err := clientObj.NodeAdd(nodeReq)
	// fmt.Println(node,err)

	// for _, node := range clusterInfo.Nodes {
	// nodeInfo, _ := clientObj.NodeInfo(node)

	//}
	volumeInfo, _ := clientObj.VolumeList()
	for i := range *&volumeInfo.Volumes {
		fmt.Println(*&volumeInfo.Volumes[i])
		// clientObj.VolumeInfo(*&volumeInfo.Volumes[i])
	}

}

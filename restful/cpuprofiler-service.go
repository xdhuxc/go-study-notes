package main

import (
	"github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"runtime/pprof"
)

type ProfilingService struct {
	rootPath   string
	cpuProfile string
	cpuFile    *os.File
}

func NewProfileService(rootPath string, outputFileName string) *ProfilingService {
	ps := new(ProfilingService)
	ps.rootPath = rootPath
	ps.cpuProfile = outputFileName
	return ps
}

func (p *ProfilingService) startProfiler(req *restful.Request, resp *restful.Response) {
	if p.cpuProfile != "" {
		if _, err := io.WriteString(resp.ResponseWriter, "[restful] CPU profiling already running."); err != nil {
			log.Infoln(err.Error())
		}
		return
	}
	cpuFile, err := os.Create(p.cpuProfile)
	if err != nil {
		log.Fatal(err)
	}

	p.cpuFile = cpuFile
	err = pprof.StartCPUProfile(cpuFile)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.WriteString(resp.ResponseWriter, "[restful] CPU profiling started, writing on:"+p.cpuProfile); err != nil {
		log.Info(err)
	}
}

func (p *ProfilingService) stopProfiler(req *restful.Request, resp *restful.Response) {
	if p.cpuFile == nil {
		if _, err := io.WriteString(resp.ResponseWriter, "[restful] CPU profiling not active."); err != nil {
			log.Info(err)
		}
		return
	}
	pprof.StopCPUProfile()
	if err := p.cpuFile.Close(); err != nil {
		log.Info(err)
	}
	p.cpuFile = nil
	if _, err := io.WriteString(resp.ResponseWriter, "[restful] CPU profiling stopped, closing:"+p.cpuProfile); err != nil {
		log.Info(err)
	}
}

func (p *ProfilingService) AddWebServiceTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path(p.rootPath).Consumes("*/*").Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/start").To(p.startProfiler))
	ws.Route(ws.GET("/stop").To(p.stopProfiler))
	container.Add(ws)
}

func main() {
	NewProfileService("./", "cpu.prof").AddWebServiceTo(restful.DefaultContainer)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

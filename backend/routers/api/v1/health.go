package v1

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/app"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

type message struct {
	Status string `json:"status"`
	Info   string `json:"info, omitempty"`
}

func HealthCheck(c *gin.Context) {
	msg := "OK"
	app.SendResponse(c, http.StatusOK, nil, message{
		Status: msg,
	})
}

func DiskCheck(c *gin.Context) {
	usage, _ := disk.Usage("/")
	
	usedMB := int(usage.Used) / MB
	usedGB := int(usage.Used) / GB
	totalMB := int(usage.Total) / MB
	totalGB := int(usage.Total) / GB
	usedPercent := int(usage.UsedPercent)
	
	status := http.StatusOK
	text := "OK"
	
	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	
	info := fmt.Sprintf("Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", usedMB, usedGB, totalMB, totalGB, usedPercent)
	app.SendResponse(c, status, nil, message{
		Status: text,
		Info:   info,
	})
}

func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)
	
	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15
	
	status := http.StatusOK
	text := "OK"
	
	if l5 >= float64(cores) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	
	info := fmt.Sprintf("Load average: %.2f, %.2f, %.2f | Cores: %d", l1, l5, l15, cores)
	app.SendResponse(c, status, nil, message{
		Status: text,
		Info:   info,
	})
}

func RAMCheck(c *gin.Context) {
	usage, _ := mem.VirtualMemory()
	
	usedMB := int(usage.Used) / MB
	usedGB := int(usage.Used) / GB
	totalMB := int(usage.Total) / MB
	totalGB := int(usage.Total) / GB
	usedPercent := int(usage.UsedPercent)
	
	status := http.StatusOK
	text := "OK"
	
	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	
	info := fmt.Sprintf("Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", usedMB, usedGB, totalMB, totalGB, usedPercent)
	app.SendResponse(c, status, nil, message{
		Status: text,
		Info:   info,
	})
}

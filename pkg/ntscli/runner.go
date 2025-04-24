package ntscli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// Define a struct to match the expected JSON body
type PostData struct {
	Module string `json:"module"`
	Cmd    string `json:"command"`
	Data   string `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	origin := r.Header.Get("Origin")

	// List of allowed origins (adjust as needed)
	allowedOrigins := map[string]bool{
		"http://localhost:8000":   true,
		"http://10.1.10.205:8000": true,
		"http://10.1.10.93:29020": true,
		"http://10.1.10.93:8000":  true,
		"http://localhost:32930":  true,
	}

	if allowedOrigins[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		// Optionally reject or set no CORS header
		http.Error(w, "Origin not allowed", http.StatusForbidden)
		return
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("posts handler called")
	// Parse form data (for application/x-www-form-urlencoded)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	data := PostData{
		Module: r.FormValue("module"),
		Cmd:    r.FormValue("command"),
		Data:   r.FormValue("data"),
	}

	serverResponse := handleResponse(data.Module, data.Cmd, data.Data)

	// Print received data
	fmt.Printf("Received: %+v\n", data)

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"msg":    serverResponse,
	})
}

var pidFile string = "/tmp/nts.pid"

func RunServers() {
	pid := os.Getpid()
	os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)

	// Create separate ServeMux and http.Server for each server
	jsMux := http.NewServeMux()
	jsMux.Handle("/", http.FileServer(http.Dir("./pkg/ntscli/web")))
	jsServer := &http.Server{Addr: ":8000", Handler: jsMux}

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/posts", handler)
	apiServer := &http.Server{Addr: ":8080", Handler: apiMux}

	go func() {
		fmt.Println("Serving static files on :8000")
		if err := jsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("8000 error:", err)
		}
	}()

	go func() {
		fmt.Println("Serving API on :8080")
		if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("8080 error:", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jsServer.Shutdown(ctx)
	apiServer.Shutdown(ctx)
	fmt.Println("Servers stopped gracefully.")
	// Call Shutdown on your servers
	// e.g., apiServer.Shutdown(ctx), staticServer.Shutdown(ctx)
	fmt.Println("Servers stopped gracefully.")
}

func StopServers() {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		fmt.Println("Could not read PID file:", err)
		return
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Println("Invalid PID:", err)
		return
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Could not find process:", err)
		return
	}
	// Send SIGTERM (or SIGINT)
	proc.Signal(syscall.SIGTERM)
	fmt.Println("Sent stop signal to server process")
}

func handleResponse(module string, command string, data string) string {
	ReadDeviceConfig()

	if DeviceHasNtpServer() == 0 {
		switch module {
		case "ntp":
			log.Println(data)
			writeNtpServerStatus(data)
			newStatus := formatNtpServerSTATUS()

			log.Println(newStatus)
			return newStatus
		case "yourmom":
			log.Println("your mom case")
		default:
			log.Println("def case")
		}
	}
	return "your mom"
}

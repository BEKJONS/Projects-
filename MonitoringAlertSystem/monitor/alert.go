package monitor

import (
	"fmt"
	"log"
	"mas/bot"
	"os"
	"time"
)

// CheckThresholds function compares the collected metrics with the defined thresholds
// and triggers an alert if any threshold is exceeded
func CheckThresholds(metrics Metrics, config *Config) {
	// Check if CPU usage exceeds the threshold
	if metrics.CPUUsage > float64(config.Thresholds.CPUUsage) {
		alert("CPU Usage", metrics.CPUUsage) // Trigger an alert for CPU usage
	}

	// Check if memory usage exceeds the threshold
	if metrics.MemoryUsage > float64(config.Thresholds.MemoryUsage) {
		alert("Memory Usage", metrics.MemoryUsage) // Trigger an alert for memory usage
	}

	// Check if disk I/O exceeds the threshold
	if metrics.DiskIO > float64(config.Thresholds.DiskIO) {
		alert("Disk I/O", metrics.DiskIO) // Trigger an alert for disk I/O
	}
}

// alert function logs an alert message when a metric exceeds its threshold
func alert(metricType string, value float64) {
	// Create an alert message with the metric type and its current value
	message := fmt.Sprintf("ALERT: %s exceeded threshold with value: %.2f", metricType, value)
	bots, err := bot.ConnectToTelegram()
	if err != nil {
		panic(err)
	}

	err = bot.SendMessageToOneUser(bots, message)
	if err != nil {
		// Handle the error
	}
	logToFile(message) // Log the alert message to a file
}

// LogMetrics function logs the current metrics data to a log file
func LogMetrics(metrics Metrics) {
	// Create a log message with timestamp, CPU, memory, and disk I/O usage
	message := fmt.Sprintf("Metrics at %s - CPU: %.2f%%, Memory: %.2f%%, Disk I/O: %.2f MB/s",
		time.Now().Format(time.RFC3339), metrics.CPUUsage, metrics.MemoryUsage, metrics.DiskIO)
	logToFile(message) // Log the metrics message to a file
}

// logToFile function writes log messages to a log file with append mode
func logToFile(message string) {
	// Open the log file in append mode; create it if it doesn't exist
	f, err := os.OpenFile("logs/monitoring.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err) // Log an error if the file can't be opened
		return
	}
	defer f.Close() // Ensure the file is closed after the function completes

	// Write the log message to the file with a newline at the end
	if _, err := f.WriteString(fmt.Sprintf("%s\n", message)); err != nil {
		log.Printf("Error writing to log file: %v", err) // Log an error if writing to the file fails
	}
}

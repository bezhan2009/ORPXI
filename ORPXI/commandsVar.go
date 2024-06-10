package ORPXI

import "goCmd/structs"

var commands = []structs.Command{
	{"help", "List of available commands"},
	{"whois", "Domain information"},
	{"pingview", "Displays ping"},
	{"traceroute", "Route tracing"},
	{"extractzip", "Extracts .zip archives"},
	{"copysource", "from file copy his source"},
	{"signout", "User signs out of ORPXI"},
	{"newshablon", "Creates a new command template for execution"},
	{"shablon", "Executes a specific command template"},
	{"newuser", "New user for ORPXI"},
	{"promptSet", "Changes ORPXI"},
	{"systemgocmd", "Displays information about ORPXI"},
	{"rename", "Renames a file"},
	{"remove", "Deletes a file"},
	{"read", "Displays file contents"},
	{"write", "Writes data to a file"},
	{"create", "Creates a new file"},
	{"exit", "Exits the program"},
	{"orpxi", "Starts another ORPXI"},
	{"wifiutils", "Launches utility for WiFi operations"},
	{"clean", "Clears the screen"},
	{"matrixmul", "Multiplication of large matrices"},
	{"primes", "Search for large prime numbers"},
	{"picalc", "Calculates the value of π"},
	{"fileio", "File I/O intensive test"},
	{"cd", "Changes the current directory"},
	{"edit", "Edits a file"},
	{"ls", "Displays directory contents"},
	{"scanport", "Port scanning"},
	{"dnslookup", "DNS queries"},
	{"ipinfo", "IP address information"},
	{"geoip", "IP address geolocation"},
}

var commandHistory []string

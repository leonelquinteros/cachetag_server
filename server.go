package main

import "github.com/leonelquinteros/cachetag"

func main() {
    // Load configuration
    config := cachetag.GetConfig()

    // Start server
    cachetag.StartServer(config)
}

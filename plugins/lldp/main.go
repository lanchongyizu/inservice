package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/RackHD/inservice/agent"
	"github.com/spf13/viper"
)

var binaryName, buildDate, buildUser, commitHash, goVersion, osArch, releaseVersion string

func main() {
	log.Println(binaryName)
	log.Println("  Release version: " + releaseVersion)
	log.Println("  Built On: " + buildDate)
	log.Println("  Build By: " + buildUser)
	log.Println("  Commit Hash: " + commitHash)
	log.Println("  Go version: " + goVersion)
	log.Println("  OS/Arch: " + osArch)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("InService LLDP Configuration Error: %s\n", err)
	}

	viper.SetConfigName("inservice")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/inservice.d")
	viper.AddConfigPath(dir)
	viper.AddConfigPath("$GOPATH/bin")
	viper.AddConfigPath("$HOME")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("InService LLDP Configuration Error: %s\n", err)
	}

	log.Printf("InService LLDP Configuration: %s\n", viper.ConfigFileUsed())

	lldp, err := NewLLDPPlugin(
		viper.GetString("agent.http.address"),
		viper.GetInt("plugins.inservice-lldp.port"),
		viper.GetStringSlice("plugins.inservice-lldp.interfaces"),
	)
	if err != nil {
		log.Fatalf("Unable to initialize Plugin: %s\n", err)
	}

	p, err := plugins.NewPlugin(lldp)
	if err != nil {
		log.Fatalf("Unable to host Plugin: %s\n", err)
	}

	p.Serve()
}

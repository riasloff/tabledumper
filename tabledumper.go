package tabledumper

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

// Config holds
type Config struct {
	Host     string
	User     string
	Password string
	DbName   string
	Tables   []string
}

func dumpTable(cfg Config, tableName string, wg *sync.WaitGroup) {
	defer wg.Done()

	outputFile, err := os.Create(fmt.Sprintf("%s.sql", tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	fmt.Println("dumping: ", tableName)
	dumpCmd := fmt.Sprintf("mysqldump -h%s -u%s -p%s %s %s > %s.sql", cfg.Host, cfg.User, cfg.Password, cfg.DbName, tableName, tableName)
	cmd := exec.Command("bash", "-c", dumpCmd)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("dump completed: ", tableName)
}

// Start dumping
func Start(cfg Config) {
	var wg sync.WaitGroup
	fmt.Printf("pid: %d\n", os.Getpid())
	for _, table := range cfg.Tables {
		wg.Add(1)
		go dumpTable(cfg, table, &wg)
	}

	wg.Wait()
}

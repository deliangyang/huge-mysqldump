package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

type MysqlConfig struct {
	Host string
	Username string
	Password string
	Database string
}

type Backup struct {
	config MysqlConfig
}

type Recover struct {
	config MysqlConfig
}

var (
	backup Backup
	recover Recover
	tables[] string
	savePath string
)

func CheckParams(config MysqlConfig) {
	if len(config.Host) <= 0 {
		log.Panic("empty host")
	}
	if len(config.Username) <= 0 {
		log.Panic("empty username")
	}
	if len(config.Password) <= 0 {
		log.Panic("empty password")
	}
	if len(config.Database) <= 0 {
		log.Panic("empty database")
	}
}

func (backup Backup)ShowTables() (tables []string, err error) {
	var cmd *exec.Cmd

	cmd = exec.Command("mysql", "-h" + backup.config.Host,
		"-u" + backup.config.Username, "-p" + backup.config.Password, backup.config.Database,
		"-e", "show tables")
	stdout, err := cmd.StdoutPipe()
	if  err != nil {
		log.Fatal(err)
		return tables, err
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return tables, err
	}

	bytes, err := ioutil.ReadAll(stdout)
	items := strings.Split(string(bytes), "\n")
	for i := range items {
		if !strings.Contains(items[i], "| ") ||
			strings.Contains(items[i], "Tables_in_" + backup.config.Database) {
			continue
		}
		table := strings.Replace(items[i], "|", "", 2)
		table = strings.TrimSpace(table)
		tables = append(tables, table)
	}
	log.Println(items)
	return tables, nil
}

func (backup Backup)SaveTable(table string) (err error) {
	var cmd *exec.Cmd
	cmd = exec.Command("mysqldump", "--opt", "-h" + backup.config.Host,
		"-u" + backup.config.Username, "-p" + backup.config.Password, backup.config.Database,
		">", savePath + table + ".sql")

	if _, err := cmd.StdoutPipe(); err != nil {
		log.Fatal(err)
		return err
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("save " + table + " success!!")
	return nil
}

func init()  {
	flag.StringVar(&recover.config.Host, "lHost", "", "local host")
	flag.StringVar(&recover.config.Username, "lUser", "root", "local Username")
	flag.StringVar(&recover.config.Password, "lPassword", "", "local password")
	flag.StringVar(&recover.config.Database, "lDb", "", "local database")

	flag.StringVar(&backup.config.Host, "dHost", "", "dis host")
	flag.StringVar(&backup.config.Username, "dUsername", "root", "dis Username")
	flag.StringVar(&backup.config.Password, "dPassword", "", "dis password")
	flag.StringVar(&backup.config.Database, "dDb", "", "dis database")

	flag.Parse()
}

func main() {
	CheckParams(backup.config)

	var tables []string
	tables, err := backup.ShowTables()
	if  err != nil {
		log.Panic("not find tables")
	}
	println(tables)
	for i := range tables {
		if err := backup.SaveTable(tables[i]); err != nil {
			log.Panic(err)
		}
	}
}

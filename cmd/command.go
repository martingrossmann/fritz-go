package cmd

import (
	"github.com/magiconair/properties"
	"github.com/martingrossmann/fritz-go/common"
	"github.com/martingrossmann/fritz-go/fritz"
	"github.com/martingrossmann/fritz-go/writer"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "counter",
	Short: "Read online counter from Fritz.box",
	Long:  "Read online counter from Fritz.box",
	Run: func(cmd *cobra.Command, args []string) {
		readCounterAndWrite()
	},
}

var cmdConvert = &cobra.Command{
	Use:   "convert",
	Short: "Convert CSV file to InfluxDB import format",
	Run: func(cmd *cobra.Command, args []string) {
		convertFile()
	},
}

func Exec() {
	rootCmd.AddCommand(cmdConvert)
	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func readCounterAndWrite() {
	conf := loadConfig()

	instance := &fritz.FritzBox{
		Host:     conf.FritzHost,
		Insecure: false,
		Passw:    conf.FritzPassword,
	}

	err := instance.PerformLogin()
	checkErr(err, "Cannot login to Fritz.Box")
	counter, err := instance.ReadOnlineCounter()
	checkErr(err, "Cannot handle online counter data from Fritz.Box")

	if conf.CsvActive {
		writer.WriteToCSV(conf, counter)
	}

	if conf.InfluxActive {
		writer.WriteToInnflux(conf, counter)
	}
}

func convertFile() {
	conf := loadConfig()

	writer.ConvertCsvToInfluxData(conf)
}

func loadConfig() common.Config {
	conf := common.Config{}
	p := properties.MustLoadFile("settings.conf", properties.UTF8)
	err := p.Decode(&conf)
	checkErr(err, "Error at parsing settings.conf: ")
	return conf
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

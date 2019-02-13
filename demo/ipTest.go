package classTest

import (
	_ "encoding/json"
	"fmt"
	"net"
	_ "testing"
	"github.com/cc14514/go-geoip2-db"
	"sync"
)

func TestNewGeoipDbByStatik() {
	db, _ := geoip2db.NewGeoipDbByStatik()
	defer db.Close()
	record, _ := db.City(net.ParseIP("115.35.95.90"))
	fmt.Println(record)
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)

}
func main() {
	TestNewGeoipDbByStatik()


	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			fmt.Println("i",i)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		<-done
	}
}

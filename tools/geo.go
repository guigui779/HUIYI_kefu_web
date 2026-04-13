package tools

import (
	"github.com/oschwald/geoip2-golang"
	"net"
)

func GetCity(path, ipAddress string) (string, string) {
	db, err := geoip2.Open(path)
	if err != nil {
		return "", ""
	}
	defer db.Close()
	record, err := db.City(net.ParseIP(ipAddress))
	if err != nil {
		return "", ""
	}
	country := record.Country.Names["zh-CN"]
	if country == "" {
		country = record.Country.Names["en"]
	}
	city := record.City.Names["zh-CN"]
	if city == "" {
		city = record.City.Names["en"]
	}
	return country, city
}

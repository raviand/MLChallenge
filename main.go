package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	model "MLChallange/interfaces"

	"github.com/gin-gonic/gin"
	sdk "github.com/mercadolibre/golang-sdk/sdk"
)

var db = make(map[string]string)
var client *sdk.Client

const clientID = int64(7045571514070392)
const secretID = "ZMCSI6UR7HA2eiuAuitPO2kTE4MZv7lt"

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/show/:itemId/*action", func(c *gin.Context) {
		itemId := c.Params.ByName("itemId")
		actions := strings.Split(c.Param("action"), "/")

		siteChannel := make(chan *model.SiteInfo)
		sellerChannel := make(chan *model.SellerInfo)
		categoryChannel := make(chan *model.CategoryData)

		//Get the client Connection
		client, err := GetClient()
		if err != nil {
			log.Println("error: ", err.Error())
			c.JSON(http.StatusOK, gin.H{"response": "Error while trying to connect with the API", "code": "100"})
		}
		//Gets all basic items data
		items, _ := GetItemData(client, itemId)

		siteSent := false
		sellerSent := false
		categorySent := false
		for index, action := range actions {
			fmt.Printf("\nActions %v, %v", action, index)
			if action == "site" && !siteSent {
				siteSent = true
			}
			if action == "seller" && !sellerSent {
				sellerSent = true
			}
			if action == "category" && !categorySent {
				categorySent = true
			}
		}
		go func(siteSent bool) {
			//Get site data by id
			if siteSent {
				siteRes, err := GetSiteData(client, items.SiteID)
				if err != nil {
					//raw["SiteData"] = nil
					siteChannel <- nil
					return
				}
				siteChannel <- siteRes
			} else {
				siteChannel <- nil
			}
		}(siteSent)

		go func(sellerSent bool) {
			//get Seller Data by Id
			if sellerSent {
				seller, err := GetSellerData(client, items.SellerID)
				if err != nil {
					//raw["SiteData"] = nil
					sellerChannel <- nil
					return
				}
				sellerChannel <- seller
			} else {
				sellerChannel <- nil
			}
		}(sellerSent)

		go func(categorySent bool) {
			//get category Data by Id
			if categorySent {
				category, err := GetCategoryData(client, items.CategoryID)
				if err != nil {
					//raw["SiteData"] = nil
					categoryChannel <- nil
					return
				}
				categoryChannel <- category
			} else {
				categoryChannel <- nil
			}
		}(categorySent)

		for i := 0; i < 3; i++ {
			select {
			case items.SiteData = <-siteChannel:
				fmt.Println("\nsite loaded")
			case items.SellerData = <-sellerChannel:
				fmt.Println("\nseller loaded")
			case items.CategoryData = <-categoryChannel:
				fmt.Println("\ncategory loaded")

			}
		}

		c.JSON(http.StatusOK, gin.H{"response": items})

	})

	return r
}
func GetItemData(client *sdk.Client, itemId string) (*model.ItemInterface, error) {
	var items model.ItemInterface

	response, _ := client.Get(fmt.Sprintf("/items/%s", itemId))
	//var site SiteInfo
	itemInfo, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(itemInfo, &items)
	return &items, nil
}

/*
Gets all category data from the ID
*/
func GetCategoryData(client *sdk.Client, categoryId string) (*model.CategoryData, error) {

	var category model.CategoryData

	categoryRes, err := client.Get(fmt.Sprintf("/categories/%s", categoryId))
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil, err
	}

	categoryInfo, _ := ioutil.ReadAll(categoryRes.Body)
	json.Unmarshal(categoryInfo, &category)
	return &category, nil
}

/*
	Gets all the Seller information from an Id
*/
func GetSellerData(client *sdk.Client, sellerId int64) (*model.SellerInfo, error) {

	var seller model.SellerInfo

	sellerRes, err := client.Get(fmt.Sprintf("/users/%d", sellerId))
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil, err
	}

	sellerInfo, _ := ioutil.ReadAll(sellerRes.Body)
	json.Unmarshal(sellerInfo, &seller)
	return &seller, nil
}

/*
	Gets all the site information from an Id
*/
func GetSiteData(client *sdk.Client, item string) (*model.SiteInfo, error) {
	//Get site data by id
	var site model.SiteInfo
	siteRes, err := client.Get(fmt.Sprintf("/sites/%s", item))
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil, err
	}
	//var raw map[string]interface{}
	siteInfo, _ := ioutil.ReadAll(siteRes.Body)
	json.Unmarshal(siteInfo, &site)
	return &site, nil
}

func GetClient() (*sdk.Client, error) {
	return sdk.Meli(clientID, "", secretID, "http://localhost")
}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

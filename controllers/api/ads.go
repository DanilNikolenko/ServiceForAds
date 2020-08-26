package api

import (
	"ServiceForAds/controllers"
	"ServiceForAds/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/labstack/gommon/log"
)

// AdsController operations for Ads
type AdsController struct {
	controllers.BaseController
}

// URLMapping ...
func (c *AdsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @Description create Ads
// @Param	body		body 	models.Ads	true		"body for Ads content"
// @Success 201 {int} models.Ads
// @Failure 403 body is empty
// @router / [post]
func (c *AdsController) Post() {
	var v models.Ads

	// unmarshal request body
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		log.Error(err)
		c.Response(400, nil, err)
	}

	// add to DB
	if _, err := models.AddAds(&v); err != nil {
		log.Error(err)
		c.Response(400, nil, err)
	}

	c.Response(201, v.Id, nil)
}

// GetOne ...
// @Title Get One
// @Description get Ads by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Ads
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AdsController) GetOne() {
	var query = make(map[string]string)

	// check param
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err)
		c.Response(400, nil, err)
	}

	// Get ad by id
	v, err := models.GetAdsById(id)
	if err != nil {
		log.Error(err)
		c.Response(400, nil, err)
	}

	if v := c.GetString("filds"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				err := errors.New("Error: invalid query key/value pair")
				log.Error(err)
				c.Response(400, nil, err)
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	rez := models.Ads{
		Id:   v.Id,
		Name: v.Name,
		//Description: "",
		Img1: v.Img1,
		//Img2:        "",
		//Img3:        "",
		Price: v.Price,
	}

	if query["description"] != "" {
		rez.Description = v.Description
	}

	if query["imgs"] != "" {
		rez.Img2 = v.Img2
		rez.Img3 = v.Img3
	}

	c.Response(201, rez, nil)
}

// GetAll ...
// @Title Get All
// @Description get Ads
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Ads
// @Failure 403
// @router / [get]
func (c *AdsController) GetAll() {
	var query = make(map[string]string)
	if v := c.GetString("filds"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				err := errors.New("Error: invalid query key/value pair")
				log.Error(err)
				c.Response(400, nil, err)
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	sort := query["sort"]
	crease := query["crease"]

	if sort != "price" && sort != "date" {
		err := errors.New("Error: invalid sort!")
		log.Error(err)
		c.Response(400, nil, err)
	}
	if crease != "increase" && crease != "decrease" {
		err := errors.New("Error: invalid crease!")
		log.Error(err)
		c.Response(400, nil, err)
	}

	limitINT, err := strconv.Atoi(query["limit"])
	if err != nil {
		log.Error(err)
		c.Response(400, nil, err)
	}

	offsetINT, err := strconv.Atoi(query["offset"])
	if err != nil {
		log.Error(err)
		c.Response(400, nil, err)
	}

	ads, err := models.GetAllAds(limitINT, offsetINT, sort, crease)
	if err != nil {
		log.Error(err)
		c.Response(500, nil, err)
	}

	c.Response(201, ads, nil)
}

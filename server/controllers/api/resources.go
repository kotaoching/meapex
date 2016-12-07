package api

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meapex/meapex/server/models"
	"github.com/mozillazg/go-slugify"
)

// GetAllResource ...
func GetAllResource(c *gin.Context) {
	resources, _ := models.GetAllResource()
	c.JSON(200, gin.H{
		"data": resources,
	})
}

// GetResourceById ...
func GetResourceById(c *gin.Context) {
	guid := c.Param("id")
	resource, _ := models.GetResourceById(guid)

	var objectReference map[string](interface{})
	json.Unmarshal([]byte(resource.Reference), &objectReference)

	c.JSON(200, gin.H{
		"data": resource,
	})
}

// CreateResource ...
func CreateResource(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	attribute := c.PostForm("attribute")

	reference := map[string](interface{}){
		"num_attribute": map[string](interface{}){
			"0": []string{"author", "d.16Vtz98QRUunv3Rpqfnl5Ep1wXy1Um7kVSWGOTG2hPyYkGMx"},
		},
		"next": 1,
	}
	jsonReference, _ := json.Marshal(reference)

	slug := slugify.Slugify(title)
	if count, _ := models.GetResourceCountBySlug(slug); count != 0 {
		slug = slug + "-sn-" + strconv.Itoa(count)
	}

	resource := &models.Resource{
		Title:     title,
		Slug:      slug,
		Content:   content,
		Attribute: attribute,
		Reference: string(jsonReference),
	}

	err := resource.Create()
	if err == nil {
		var objectReference map[string](interface{})
		json.Unmarshal([]byte(resource.Reference), &objectReference)

		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"guid":      resource.GUID,
				"title":     resource.Title,
				"slug":      resource.Slug,
				"content":   resource.Content,
				"attribute": resource.Attribute,
				"reference": objectReference,
			},
		})
	} else {
	}
}

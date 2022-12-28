package main

import (
	"encoding/json"
	"image/color"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

// albums slice to seed record album data.
var albums = []album{
	{"1", "Blue Train", "John Coltrane", 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	router := gin.Default()
	router.Static("/attrs", "./resources")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// router.LoadHTMLFiles("index.html", "createHero.html", "heros.html")
	router.LoadHTMLGlob("templates/*")

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.GET("/createHero", getHero)
	router.POST("/createHero", postHero)

	router.GET("/heros", getHeros)

	router.GET("", getPro)

	router.Run("localhost:8081")
}

var heros []Hero
var nameRole = make(map[string]Role) //英雄-名字:角色 映射
var names, roles []string            //名字,职业 的集合
var hcolors []color.RGBA

func getHeros(c *gin.Context) {
	herosj, _ := json.Marshal(heros)
	herosjs := string(herosj)
	hcolorsj, _ := json.Marshal(hcolors)
	hcolorsjs := string(hcolorsj)
	c.HTML(200, "heros.html",
		gin.H{"role": roles, "name": names, "nameRole": nameRole, "heros": herosjs, "colors": hcolorsjs},
		// map[string][]Hero{
		// 	"heros": heros,
		// },
	)
	// c.IndentedJSON(http.StatusOK, heros)

}

func getHero(c *gin.Context) {
	pro, _ := c.GetQuery("pro")

	for _, v := range rolePro {
		if v.Profession == pro {
			c.HTML(200, "createHero.html", gin.H{"color": v.Color, "pro": v.Profession})
		}
	}

}

func postHero(c *gin.Context) {
	name := c.PostForm("name")
	role := c.PostForm("role")
	// name := c.Query("name")
	// role := c.Query("role")
	var hp, mp, damage, defence float32
	var hcolor color.RGBA
	//延时一秒才创建角色====没用客户端还是可以快速发送请求
	time.Sleep(1000 * time.Millisecond)

	//检测重名
	for _, v := range names {
		if v == name {
			c.IndentedJSON(http.StatusNotAcceptable, gin.H{
				"message": "重名"})
			panic("重名")
		}
	}
	//根据职业获得初始属性值
	for _, v := range rolePro {
		if v.Profession == role {
			nameRole[name] = v
			hcolor = v.Color
			hp = v.Hp
			mp = v.Mp
			damage = v.Damage
			defence = v.Defence
		}
	}
	//防止curl 注入不存在的职业
	if hp == 0 {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"message": "无效的职业"})
		panic("无效的职业")
	}

	hero := Hero{name, role, hp, mp, 1, damage, defence}
	hcolors = append(hcolors, hcolor)
	heros = append(heros, hero)
	names = append(names, name)
	roles = append(roles, role)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "收到",
	})
}

func getPro(c *gin.Context) {
	// c.Redirect(http.StatusFound, "/")

	// c.JSON(200, gin.H{
	// 	"message": zhanShi.color,
	// })

	c.HTML(200, "index.html", gin.H{
		"zhanshi":    zhanShi.Color,
		"prozhanshi": zhanShi.Profession,

		"fashi":    faShi.Color,
		"profashi": faShi.Profession,

		"mushi":    muShi.Color,
		"promushi": muShi.Profession,

		"lieren":    lieRen.Color,
		"prolieren": lieRen.Profession,

		"xiaode":    xiaoDe.Color,
		"proxiaode": xiaoDe.Profession,

		"siqi":    siQi.Color,
		"prosiqi": siQi.Profession,

		"shengqi":    shengQi.Color,
		"proshengqi": shengQi.Profession,

		"shushi":    shuShi.Color,
		"proshushi": shuShi.Profession,

		"liemo":    lieMo.Color,
		"proliemo": lieMo.Profession,
	})
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

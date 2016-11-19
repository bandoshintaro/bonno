package controllers

import (
    "io/ioutil"
    "math/rand"
    "time"
    "github.com/revel/revel"
    "bonno/app/models"
    "bonno/app/routes"
)

type App struct {
    *revel.Controller
}

//インデックス画面を表示する
func (c App) Index() revel.Result {
    movielist, _ := DbMap.Select(models.Movie{}, "select * from Movie order by RANDOM() limit 1")
    catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
    return c.Render(movielist,catelist)
}
//かっちょいいトップ画面を表示する
func (c App) Top() revel.Result {
    abdir := revel.Config.StringDefault("bonno.abdir", "./")
    reldir := revel.Config.StringDefault("bonno.reldir","./")
    topmoviedatas, _ := ioutil.ReadDir(abdir+"/other/top")
	rand.Seed(time.Now().UnixNano())
    i := len(topmoviedatas)
    topmoviepath := reldir+"/other/top/"+topmoviedatas[rand.Intn(i)].Name()
    return c.Render(topmoviepath)
}

//動画一覧を表示する
func (c App) Movie() revel.Result {
    movielist, _ := DbMap.Select(models.Movie{}, "select * from Movie order by RANDOM() limit 10")
    catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
    return c.Render(movielist,catelist)
}


func (c App) Search(word string) revel.Result {
    movielist, _ := DbMap.Select(models.Movie{}, "select * from Movie where Name like ?",word)
    catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
    return c.Render(word,movielist,catelist)
}

func (c App) Category(tag string) revel.Result {
    movielist, _ := DbMap.Select(models.Movie{}, "select * from Movie where Tag1=? or Tag2=? or Tag3=? or Tag4=?",tag,tag,tag,tag)
    catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
    return c.Render(tag,movielist,catelist)
}

//動画詳細を表示する
func (c App) Detail(id int) revel.Result {
    rows, _  := DbMap.Select(models.Movie{}, "select * from Movie where Id=?",id)
    movie := rows[0].(*models.Movie)
    abdir := revel.Config.StringDefault("bonno.abdir", "./")
    moviedata, _ := ioutil.ReadDir(abdir+"/movie/"+movie.Tag1+"/"+movie.Name)
    catelist, _ := DbMap.Select(models.Movie{}, "select distinct Tag1 from Movie")
    return c.Render(movie,moviedata,catelist)
}

//動画のディレクトリを読み込んでDBに追加する
func (c App) Init() revel.Result {
    abdir := revel.Config.StringDefault("bonno.abdir", "./")
    reldir := revel.Config.StringDefault("bonno.reldir","./")
    dirlist, _ := ioutil.ReadDir(abdir+"/movie")
    for _, dir := range dirlist{
        titlelist, _ := ioutil.ReadDir(abdir+"/movie/"+dir.Name())
        for _, title := range titlelist{
            rows, _ := DbMap.Select(models.Movie{}, "select * from Movie where Name=?",title.Name())
            if len(rows)==0{
                DbMap.Insert(&models.Movie{0, title.Name(),reldir+"/movie/"+dir.Name()+"/"+title.Name(),reldir+"/images/thumb/"+title.Name()+".png",0,"",title.ModTime(),dir.Name(),"","",""})
            }
        }
    }
    return c.Redirect(routes.App.Index())
}

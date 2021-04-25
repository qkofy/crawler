# Crawler

Crawler是对Gobot的重构。

Crawler是一个轻量级网页爬虫，用于采集数据，数据采集规则使用re2正则进行匹配。

## 一、安装

```
go get github.com/qkofy/crawler
```

## 二、配置

**配置格式**：

以.json文件保存的JSON数据，路劲：./config

**配置参数**：

| 参数名 |  类型  | 必填 |                             说明                             |
| :----: | :----: | :--: | :----------------------------------------------------------: |
|  filt  | string |  否  | 目标网址过滤，过滤网址以string格式拼接，直接设为参数值，或保存于./config/配置文件名.txt，远程文件参数值设为文件网址，本地文件参数值设为local |
|  hurl  | string |  否  |          目标首页网址，目标网址需要补全时设置该参数          |
|  list  | array  |  是  | 列表页采集规则，规则为键值对形式出现（值为re2正则），1- 固定键名（目标规则）： target，2 - 可选键名（目标区域）：extent |
|  mode  | number |  否  |  采集模式：1 - 采集列表，2 - 采集详情，默认 - 1、2顺序采集   |
|  murl  | string |  是  |       （列表）有序分页网址组合规则（值为sprintf规则）        |
|  page  | array  |  是  | 详情页采集规则，规则为键值对形式出现（值为re2正则），键名自定义 |
|  surl  | string |  是  |        （列表）频道页网址组合规则（值为sprintf规则）         |
|  tout  | number |  否  |                       超时时间，默认3s                       |

**配置示例**：

```
[
    'filt' => 'local',
    'hurl' => '',
    'list' => [
    	'target' => '<a href="(http://www.abc.com/[a-z]+/\d+.html)">',
    	'extent' => '(?s)<div class="part-box cl">(.*)<div class="page-num">',
    ],
    'mode' => '',
    'murl' => 'http://www.abc.com/%s/page/%d.html',
    'page' => [
        'description' => '<meta name="description" content="([^"]+)">',
		'imgurls' => '(?sU)<div class="scoll-soft-pic">\s*<div class="box">(.*)</div>',
		'introduce' => '(?s)<div class="article-con  padlf yyb_app">(.+)\s</div>\s</div>\s<div class="part-box">',
		'keywords' => '<meta name="keywords" content="([^"]+)">',
		'litpic' => '<div class="link-img"> <img src="([^"]+)" alt="[^"]*"></div>',
		'seotitle' => '<title>([^|]+) [^<]*</title>',
		'softlinks' => '<a class="and" href=\'([^\']*)\' target=\'_blank\'>[^<]*</a>',
		'title' => '<h1 class="link-name">([^<]+)</h1>',
        'typename' => '<p>类别：([^<]*)</p>',
		'versionnumber' => '<title>[^ ]* ([^ ]*) [^<]*</title>',
    ],
    'surl' => 'http://www.abc.com/%s',
    'tout' => '6'
]
```

上述示例为数组格式需转换为json格式，然后保存于./config/文件名.json，或以字符串形式传入

```
{
    "filt": "local", 
    "hurl": "", 
    "list": {
        "target": "<a href=\"(http://www.abc.com/[a-z]+/\\d+.html)\">", 
        "extent": "(?s)<div class=\"part-box cl\">(.*)<div class=\"page-num\">"
    }, 
    "mode": "", 
    "murl": "http://www.abc.com/%s/page/%d.html", 
    "page": {
        "description": "<meta name=\"description\" content=\"([^\"]+)\">", 
        "imgurls": "(?sU)<div class=\"scoll-soft-pic\">\\s*<div class=\"box\">(.*)</div>", 
        "introduce": "(?s)<div class=\"article-con  padlf yyb_app\">(.+)\\s</div>\\s</div>\\s<div class=\"part-box\">", 
        "keywords": "<meta name=\"keywords\" content=\"([^\"]+)\">", 
        "litpic": "<div class=\"link-img\"> <img src=\"([^\"]+)\" alt=\"[^\"]*\"></div>", 
        "seotitle": "<title>([^|]+) [^<]*</title>", 
        "softlinks": "<a class=\"and\" href='([^']*)' target='_blank'>[^<]*</a>", 
        "title": "<h1 class=\"link-name\">([^<]+)</h1>", 
        "typename": "<p>类别：([^<]*)</p>", 
        "versionnumber": "<title>[^ ]* ([^ ]*) [^<]*</title>"
    }, 
    "surl": "http://www.abc.com/%s", 
    "tout": "6"
}
```

## 三、使用

**1、参数说明**

| 参数名 |  类型  | 必填 | 默认值 |                      说明                      |
| :----: | :----: | :--: | :----: | :--------------------------------------------: |
|  arg1  | string |  是  |        | 配置文件名（不含后缀），或配置字符串，格式json |
|  arg2  | string |  是  |        |                采集目标列表路劲                |
|  arg3  |  int   |  是  |   0    |                 采集目列表页码                 |
|  arg4  |  int   |  是  |   0    |                采集目标列表页数                |

**2、调用内置函数**

```
crawler.Run(arg1, arg2, arg3, arg4)  //返回json字符串
```

示例

```
//配置文件./config/abc.json
crawler.Run("abc", "list", 0, 0)   //如：http://www.abc.com/list
crawler.Run("abc", "list", 2, 3)   //如：http://www.abc.com/list/page/2.html
```

或

```
//配置变量 cfg := `{...}`
crawler.Run(cfg, "list", 0, 0)
crawler.Run(cfg, "list", 2, 3)
```

**3、使用内置方法**

```
crawler.NewConfig(arg1)
url := fmt.Sprintf(cfg.Surl, arg2)
L := crawler.NewList()
L.ParseList(url, cfg)
P := crawler.NewPage()
P.ParsePage(list.Urls[0], cfg)
crawler.GetResult(P).ToJson()
```

或

```
crawler.NewConfig(arg1)
for i := arg3; i < arg4; i++
url := fmt.Sprintf(cfg.Murl, arg2, i)
L := crawler.NewList()
L.ParseList(url, cfg)
P := crawler.NewPage()
P.BatchParsePage(list.Urls, cfg)
crawler.GetResult(P).ToJson()
```

以上示例为伪代码，完整程序需自行完善
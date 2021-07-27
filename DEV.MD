# HTTP SOCKET
## POST /api/new
Request:
```HTTP
POST /api/new HTTP/1.1
Host: ks.sorakado.com

```

Response:
```json
{
    "id":"60ffba09d9bb72512b5094e0",
    "base":"",
    "pieces":[
        {
            "path":"m4tubw/4.mp3",
            "full":"m4tubw/full.mp3",
            "name":"Bravely You",
            "title":"Bravely You (Instrumental)",
            "choices":[
                "Saya's Song",
                "羽のゆりかご",
                "風の辿り着く場所"
            ]
        },
        {
            "path":"nv3y92/2.mp3",
            "full":"nv3y92/full.mp3",
            "name":"日々の遑",
            "title":"日々の遑",
            "choices":[
                "伝承",
                "木陰の憩",
                "kanade"
            ]
        },
        {
            "path":"7vfy9o/9.mp3",
            "full":"7vfy9o/full.mp3",
            "name":"かの記憶、あの記憶",
            "title":"かの記憶、あの記憶",
            "choices":[
                "灰燼に帰す",
                "水たまり",
                "Mission possible"
            ]
        },
        {
            "path":"22x0kw/6.mp3",
            "full":"22x0kw/full.mp3",
            "name":"存在",
            "title":"存在 -piano-",
            "choices":[
                "資料室のお茶会",
                "Little Busters!",
                "絵空事"
            ]
        },
        {
            "path":"zrlygs/2.mp3",
            "full":"zrlygs/full.mp3",
            "name":"白詰草",
            "title":"白詰草",
            "choices":[
                "降り続く雨の街で",
                "theme of SSS",
                "Deep Blue Blue"
            ]
        },
        {
            "path":"jem5a7/7.mp3",
            "full":"jem5a7/full.mp3",
            "name":"少年たちとの別れ",
            "title":"少年たちとの別れ",
            "choices":[
                "坂の上の陽炎",
                "Twinkle of Alcor",
                "影は色濃くて"
            ]
        },
        {
            "path":"y9jlzr/3.mp3",
            "full":"y9jlzr/full.mp3",
            "name":"Other side Blue",
            "title":"Other side Blue",
            "choices":[
                "魔法のアンサンブル",
                "影は色濃くて",
                "Other side Blue"
            ]
        },
        {
            "path":"illuti/14.mp3",
            "full":"illuti/full.mp3",
            "name":"Song for friends",
            "title":"Song for friends",
            "choices":[
                "Wing of Glass",
                "凍土高原",
                "猫と硝子と円い月"
            ]
        },
        {
            "path":"piqzx8/0.mp3",
            "full":"piqzx8/full.mp3",
            "name":"双星",
            "title":"双星",
            "choices":[
                "かの記憶、あの記憶",
                "夏影",
                "ぼくのいやなこと"
            ]
        },
        {
            "path":"9vbdiz/10.mp3",
            "full":"9vbdiz/full.mp3",
            "name":"幻想",
            "title":"幻想",
            "choices":[
                "冬の花火",
                "pure snows",
                "願いが叶う場所"
            ]
        }
    ]
}
```

## POST /api/get   
Request:
```HTTP
POST /api/get HTTP/1.1
Host: ks.sorakado.com
Content-Type: application/x-www-form-urlencoded

id=60ffba09d9bb72512b5094e0
```

Response:
```json
{
    "id":"60ffba09d9bb72512b5094e0",
    "base":"",
    "pieces":[
        {
            "path":"m4tubw/4.mp3",
            "full":"m4tubw/full.mp3",
            "name":"Bravely You",
            "title":"Bravely You (Instrumental)",
            "choices":[
                "Saya's Song",
                "羽のゆりかご",
                "風の辿り着く場所"
            ]
        },
        {
            "path":"nv3y92/2.mp3",
            "full":"nv3y92/full.mp3",
            "name":"日々の遑",
            "title":"日々の遑",
            "choices":[
                "伝承",
                "木陰の憩",
                "kanade"
            ]
        },
        {
            "path":"7vfy9o/9.mp3",
            "full":"7vfy9o/full.mp3",
            "name":"かの記憶、あの記憶",
            "title":"かの記憶、あの記憶",
            "choices":[
                "灰燼に帰す",
                "水たまり",
                "Mission possible"
            ]
        },
        {
            "path":"22x0kw/6.mp3",
            "full":"22x0kw/full.mp3",
            "name":"存在",
            "title":"存在 -piano-",
            "choices":[
                "資料室のお茶会",
                "Little Busters!",
                "絵空事"
            ]
        },
        {
            "path":"zrlygs/2.mp3",
            "full":"zrlygs/full.mp3",
            "name":"白詰草",
            "title":"白詰草",
            "choices":[
                "降り続く雨の街で",
                "theme of SSS",
                "Deep Blue Blue"
            ]
        },
        {
            "path":"jem5a7/7.mp3",
            "full":"jem5a7/full.mp3",
            "name":"少年たちとの別れ",
            "title":"少年たちとの別れ",
            "choices":[
                "坂の上の陽炎",
                "Twinkle of Alcor",
                "影は色濃くて"
            ]
        },
        {
            "path":"y9jlzr/3.mp3",
            "full":"y9jlzr/full.mp3",
            "name":"Other side Blue",
            "title":"Other side Blue",
            "choices":[
                "魔法のアンサンブル",
                "影は色濃くて",
                "Other side Blue"
            ]
        },
        {
            "path":"illuti/14.mp3",
            "full":"illuti/full.mp3",
            "name":"Song for friends",
            "title":"Song for friends",
            "choices":[
                "Wing of Glass",
                "凍土高原",
                "猫と硝子と円い月"
            ]
        },
        {
            "path":"piqzx8/0.mp3",
            "full":"piqzx8/full.mp3",
            "name":"双星",
            "title":"双星",
            "choices":[
                "かの記憶、あの記憶",
                "夏影",
                "ぼくのいやなこと"
            ]
        },
        {
            "path":"9vbdiz/10.mp3",
            "full":"9vbdiz/full.mp3",
            "name":"幻想",
            "title":"幻想",
            "choices":[
                "冬の花火",
                "pure snows",
                "願いが叶う場所"
            ]
        }
    ]
}
```
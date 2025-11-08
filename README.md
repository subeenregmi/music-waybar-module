# music-waybar-module
small module to show what is being song is being played on music players

supported: spotify

### showcase

no spotify open
![](docs/notopen.png)

spotify open (paused)
![](docs/paused.png)

spotify open (playing)
![](docs/playing.png)

### setup

dependencies
- go
- gnumake
- unix system

waybar config

```jsonc
"custom/music" : {
    "exec": "../../music-waybar-module/build/music-waybar-module",
    "exec-if": "pgrep spotify",
    "format": "{text}",
    "return-type": "json",
    "max-length": 25,
},
```

styles.css

```css
#custom-music {
    background-color: @mantle;
    border: 2px solid @lavender;
    border-radius: 12.5px;
    margin-left: 5px;
    padding: 0px 7.5px;
}

#custom-music.unknown {
    padding-right: 12.5px;
}

#custom-music.normal {
    color: @green;
}

#custom-music.paused {
    color: @yellow;
}

```

to build

```bash
make 
```


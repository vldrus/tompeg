Simple video to h264 mp4 converter (using `ffmpeg` to convert)

```
ffmpeg -i INPUT -vf scale="-2:'min(480,ih)':flags=lanczos",setsar="1:1" -map_metadata -1 -preset slower -pix_fmt yuv420p -b:v 600k -ar 44100 -b:a 66k OUT
```

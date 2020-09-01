# mkimg

```
Usage of this command:
  -c string
    	Set the color of the image. (default "red")
  -fc string
    	Set the color of the font. (default "white")
  -fs int
    	Set the font size. (default 24)
  -h int
    	Set the height of the image. (default 200)
  -n string
    	Set the file name of the image. (default "sampleImage.png")
  -t string
    	Set the text to be added to the image.
  -w int
    	Set the width of the image. (default 200)
```

## Demo

```
$ cd ./cmd/mkimg ; go run main.go -t 私はGoを勉強中です -fs 20
```

![output image](https://github.com/kshiva1126/goybox/blob/master/toys/mkimg/cmd/mkimg/sampleImage.png)

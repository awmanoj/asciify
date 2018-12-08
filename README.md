# asciify
ASCIIfy the pictures 

# Requirements 

```
$ go version 
go version go1.9.2 linux/amd64
```

# Setup 

```
$ git clone https://github.com/awmanoj/asciify.git
... 
$ cd asciify 
... 
$ go build 
$ ./asciify 
... 
```
You can open your browser and send a request to the server. For example: 

```
http://localhost:9000/?image_url=https://upload.wikimedia.org/wikipedia/commons/thumb/5/5f/Jnehru.jpg/220px-Jnehru.jpg 
``` 
# Sample 

Following image generated from [this image](https://upload.wikimedia.org/wikipedia/commons/thumb/5/5f/Jnehru.jpg/220px-Jnehru.jpg) of Jawahar Lal Nehru, the first prime minister of India. 

![alt text](https://raw.githubusercontent.com/awmanoj/asciify/master/samples/jnh.png)

If image is not clear (if there is lot of background for example) then pass `?dscale=1` or `dscale=2` parameter. Default value is `dscale=3`. To increase graininess increase this value.

# Parameters 

Following are the parameters: 

| Parameter Name | Description | Mandatory | Default |
| ---            | ---         | ---       | ----    | 
| image_url      | URL of the image to be asciified | Yes     | No | 
| rhint          | Reduce the size of the image by this fraction | No   | 4  | 
| maxw           | Maximum width of the output image (aspect ratio always maitained)  | No  | 200  | 





## HTTP Package

- HTTP request to google.com
- Print response to terminal

<img src="./http-pictures/001.png" alt="*Response" style="width:50%">

## GET Response

- `func Get(url string) (resp *Response, err error)`
- `resp, err := http.Get("http://example.com/")`

## Reader Interface

- A **common point of contact** for the multiple sources of input coming into the application.
- So many functions but they all just mean printing the source of input with differnt data types.

<img src="./http-pictures/002.png" alt="*Response" style="width:49%;display:inline-block">
<img src="./http-pictures/003.png" alt="*Response" style="width:50%;display:inline-block">

<img src="./http-pictures/004.png" alt="*Response" style="width:49%;display:inline-block">
<img src="./http-pictures/005.png" alt="*Response" style="width:50%;display:inline-block">

## Read Function

<img src="./http-pictures/006.png" alt="*Response" style="width:50%;display:inline-block">

```
func main() {
	resp, err := http.Get("https://www.google.com")

	// error handling
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// takes a type of slice, and number of elements that the slice can contain
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
```

## Writer Interface

<img src="./http-pictures/007.png" alt="*Response" style="width:49%;display:inline-block">
<img src="./http-pictures/008.png" alt="*Response" style="width:50%;display:inline-block">

```
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

## Copy Function

`func Copy(dst Writer, src Reader) (written int64, err error)`

- `dst Writer`: taking data and send outside the application
- `src Reader`: implement the Reader interface

- Simpler one-line method of logging out the HTML response.
  `io.Copy(os.Stdout, resp.Body)`

<img src="./http-pictures/009.png" alt="*Response" style="width:50%;display:inline-block">

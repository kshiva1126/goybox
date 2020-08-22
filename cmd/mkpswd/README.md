# mkpswd

```
Usage of this command:
  -char value
    	Multiple selections are available (default: n)
    	
    	l: Lowercase characters
    	u: Uppercase characters
    	n: Including number
    	s: Including symbolic characters
    	c: Not including confusing characters, like l o I O 0 1 " ' , . : ; ^ _ ` | ~
  -nchar int
    	The number of characters (default 8)
  -npass int
    	The number of passwords (default 1)
```

## Demo

```
$ go run main.go character_flags.go -char l -char u -char n -char s -char c -nchar 16 -npass 2
[NGpjFv%RT{Pu}r=f kW/j*cDzsFYAN>Ey]
```
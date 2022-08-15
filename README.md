# random string
Generate a random string

# Usage

Return a string of 32 characters.

```bash
./randomstring
zEbNgSn1oOed7HsuzL2cl4TjPXv20mI3
```


Return a string using the input number

```bash
./randomstring 22
OWrslUCVNBFlylp4StQOxa

./randomstring 50
3zlmiyPmueMIi306FY2sEAHAfJVgFPOsVq69V6FSLosLsDQAYk
```

## Help

If you use the argument **-h** or **help** return this:

```bash
./randomstring -h
randomstring [Argument]
                  
randomstring:
  Without arguments return a string of 32 characters between uppercases, lowercases and numbers.
                  
randomstring version | randomstring -v:
  Return the version of this example.

randomstring help | randomstring -h:
  Return this help.
```

## Version

If you use the argument **-v** or **version** return this:

```bash
./randomstring -v
0.1.1
```

# Usage example

```bash
mkdir $(randomstring)
cp /home/user/file.xyz /home/user/$(randomstring 45)
mv /home/user/file.xyz $(randomstring)
```

# License

[GPLv3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)

# random string
A golang library for generate a random string lowercases, string uppercase and numbers. Util for unix-like shell scripting.

# Usage

Return a string of 32 characters.

```bash
./randstr
zEbNgSn1oOed7HsuzL2cl4TjPXv20mI3
```


Return a string using the input number

```bash
./randstr 22
OWrslUCVNBFlylp4StQOxa

./randstr 50
3zlmiyPmueMIi306FY2sEAHAfJVgFPOsVq69V6FSLosLsDQAYk
```

## Help

If you use the argument **-h** or **help** return this:

```
randstr [Argument]
                 
    -h or help    : Return this help.
    -v or version : Return the version. 
    -n            : Return numbers.
    -l            : Return lowercase letters.
    -u            : Return uppercase letters.
    -nl or -ln    : Return lowercase letters with numbers.
    -un or -nu    : Return uppercase letters with numbers.
    -ul or -lu    : Return uppercase letters with lowercase letters.
    default       : Return a string of 32 characters between uppercases, lowercases and numbers.
    
    Note: All previous parameters can use a number for the output length.
        
    Some examples:
        
    $ randstr
    cDUKhgpkuIgMdqs3JibEjJoyBiBjgjZi
        
    $ randstr 35
    ghrzAtgElMIMxCRFMFCbUCnwiDuEkejKocQ
        
    $ randstr -n
    54770780734477138272862081928945
     
    $ randstr -n 33
    280309865114201905814323059174146
        
    $ randstr -nl 35
    hs9wsuhr6m2dyez2zdfyxduglsccygtyzaq
```

## **Warning**                                                                                                                             
        On most GNU/Linux file systems, filenames can be up to 255 characters long, 
		this is the most portable character set [-._A-Za-z0-9] 

## Version

If you use the argument **-v** or **version** return this:

```bash
./randstr -v
0.1.2
```

# Usage example

```bash
mkdir $(randstr -l 27)
cp /home/user/file.xyz /home/user/$(randstr -u 45)
mv /home/user/file.xyz $(randstr -l)
```

# License

[GPLv3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)

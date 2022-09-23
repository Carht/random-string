# random string
A golang library for generate a random string lowercases, string uppercase and numbers. Util for unix-like shell scripting.

# Usage

Return a string of 32 characters.

```bash
./randstr
GAPkKnVFRICHskSCDxeTFTknIIsQBbdd
```


Cases:

* Default (length of this string: 32 letters, upper and lowercases):

```bash
./randstr
UDckrRnCvURNRuKvNWziyRIBxREtjaEQ
```

* Another cases:

```bash
./randstr -lu 35
WIMVngPABZGiBHcCuhmDLTVcoBAHrdJJUGH

./randstr -l 35
sgjxwcoqnkkqogzotoigozvyjdmzsegixul

./randstr -u 35
IBZDSTKGTFQSUWIICQSOTUBBNYISCIAWOYC

./randstr -nl 30
en8usol16pe8e6fhwy5c8xz7d5v9sr

./randstr -nu 32
DYCA7VFE1AURUI66PZGISG7IHQ8707HX

./randstr -n 20
46495882363454181740

./randstr -v
0.1.3
```

## **Warning**                                                                                                                             
        On most GNU/Linux file systems, filenames can be up to 255 characters long, 
		this is the most portable character set [-._A-Za-z0-9] 

## Version

If you use the argument **-v** or **version** return this:

```bash
./randstr -v
0.1.3
```

# Usage example

```bash
mkdir $(randstr -l 27)
cp /home/user/file.xyz /home/user/$(randstr -u 45)
mv /home/user/file.xyz $(randstr -l 17)
```

# License

[GPLv3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)

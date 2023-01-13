According to your operating system it will create an executable. Like windows have .exe and same for other Operarting Systems.
1. you can check `GOOS` by th following command -

```
go env
```

2. For linux you can use following command -

```
GOOS="linux" go build main.go
```

The above command will an executable and you can check it's output by following command- 

```
>> ./main

This is Time !!
2022-09-03 19:17:11.0181496 +0530 IST m=+0.000230701
09-03-2022 19:17:11 Saturday
Today's date is 09-03-2022
Current Time is 19:17:11
Today's day is Saturday
1999-08-12 23:23:05 +0000 UTC
08-12-1999 Thursday
year is  1999
month is  August
date is  12
```

3. For Windows you can use following command -

```
GOOS="windows" go build main.go
```

The above command will an executable and you can check it's output by following command-

```
>> main.exe

This is Time !!
2022-09-03 19:09:19.8800764 +0530 IST m=+0.005353301
09-03-2022 19:09:19 Saturday
Today's date is 09-03-2022
Current Time is 19:09:19
Today's day is Saturday
1999-08-12 23:23:05 +0000 UTC
08-12-1999 Thursday
year is  1999
month is  August
date is  12
```

4. For MacOS you can use following command -

```
GOOS="darwin" go build main.go
```


# goAnyDate.go
D:\hai>go build goAnyDate.go

D:\hai>goAnyDate.exe a

Err:验证字符串是否为数字(包含正负数)

D:\hai>goAnyDate.exe -1

20180927

D:\hai>goAnyDate.exe 1

20180929

D:\hai>goAnyDate.exe %

Err:验证字符串是否为数字(包含正负数)

D:\hai>goAnyDate.exe 0.1

Err:验证字符串是否为数字(包含正负数)

D:\hai>goAnyDate.exe -0.1

Err:验证字符串是否为数字(包含正负数)

D:\hai>goAnyDate.exe -100

20180620

D:\hai>goAnyDate.exe 100

20190106

# time
time


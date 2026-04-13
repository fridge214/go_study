---
title: "Go语言复数 | Go语言教程"
source: "https://go.mosong.cc/tutorial/view_4033.html"
author:
published:
created: 2026-04-13
description: "Go教程 Go入门到精通 Go中文教程 Go教程推荐"
tags:
  - "clippings"
---
## Go语言复数

在计算机中，复数是由两个浮点数表示的，其中一个表示实部（real），一个表示虚部（imag）。

Go语言中复数的类型有两种，分别是 complex128（64 位实数和虚数）和 complex64（32 位实数和虚数），其中complex128 为复数的默认类型。

复数的值由三部分组成 RE + IMi，其中 RE 是实数部分，IM 是虚数部分，RE 和 IM 均为 float 类型，而最后的 i 是虚数单位。

声明复数的语法格式如下所示：

var name complex128 = complex(x, y)

其中 name 为复数的变量名，complex128 为复数的类型，“=”后面的 complex 为Go语言的内置函数用于为复数赋值，x、y 分别表示构成该复数的两个 float64 类型的数值，x 为实部，y 为虚部。

上面的声明语句也可以简写为下面的形式：name:= complex(x, y)

对于一个复数 `z := complex(x, y)` ，可以通过Go语言的内置函数 `real(z)` 来获得该复数的实部，也就是 x；通过 `imag(z)` 获得该复数的虚部，也就是 y。

【示例】使用内置的 complex 函数构建复数，并使用 real 和 imag 函数返回复数的实部和虚部：

```
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"
```

复数也可以用 `==` 和 `!=` 进行相等比较，只有两个复数的实部和虚部都相等的时候它们才是相等的。

Go语言内置的 math/cmplx 包中提供了很多操作复数的公共方法，实际操作中建议大家使用复数默认的 complex128 类型，因为这些内置的包中都使用 complex128 类型作为参数。
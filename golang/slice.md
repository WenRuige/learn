##### slice





> golang slice 和 array的区别

数组传的是值,不会改变元素的值

`slice`也是值语义,但本身是指针类型的,所以会改变值

`array`简单数组:是一组同类型元素的序列,序列长度也是数组定义的一部分(长度不可变)






>  slice 和 unsafe.Pointer相互转换

可以使用`reflect.SliceHeader`方式来构造`slice`

```go
    arr := [10]int{1,2,3,4,5,6,7,8,9,0}
    size:=len(arr)
    p:=uintptr(unsafe.Pointer(&arr))
    
    var data []byte
    sh:=(*reflect.SliceHeader)(unsafe.Pointer(&data))
    sh.Data = p
    sh.Len  = size
    sh.Cap  = size
```



```go
type slice  struct{
	array unsafe.Pointer  //指向数组的指针
	len int               //当前切片的长度
	cap int               //当前切片的容量
}
```

[golang slice](https://halfrost.com/go_slice/)
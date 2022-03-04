#插件说明
-go自带的语法糖比较少...连max都需要自己写..更不用说判断元素内是否包含某元素了
http://wxnacy.com/2018/11/20/go-in-array/ ←这个插件内置了几个方法，暂时满足需求了

##安装
> go get -u  github.com/wxnacy/wgo
##导入
> import "github.com/wxnacy/wgo/arrays"

##常用 Contains 方法
>func Contains(array interface{}, val interface{}) (index int)
>func ContainsString(array []string, val string) (index int)
>func ContainsInt(array []int64, val int64) (index int)
> func ContainsFloat(array []float64, val float64) (index int)
> func ContainsBool(array []bool, val bool) (index int)
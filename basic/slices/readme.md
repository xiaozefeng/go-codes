## slice
## slice 是对底层数组的视图
## reSlice 指的是对slice 再进行slice, reSlice出来的 slice 也是指向底层数组的
## slice 可以向后扩展，不可以向前扩展
## s[i]不可以超过len(s), 向后扩展不可以超越底层数组cap(s)
## 使用append添加元素时如果超过了 slice 的 cap, 系统会分配更大的底层数组
## 由于是值传递的原因，必须接受append的返回值 s = append(s, val)

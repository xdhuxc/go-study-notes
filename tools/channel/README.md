
### select 语句

1、每次执行 select 时，都会只执行其中 1 个 case 或者执行 default 语句

2、当没有 case 或者 default 可以执行时，select 则阻塞，等待直到有 1 个 case 可以执行

3、当有多个 case 可以执行时，则随机选择 1 个 case 执行

4、case 后面跟的必须是读或写通道的操作，否则编译出错

5、当在 case 上读一个通道时，如果这个通道是 nil，则该通道永远阻塞

6、在 select 内的 break 并不能跳出 for-select 循环，如果想结束该循环，可以考虑如下方式：

1）在满足条件的 case 内，使用 return 结束携程，如果还有收尾工作，尝试交给 defer 来处理

2）在 select 外 for 内使用 break 跳出循环

3）使用 goto 语句跳出循环

7、select{} 永远阻塞
go字符串拼接总结 ：
1. + 连接适用于短小的、常量字符串（明确的，非变量），因为编译器会给我们优化。
2. Join是比较统一的拼接，不太灵活
3. fmt和buffer基本上不推荐
4. builder从性能和灵活性上，都是上佳的选择
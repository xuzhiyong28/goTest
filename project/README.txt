目录结构最佳实践
    - /pkg          这个目录中存放的就是项目中可以被外部应用使用的代码库
    - /internal     存放私有代码，该文件夹下的所有包及相应文件都有一个项目保护级别，即其他项目是不能导入这些包的，仅仅是该项目内部使用
    - /conf         存放配置文件
    - /cmd          启动路径，后面编译成可执行路径
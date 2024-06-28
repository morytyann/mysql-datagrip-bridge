## 项目目的

某些运维软件不支持DataGrip，遂用此项目做桥，接收MySQL数据库连接参数，转化为DataGrip项目并自动打开

## 使用方法

1. 自行编译`mysql-datagrip-bridge.exe`
2. 将`mysql-datagrip-bridge.exe`移动到`datagrip64.exe`同级目录下
3. 在运维软件中配置工具，将`MySQL 命令行客户端`的路径改为`mysql-datagrip-bridge.exe`的路径
4. 使用`MySQL 命令行客户端`连接数据库服务
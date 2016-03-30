# akbs

基于gin的http-api服务框架 


#### Start using it
1. 下载:
    ```sh
    $ go get github.com/phillihq/akbs
    ```

2. 进行项目构建:
    ```sh
    $ make build
    ```

3. 生成输出文件
    ```sh
    上面步骤生成builds文件夹,进入此文件夹会发现生成了一个二进制文件akb和一个包含配置文件的文件夹config
    ```

4. 执行二进制文件
    ```go
    ./akb -config config/conf.yaml
    ```
    
5. 浏览器访问
    ```sh
    在地址栏输入 http://localhost:8080
    
    若显示 "Power by akbs" 则程序运行正常
    ```




# grpc-demo


1. 
   
2. 生成私钥文件:

     ```go
     -- 需要输入密码
     openssl genrsa -des3 -out ca.key 2048
     ```

3. 如何查看openssl.cnf文件具体的路径:

     ```go
      打开命令窗口后输入:openssl version-d
     ```

4. 创建证书请求:

   ```go
   openssl req -new -key ca.key -out ca.csr
   ```

5. 生成ca.crt:

   ```go
   openssl x509 -req -days 365 -in ca.csr -signkey ca.key -out ca.crt
   ```

6. 生成证书私钥server.key:

   ```go
   openssl genpkey -algorithm RSA -out server.key
   ```

7. 通过私钥server.key生成证书请求文件server.csr:

8. ```go
   openssl req -new -nodes -key server.key -out server.csr -days 3650 -config ./openssl.cnf -extensions v3_req
   ```

9. 生成SAN证书:

   ```go
   openssl x509 -req -days 365 -in server.csr -out server.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
   
   
   
    -- key：** 服务器上的私钥文件，用于对发送给客户端数据的加密，以及对从客户端接收到数据的解密。
   ```




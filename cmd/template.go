package cmd

var MysqlDeployTemplate = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kube-db-mysql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kube-db-mysql
  template:
    metadata:
      labels:
        app: kube-db-mysql
    spec:
      containers:
        - name: kube-db-mysql
          image: "mysql/mysql-server:latest"
          ports:
            - containerPort: 3306
          volumeMounts:
            - mountPath: /docker-entrypoint-initdb.d
              name: kube-mysql-data
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: "123456"
            - name: MYSQL_DATABASE    # 新建的数据库
              value: "fydb"
      volumes:
        - name: kube-mysql-data
          hostPath:
            path: /Users/jason/work/go/gin-fanyi-project/.mysql-entrypoint-initdb.d
            type: Directory
`


var RedisDeployTemplate = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kube-db-redis
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kube-db-redis
  template:
    metadata:
      labels:
        app: kube-db-redis
    spec:
      containers:
        - name: kube-db-redis
          image: "redis:latest"
          ports:
            - containerPort: 6379
      volumes:
        - name: kube-redis-data
          hostPath:
            path: /Users/jason/work/go/gin-fanyi-project/redis_data
            type: Directory
`

var GoDeployTemplate = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kube-go
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kube-go
  template:
    metadata:
      labels:
        app: kube-go
    spec:
      containers:
        - name: kube-go
          image: "fanyi:v1"
          ports:
            - name: http
              containerPort: 8080
`

var MysqlSvc = `
apiVersion: v1
kind: Service
metadata:
name: kube-go
spec:
selector:
app: kube-go
ports:
- protocol: TCP
port: 8080
targetPort: http
type: NodePort
`

var RedisSvc = `
apiVersion: v1
kind: Service
metadata:
  name: kube-db-redis
spec:
  selector:
    app: kube-db-redis
  ports:
    - port: 6379
`

var GoSvc = `apiVersion: v1
kind: Service
metadata:
  name: kube-go
spec:
 selector:
   app: kube-go
 ports:
 - protocol: TCP
   port: 8080
   targetPort: http
 type: NodePort
`

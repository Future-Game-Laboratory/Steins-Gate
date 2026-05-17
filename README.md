# 这里是《创世纪》项目组

## Docker管理环境，后端使用GO+Redis+MySQL


## MySQL数据库连接
默认 docker-compose 会创建：

- 地址：localhost:3306
- 数据库：game_db
- 用户：game_app
- 密码：123456


## Docker常用指令
### 后台启动（-d 表示守护进程）
docker-compose up -d

### 查看服务运行状态
docker-compose ps

### 停止服务（不删除数据）
docker-compose stop

### 停止并删除容器（数据卷保留）
docker-compose down

### 查看日志
docker-compose logs -f  # 实时日志
docker-compose logs redis  # 只看 Redis 日志
docker-compose logs mysql  # 只看 MySQL 日志


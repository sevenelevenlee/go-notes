version: '2'
services:
  app:
    image: update.webvpn.net.cn/s002-testing
    restart: always
    environment:
    - DB_USER=s002
    - DB_NAME=s002
    - DB_HOST=172.23.0.16:3306
    - DB_PASSWORD=wvZJFExMtWQpvLkqaoakbQT88
    - MONGO_HOST=root:wvZJFEsMgWQpvLkqaoakbQT88@172.23.0.17
    - ADMIN_HOST=59.65.196.32:8080
    - PPROF=true
    - LOG_LEVEL=debug
    ports:
    - 80:80
    - 8080:8080
    - 443:443
    - 127.0.0.1:6060:6060
    network_mode: host
  mysql:
    image: mysql:5.7
    environment:
    - MYSQL_ROOT_PASSWORD=sQvtf3KKUKgu
    - MYSQL_USER=s002
    - MYSQL_PASSWORD=wvZJFExMtWQpvLkqaoakbQT88
    - MYSQL_DATABASE=s002
    volumes:
    - ./mysql:/var/lib/mysql
    networks:
      innernet:
        ipv4_address: 172.23.0.16
    command: ['--character-set-server=utf8mb4', '--collation-server=utf8mb4_unicode_ci']
  mongodb:
    image: mongo:4.0
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: wvZJFEsMgWQpvLkqaoakbQT88
    volumes:
    - ./mongo:/data/db
    ports:
    - 127.0.0.1:27017:27017
    - 127.0.0.1:27018:27018
    - 127.0.0.1:27019:27019
    networks:
      innernet:
        ipv4_address: 172.23.0.17

networks:
  innernet:
    ipam:
      driver: default
      config:
      - subnet: 172.23.0.1/24
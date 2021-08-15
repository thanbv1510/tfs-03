# Crawler

### Issues

- [ ] **[feat]** Database connection pool
- [x] **[bug]**  Duplicate element data when insert
- [x] **[feat]** Insert batch
- [x] **[feat]** <del>goroutine (single)</del> multiple goroutine
- [x] **[feat]** Logging with zap (console)
- [ ] **[feat]** Logging to file
- [x] **[feat]** reading info from config file
- [x] **[feat]** Using GORM
- [ ] **[feat]** Test case

### Setup environment

- run docker file:

```bash
$ docker-compose up -d
```

### Information

- Website: ```https://www.imdb.com/chart/top/?ref_=nv_mv_250```
- Database: MySQL
    - database name: crawler
    - username: root
    - password: 123456Aa@

### ERD

![ERD](https://github.com/thanbv1510/tfs-03/blob/master/lec-04/crawler/resources/ERD.png)
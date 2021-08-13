# Crawler

### Issues
- [ ] Connection pool
- [x] **[Bug]** - Duplicate element data when insert
- [x] Insert batch
- [x] goroutine
- [ ] multiple goroutine :)
- [x] Logging with zap
- [x] reading info from config file
- [x] Using GORM
- [ ] Test case

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
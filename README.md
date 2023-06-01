# Clone repository:
```
git clone https://github.com/Allayar07/prometheus-gin.git
```
* if you use ssh then this command:
```
git clone git@github.com:Allayar07/prometheus-gin.git
```
# Database schema:
first check your directory. Are you in prometheus-gin/first-service directory? If you aren't in directory, you must enter folder prometheus-gin.(Use this command ``` cd prometheus-gin/first-service```)
# Command for initializing db table:
```migrate -path ./schema -database 'postgres://postgres:password0701@localhost:5432/practice?sslmode=disable' up```

Then check it created successfully follow this commands:
* ```docker exec -it practice-db bash```
* ```psql -U postgres```
*  ```\c practice```
*  ```\d```
Then you will see:
```               List of relations
Schema |       Name        | Type  |  Owner   
--------+-------------------+-------+----------
public | schema_migrations | table | postgres
public | users             | table | postgres
(2 rows)
```
# Run project step by step:
# STEP 1:
first check your directory. Are you in prometheus-gin/first-service directory? If you aren't in directory, you must enter folder prometheus-gin.(Use this command ``` cd prometheus-gin/first-service```)
# STEP 2:
Use ```docker compose up``` command for running prometheus and app!
# For request:
then do request this endpoint from postman, insomnia or browser(Note: METHOD "GET"): http://localhost:8079/say
# See result:
Do request this url for observe metrics: http://localhost:9911/metrics.
* See prometheus UI in this url: http://localhost:19090/

# Grafana configuration:

* import node-exporter's dashboard its id ```1860```

* import cadvisor's dashboard its id ```14282```

[//]: # (* ![img.png]&#40;img.png&#41;)

* write ids and click ```load``` and then select prometheus and then click ```import```
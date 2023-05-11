# Clone repository:
```
git clone https://github.com/Allayar07/prometheus-gin.git
```
* if you use ssh then this command:
```
git clone git@github.com:Allayar07/prometheus-gin.git
```
# Run project step by step:
# STEP 1:
first check your directory. Are you in prometheus directory? If you aren't in directory, you must enter folder prometheus.(Use this command ``` cd prometheus```)
# STEP 2:
Use ```docker compose up``` command for running prometheus!

# STEP 3
Run the project(you must be in prometheus folder) : ```go run cmd/main.go```
Do request this endpoint from postman, insomnia or browser(Note: METHOD "GET"): http://localhost:8079/say

# See result:
Do request this url for observe metrics: http://localhost:9911/metrics.
See prometheus UI in this url: http://localhost:19090/
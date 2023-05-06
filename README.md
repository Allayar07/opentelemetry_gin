# Clone repository:
```
git clone https://github.com/Allayar07/opentelemetry_gin.git
```
* if you use ssh then this command:
```
git clone git@github.com:Allayar07/opentelemetry_gin.git
```
* **Note** if your system don't have docker engine then install it your system.
* for installing docker engine in your system:
* Ubuntu users: click here ===> https://docs.docker.com/engine/install/ubuntu/
* Windows users: click here ===> https://www.simplilearn.com/tutorials/docker-tutorial/install-docker-on-windows
* Mac users: click here ===> https://www.makeuseof.com/how-to-install-docker-mac/

# Run project step by step:
# STEP 1:
first check your directory. Are you in opentelemetry_gin directory? If you aren't in directory, you must enter folder opentelemetry_gin.(Use this command ``` cd opentelemetry_gin```)
# STEP 2:
Use ```docker compose up``` command for running project!

# See result:
Do request this url: http://localhost:8089/call-service.
And go to zipkin UI ===> http://127.0.0.1:9411/
* Then go to right corner and click ```RUN QUERY``` button, and you can see result.


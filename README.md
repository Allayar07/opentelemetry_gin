# Clone repository:
```
git clone https://github.com/Allayar07/opentelemetry_gin.git
```
* if you use ssh then this command:
```
git clone git@github.com:Allayar07/opentelemetry_gin.git
```
# Run project step by step:
# STEP 1:
first check your directory. Are you in opentelemetry_gin directory? If you aren't in directory, you must enter folder opentelemetry_gin.(Use this command ``` cd opentelemetry_gin```)
# STEP 2:
Use ```docker compose up``` command for running project!

Wait for docker, for finishing its work
# See result:
Do request this url: http://localhost:8089/call-service.
And go to zipkin UI ===> http://127.0.0.1:9411/
* Then go to right corner and click ```RUN QUERY``` button, and you can see result.

# See trace in grafana :

* Step-1. Go to this url: http://localhost:3000

* Step-2. Write ```username``` "admin", ```password``` admin.

* Step-3. Then go to Configuration button (it's situated left sidebar, under)

* Step-4. Click ```Explore``` button

* Step-5. Choose trace from ```Traces``` button

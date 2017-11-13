# My Onboaring Project Two

## Description
This is my second project on SRE onboarding to understand the ruby sinatra web application.

## SLO and SLI
* Availability: To be Defined
* Mean Response Time: To be Defined

## Architecture Diagram
```
+--------------------+        +---------------------+         +---------+
|  Project API/CMD   | <--->  |  Ministore Package  |  <--->  |  MySql  |
+--------------------+        +---------------------+         +---------+
```

## Owner
[Setya Kurniawan](https://github.com/SetyaK)

## Contact and On-Call Information
[Setya Kurniawan](setya.kurniawan@bukalapak.com)

## Links

## Onboarding and Development Guide
### Prerequisite
1. Ministore Ruby Gem,  
  [Guide](https://github.com/SetyaK/BL-Onboarding1-Ruby)
### Setup
1. Get ministore libary  
  `$ go get github.com/SetyaK/BL-Onboarding3-Go-package`
2. Install required packages  
  `$ govendor sync`
3. Find the app directory  
  - `app/ministore_cmd` for command line app
  - `app/webAPI` for web API app
4. Create .env file from .env.sample  
  `$ cp .env.sample .env`  
  Then modify it to match your configuration  
  `ENV=development` mean existing table will be recreated when  
  `database.Migration.migrate` was called
5. Build app  
  `$ go build`
### Development Guide
Available command line and web API version to use the ministore package.
- To use the command line run `app/ministore_cmd/ministore_cmd --help`
- To use the web API run `app/webAPI/webAPI`, the web app will run in port 4567
#### API usage
You can use [POSTMAN](https://www.getpostman.com/) to test the API
- Test the app is run  
  `GET /`
- Get list of all products  
  `GET /product/`
- Get a product  
  `GET /product/[id]`
- Add new product  
  `POST /product`  
  Post variable:
  - `name`
  - `description`
  - `initial_stock`
- Update a product
  `POST /product/[id]`  
  Post variable:
  - `name`
  - `description`
- Delete a product  
  `DELETE /product/[id]`

## On-Call Runbooks

## F.A.Q.

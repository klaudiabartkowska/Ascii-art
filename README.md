
### Description 

Ascii-art-web consists in creating and running a server, in which it is possible to use a web GUI (graphical user interface) version of our latest project, ascii-art.

Webpage allow the use of the different banners:

shadow
standard
thinkertoy

Implements the following HTTP endpoints:

GET /: Sends HTML response, the main page.
POST /ascii-art: that sends data to Go server (text and a banner)


The main page includes:

text input,
radio buttons to switch between banners,
button, which sends a POST request to '/ascii-art' and outputs the result on the page.
HTTP status code.


Endpoints returns appropriate HTTP status codes :

OK (200), if everything went without errors.
Not Found, if nothing is found, for example templates or banners.
Bad Request, for incorrect requests.
Internal Server Error, for unhandled errors.

### Usage 
In order to use this program, please clone the repository and run main.go. 
Open your browser on http://localhost:8080/. 
Now you are able to generate text using diffrent banner available. 


### Authors 
Klaudia 

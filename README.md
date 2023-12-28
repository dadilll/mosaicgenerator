# Mosaic Generator

This Go application sets up a simple web server for generating mosaic images. It uses the `net/http` package to handle HTTP requests and serves an HTML form for uploading images and specifying a pixel size. The core mosaic generation logic is encapsulated in the `image` package.

## How to Run

### Clone the Repository
```
git clone https://github.com/your-username/your-repository.git
cd your-repository
```

### Run the Application
```
cd cmd
go run main.go
```
### Access the Application
Open your web browser and go to http://localhost:8080.

You'll see a form to upload an image and set the pixel size.

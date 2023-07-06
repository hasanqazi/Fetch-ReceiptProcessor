# Running the Code - Receipt Processing Service

To run the provided code for the Receipt Processing Service, follow the steps below:

## Prerequisites

- Go programming language installed on your machine.

## Steps

1. Create a new directory for your project.

2. Open a text editor and create a new file named `main.go`.

3. Copy and paste the provided code into the `main.go` file.

4. Save the file.

5. Open a terminal or command prompt and navigate to the directory where you saved `main.go`.

6. Build the Go application by running the following command:

   ```shell
   go build
   ```

   This will create an executable file in the current directory.

7. Run the executable by executing the following command:

   ```shell
   ./main
   ```

   The server will start running and listen on port 8080.

   ```shell
   Starting server on port 8080...
   ```

8. Once the server is running, you can interact with it using HTTP requests. Here are a few examples using `curl` command:

   - Process a receipt:

     ```shell
     curl -X POST -H "Content-Type: application/json" -d '{
       "retailer": "Example Retailer",
       "purchaseDate": "2023/07/06",
       "purchaseTime": "15:30",
       "items": [
         {
           "shortDescription": "Item 1",
           "price": "9.99"
         },
         {
           "shortDescription": "Item 2",
           "price": "4.50"
         }
       ],
       "total": "14.49"
     }' http://localhost:8080/receipts/process
     ```

     This will process the receipt and return a JSON response with the generated ID.

   - Get the points for a receipt:

     ```shell
     curl http://localhost:8080/receipts/{id}/points
     ```

     Replace `{id}` with the generated ID of a processed receipt. This will return a JSON response with the points calculated for the receipt.

9. To stop the server, press `Ctrl + C` in the terminal or command prompt where it's running.
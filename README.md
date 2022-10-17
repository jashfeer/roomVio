# RoomVio
The application was Designed and built for online Room booking with all the
necessary features. The application must provide customers 24*7 hours
online booking service.Design and functionalities are mostly inspired from OYO.

## Technologies and tools used
   **Server:** Golang
   
   **Frontend:** HTML5, CSS3, Bootstrap
   
   **Primary database:** PostgreSQL
   
   **ORM:** GORM
   
   **Authentication:** session and cookies  
   
   **Payment gateway:** RazorPay   
   
   **Container:** Docker  
   
   
   ## Workflow and Features
   
   - Cross Platform
   
   #### User side:-
   
   - Search available rooms with check-in and check-out date
   
   - Show all the facilities of each room/hotel
   
   - Filter all rooms/hotels by each facility, price, and rating
   
   - **Razorpay Payment Gateway** is integrated for users to book a room with pay 20% of the total amount
   
   - Show all their bookings and the status of the reservation
   
   - User can cancel their booking
   
   #### Admin side:-
   
   - Add/delete rooms and hotels
   
   - Show the user's details, block/unblock any user
   
   - Approve/cancel user's bookings
   
   - Show payment details
   
   ## Run locally
   Clone the project
   ```
   https://github.com/jashfeer/roomVio.git
   ```
   Create `mod file`
   ```
     go mod init main.go
   ```
   Add module requirements and sums
   ```
     go mod tidy
   ```
   Start the server
   ```
   go run main.go
   ```
   App will listen to `localhost:8080`
   
   ## Environment Variables
   To run this project, you will need to add the following environment variables to your .env file  
   `dbHost` `dbUser` `dbPassword` `dbName` `dbType` `dbPort` `AWS_REGION` `AWS_ACCESS_KEY_ID` `AWS_SECRET_ACCESS_KEY`
   
   ## To Run in Docker
   Download docker image from dockerhub and run in detached mode
   ```
   sudo docker run -d -p 8080:8080 jashfeer/roomvio:v-1.01 
   ```
   App will listen to `localhost:8080`
   
   ##
   I have created **roomVio** as a study project. The main objective of building this project was learn how to build real world working applications using **Go** with other technologies and tools I have learned so far.
   
   ## Feedback
   If you have any feedback, please reach out at jashfeerabdullaf@gmail.com
   
   

  



   


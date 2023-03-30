# OverlappingDates

## Progress
- added a GET endpoint to get started
- Comands to run the server
    - cd cmd
    - go run main.go

<img width="1025" alt="Screenshot 2023-03-29 at 3 34 59 PM" src="https://user-images.githubusercontent.com/22546296/228683523-e3996d42-16ae-4c90-bfc8-ada3b898636b.png">

- addded proto file for the request and response schema
- commands
    - Install protobuf using brew
    - Run ```protoc --go_out=. --go_opt=paths=source_relative <I>protofile</I>``` to generate proto files

<img width="823" alt="Screenshot 2023-03-29 at 3 47 21 PM" src="https://user-images.githubusercontent.com/22546296/228685337-7acbb0aa-3408-41f0-9bba-a18d79e6ec27.png">

- added handler function which handles the logic of comapring date ranges

<img width="1048" alt="Screenshot 2023-03-29 at 6 40 43 PM" src="https://user-images.githubusercontent.com/22546296/228706172-66c1303e-436e-401e-9acb-d9829a8cd29d.png">

- added parse function to parse the input dates to catch wrong date formats

<img width="1014" alt="Screenshot 2023-03-29 at 7 29 54 PM" src="https://user-images.githubusercontent.com/22546296/228712942-b3270361-4750-4713-94c0-d7979f885aee.png">

# gorangutan
Android / IOS push notification server written in GO

Project in development stage

1. Add registration token to db - Android
        
      Api : localhost:8080/register/
      Method : POST
      Headers : device-id // Unique device id for a device
      Form-Data : registration_id = wggdjsgnxvjsbxjksb // Registration Id recieved from Google 
      Response code : 201

2. Send push notification to registered device - Android
     
      Api : localhost:8080/androidPush/
      Method : POST
      Form-Data : device_ids=1234567,23456,345654&title=title for push&message=message for push

## url : https://hospital-playlist.herokuapp.com/

users

    GET /users
    POST /users/register
    POST /users/login
    GET /users/:user_id
    PUT /users/:user_id
    DELETE /users/:user_id


user details

    GET /user_details
    POST /user_details
    PUT /user_details

User Profile

    GET /user_profile
    POST /user_profile
    PUT /user_profile

Drugs

    GET /drugs
    GET /drug/:drug_id
    POST /drug

Spesialist

    GET    /spesialist             
    GET    /spesialist/:spesialist_id 
    POST   /spesialist             
    PUT /spesialist/:spesialist_id 

Dokter

    GET    /dokters  
    GET    /dokter/:dokter_id      
    POST   /dokter/register         
    POST   /dokter/login    

Booking        
    GET    /booking                  
    GET    /booking/:book_id         
    POST   /booking     

Booking Obat            
    GET    /bookingobat/:booking_id/:drug_id
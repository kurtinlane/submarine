# submarine
it's like two people in a submarine turning their keys to launch a nuclear bomb or in this case, deleting your data. 


[GO getting started link] (https://cloud.google.com/appengine/docs/go/gettingstarted/devenvironment)

A webserver that takes an email and returns a key. 

# overview

Here's the idea. A service provider--[bettrnet](http://bettrnet.com), for example--wants to provide reassurance to its customers that their data is in good hands. Bettrnet's customers want to use bettrnet, but they are worried about giving up their data.

Bettrnet uses submarine to give customer's power of their own data. When User signs up with bettrnet, bettrnet sends the User's email address to submarine. Submarine creates a random string that belongs to User and returns it to bettrnet. Submarine also sends instructions to User on how they can use submarine to control their data.

Bettrnet uses this random string as an encryption key to encrypt User's data. Everytime User logs in to bettrnet, bettrnet has to ask submarine for User's key. 

If at some point User does not trust bettrnet with their data anymore, User can delete their bettrnet key through submarine. Now, bettrnet no longer has access to any data that it encrypted with User's key.

